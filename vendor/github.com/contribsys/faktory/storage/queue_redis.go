package storage

import (
	"context"
	"encoding/json"
	"time"

	"github.com/contribsys/faktory/client"
	"github.com/contribsys/faktory/util"
	"github.com/go-redis/redis"
)

type redisQueue struct {
	name  string
	store *redisStore
	done  bool
}

func (store *redisStore) NewQueue(name string) *redisQueue {
	return &redisQueue{
		name:  name,
		store: store,
		done:  false,
	}
}

func (q *redisQueue) Pause() error {
	return q.store.rclient.SAdd("paused", q.name).Err()
}

func (q *redisQueue) Resume() error {
	return q.store.rclient.SRem("paused", q.name).Err()
}

func (q *redisQueue) IsPaused() bool {
	b, _ := q.store.rclient.SIsMember("paused", q.name).Result()
	return b
}

func (q *redisQueue) Close() {
	q.done = true
}

func (q *redisQueue) Name() string {
	return q.name
}

func (q *redisQueue) Page(start int64, count int64, fn func(index int, data []byte) error) error {
	index := 0

	slice, err := q.store.rclient.LRange(q.name, start, start+count).Result()
	for idx := range slice {
		err = fn(index, []byte(slice[idx]))
		if err != nil {
			return err
		}
		index += 1
	}
	return err
}

func (q *redisQueue) Each(fn func(index int, data []byte) error) error {
	return q.Page(0, -1, fn)
}

func (q *redisQueue) Clear() (uint64, error) {
	q.store.mu.Lock()
	defer q.store.mu.Unlock()

	_, err := q.store.rclient.Pipelined(func(pipe redis.Pipeliner) error {
		pipe.Unlink(q.name)
		pipe.SRem("queues", q.name)
		pipe.SRem("paused", q.name)
		return nil
	})
	if err != nil {
		return 0, err
	}

	delete(q.store.queueSet, q.name)
	return 0, nil
}

func (q *redisQueue) init() error {
	util.Debugf("Queue init: %s %d elements", q.name, q.Size())
	return nil
}

func (q *redisQueue) Size() uint64 {
	return uint64(q.store.rclient.LLen(q.name).Val())
}

func (q *redisQueue) Add(job *client.Job) error {
	job.EnqueuedAt = util.Nows()
	data, err := json.Marshal(job)
	if err != nil {
		return err
	}

	return q.Push(data)
}

func (q *redisQueue) Push(payload []byte) error {
	q.store.rclient.LPush(q.name, payload)
	return nil
}

// non-blocking, returns immediately if there's nothing enqueued
func (q *redisQueue) Pop() ([]byte, error) {
	if q.done {
		return nil, nil
	}

	return q._pop()
}

func (q *redisQueue) _pop() ([]byte, error) {
	val, err := q.store.rclient.RPop(q.name).Result()
	if val == "" {
		return nil, nil
	}
	return []byte(val), err
}

func (q *redisQueue) BPop(ctx context.Context) ([]byte, error) {
	val, err := q.store.rclient.BRPop(2*time.Second, q.name).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, err
	}

	return []byte(val[1]), nil
}

func (q *redisQueue) Delete(vals [][]byte) error {
	for idx := range vals {
		err := q.store.rclient.LRem(q.name, 1, vals[idx]).Err()
		if err != nil {
			return err
		}
	}

	return nil
}
