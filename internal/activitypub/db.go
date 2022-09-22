package activitypub

import (
	"context"
	"fmt"
	"net/url"
	"sync"

	"github.com/eriner/burr/internal/ent"
	"github.com/go-fed/activity/pub"
	"github.com/go-fed/activity/streams/vocab"
)

var _ pub.Database = (*APDB)(nil)

// APDB implements github.com/go-fed/activity/pub.Database
type APDB struct {
	Host string // fqdn of the server
	DB   *ent.Client

	content *sync.Map // TODO: replace and implement ent database schema
	locks   *sync.Map
}

// content is ap data
type content struct {
	data  vocab.Type
	local bool // true if this server owns this content
}

// Lock takes a lock for the object at the specified id. If an error
// is returned, the lock must not have been taken.
//
// The lock must be able to succeed for an id that does not exist in
// the database. This means acquiring the lock does not guarantee the
// entry exists in the database.
//
// Locks are encouraged to be lightweight and in the Go layer, as some
// processes require tight loops acquiring and releasing locks.
//
// Used to ensure race conditions in multiple requests do not occur.
func (p *APDB) Lock(c context.Context, id *url.URL) error {
	mu := &sync.Mutex{}
	mu.Lock()
	i, exists := p.locks.LoadOrStore(id.String(), mu)
	if exists {
		mu = i.(*sync.Mutex)
		mu.Lock()
	}
	return nil
}

// Unlock makes the lock for the object at the specified id available.
// If an error is returned, the lock must have still been freed.
//
// Used to ensure race conditions in multiple requests do not occur.
func (p *APDB) Unlock(c context.Context, id *url.URL) error {
	i, exists := p.locks.Load(id.String())
	if !exists {
		return fmt.Errorf("failed to find id during Unlock")
	}
	mu := i.(*sync.Mutex)
	mu.Unlock()
	return nil
}

// InboxContains returns true if the OrderedCollection at 'inbox'
// contains the specified 'id'.
//
// The library makes this call only after acquiring a lock first.
func (p *APDB) InboxContains(c context.Context, inbox *url.URL, id *url.URL) (contains bool, err error) {
	// "Our goal is to see if the `inbox`, which is an ORderedCollection,
	// contains an element in its `ordered_items` property that has a
	// matching `id`"
	var oc vocab.ActivityStreamsOrderedCollection
	oc, err = p.getOrderedCollection(inbox) // TODO
	if err != nil {
		return
	}
	// "Next, we use the ACtivityStreams vocabulary to obtain the
	// `ordered_items` property of the OrderedCollection type."
	oi := oc.GetActivityStreamsOrderedItems()
	if oi == nil {
		return
	}
	// "Finally, loop through each item in the `ordered_items` property and see
	// if the element's `id` matches the desired `id`.

	for i := oi.Begin(); i != oi.End(); i = i.Next() {
		var eid *url.URL
		id, err = pub.ToId(i)
		if err != nil {
			return
		}
		if eid.String() == id.String() {
			return true, nil
		}
	}
	return
}

// GetInbox returns the first ordered collection page of the outbox at
// the specified IRI, for prepending new items.
//
// The library makes this call only after acquiring a lock first.
func (p *APDB) GetInbox(c context.Context, inboxIRI *url.URL) (inbox vocab.ActivityStreamsOrderedCollectionPage, err error) {
	// fetch an inbox
	return p.getOrderedCollectionPage(inboxIRI) // TODO
}

// SetInbox saves the inbox value given from GetInbox, with new items
// prepended. Note that the new items must not be added as independent
// database entries. Separate calls to Create will do that.
//
// The library makes this call only after acquiring a lock first.
func (p *APDB) SetInbox(c context.Context, inbox vocab.ActivityStreamsOrderedCollectionPage) error {
	storedInbox, err := p.getOrderedCollection(inboxIRI)
	if err != nil {
		return err
	}
	updatedInbox := p.applyDiffOrderedCollection(storedInbox, inbox)

	return p.saveToContent(updatedInbox)
}

// Owns returns true if the database has an entry for the IRI and it
// exists in the database.
//
// The library makes this call only after acquiring a lock first.
func (p *APDB) Owns(c context.Context, id *url.URL) (owns bool, err error) {
	// TODO: documentation says to use something "more robust than this string comparison", whatever that means.
	return id.Host == p.host, nil
}

// ActorForOutbox fetches the actor's IRI for the given outbox IRI.
//
// The library makes this call only after acquiring a lock first.
func (p *APDB) ActorForOutbox(c context.Context, outboxIRI *url.URL) (actorIRI *url.URL, err error) {
	panic("not implemented")
}

// ActorForInbox fetches the actor's IRI for the given outbox IRI.
//
// The library makes this call only after acquiring a lock first.
func (p *APDB) ActorForInbox(c context.Context, inboxIRI *url.URL) (actorIRI *url.URL, err error) {
	panic("not implemented") // TODO: Implement
}

// OutboxForInbox fetches the corresponding actor's outbox IRI for the
// actor's inbox IRI.
//
// The library makes this call only after acquiring a lock first.
func (p *APDB) OutboxForInbox(c context.Context, inboxIRI *url.URL) (outboxIRI *url.URL, err error) {
	panic("not implemented") // TODO: Implement
}

// Exists returns true if the database has an entry for the specified
// id. It may not be owned by this application instance.
//
// The library makes this call only after acquiring a lock first.
func (p *APDB) Exists(ctx context.Context, id *url.URL) (exists bool, err error) {
	_, exists = p.content.Load(id.String())
	return exists, nil
}

// Get returns the database entry for the specified id.
//
// The library makes this call only after acquiring a lock first.
func (p *APDB) Get(ctx context.Context, id *url.URL) (value vocab.Type, err error) {
	v, ok := p.content.Load(id.String())
	if !ok {
		return nil, fmt.Errorf("Get failed")
	}
	c := v.(*content)
	return c.data, nil
}

// Create adds a new entry to the database which must be able to be
// keyed by its id.
//
// Note that Activity values received from federated peers may also be
// created in the database this way if the Federating Protocol is
// enabled. The client may freely decide to store only the id instead of
// the entire value.
//
// The library makes this call only after acquiring a lock first.
//
// Under certain conditions and network activities, Create may be called
// multiple times for the same ActivityStreams object.
func (p *APDB) Create(ctx context.Context, asType vocab.Type) error {
	// could be local or federated content.
	id, err := pub.GetId(asType)
	if err != nil {
		return err
	}
	owns, err := p.Owns(ctx, id)
	if err != nil {
		return err
	}
	p.content.Store(id.String(), &content{
		data:  asType,
		local: owns,
	})
	return nil
}

// Update sets an existing entry to the database based on the value's
// id.
//
// Note that Activity values received from federated peers may also be
// updated in the database this way if the Federating Protocol is
// enabled. The client may freely decide to store only the id instead of
// the entire value.
//
// The library makes this call only after acquiring a lock first.
func (p *APDB) Update(ctx context.Context, asType vocab.Type) error {
	// For the actual DB, this will need to be a real update
	// call, but this in-memory map can just call create to
	// replace the value.
	return p.Create(ctx, asType)
}

// Delete removes the entry with the given id.
//
// Delete is only called for federated objects. Deletes from the Social
// Protocol instead call Update to create a Tombstone.
//
// The library makes this call only after acquiring a lock first.
func (p *APDB) Delete(ctx context.Context, id *url.URL) error {
	p.content.Delete(id.String())
	return nil
}

// GetOutbox returns the first ordered collection page of the outbox
// at the specified IRI, for prepending new items.
//
// The library makes this call only after acquiring a lock first.
func (p *APDB) GetOutbox(c context.Context, outboxIRI *url.URL) (inbox vocab.ActivityStreamsOrderedCollectionPage, err error) {
	return p.getOrderedCollectionPage(outboxIRI) // TODO
}

// SetOutbox saves the outbox value given from GetOutbox, with new items
// prepended. Note that the new items must not be added as independent
// database entries. Separate calls to Create will do that.
//
// The library makes this call only after acquiring a lock first.
func (p *APDB) SetOutbox(c context.Context, outbox vocab.ActivityStreamsOrderedCollectionPage) error {
	storedOutbox, err := p.getOrderedCollection(outboxIRI)
	if err != nil {
		return err
	}
	updatedOutbox := p.applyDiffOrderedCollection(storedOutbox, outbox)
	return p.saveToContent(updatedOutbox)
}

// NewID creates a new IRI id for the provided activity or object. The
// implementation does not need to set the 'id' property and simply
// needs to determine the value.
//
// The go-fed library will handle setting the 'id' property on the
// activity or object provided with the value returned.
func (p *APDB) NewID(c context.Context, t vocab.Type) (id *url.URL, err error) {
	panic("not implemented") // TODO: Implement
}

// Followers obtains the Followers Collection for an actor with the
// given id.
//
// If modified, the library will then call Update.
//
// The library makes this call only after acquiring a lock first.
func (p *APDB) Followers(c context.Context, actorIRI *url.URL) (followers vocab.ActivityStreamsCollection, err error) {
	var person vocab.ActivityStreamsPerson
	person, err = p.getPerson(actorIRI)
	if err != nil {
		return
	}
	f := person.GetActivityStreamsFollowers()
	if f == nil {
		return nil, fmt.Errorf("no followers collection")
	}
	// f could be an OrderedCollection, IRI, or something extending an OrderedCollection
	fid, err := pub.ToId(f)
	if err != nil {
		return nil, fmt.Errorf("failed to get id: %w", err)
	}
	return p.getOrderedCollection(fid)
}

// Following obtains the Following Collection for an actor with the
// given id.
//
// If modified, the library will then call Update.
//
// The library makes this call only after acquiring a lock first.
func (p *APDB) Following(c context.Context, actorIRI *url.URL) (followers vocab.ActivityStreamsCollection, err error) {
	// similar to Followers()
	panic("not implemented")
}

// Liked obtains the Liked Collection for an actor with the
// given id.
//
// If modified, the library will then call Update.
//
// The library makes this call only after acquiring a lock first.
func (p *APDB) Liked(c context.Context, actorIRI *url.URL) (followers vocab.ActivityStreamsCollection, err error) {
	// similar to Followers()
	panic("not implemented")
}

func (p *APDB) getOrderedCollection(iri *url.URL) (vocab.ActivityStreamsOrderedCollection, error) {
	panic("TODO")
}

func (p *APDB) getPerson(iri *url.URL) (vocab.ActivityStreamsPerson, error) {
	panic("TODO")
}

func (p *APDB) saveToContent(vocab.ActivityStreamsCollection) error {
	panic("TODO")
}

func (p *APDB) applyDiffOrderedCollection(x, y vocab.ActivityStreamsOrderedCollection) vocab.ActivityStreamsCollection {
	panic("TODO")
}

func (p *APDB) getOrderedCollectionPage(iri *url.URL) (vocab.ActivityStreamsOrderedCollectionPage, error) {
	panic("TODO")
}
