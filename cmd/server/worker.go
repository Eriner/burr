package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
)

var workerCmd *cobra.Command = func() *cobra.Command {
	var threads int
	var queues []string
	cmd := &cobra.Command{
		Use:     "worker",
		Aliases: []string{"work", "w"},
		RunE: func(cmd *cobra.Command, _ []string) error {
			f, err := cmd.Flags().GetString("config")
			if err != nil {
				return fmt.Errorf("failed to read config flag: %w", err)
			}
			cfg, err := ConfigFromFile(f)
			if err != nil {
				return fmt.Errorf("failed to read config file %q: %w", f, err)
			}
			return worker(cfg, threads, queues)
		},
	}
	cmd.Flags().StringSliceVarP(&queues, "queues", "q", []string{"default"}, "queues to process with this worker")
	cmd.Flags().IntVarP(&threads, "threads", "t", 20, "maximum threads per core. Determines worker concurrency")
	return cmd
}()

func worker(cfg *Config, threads int, queues []string) error {
	secrets, err := secretsFromCfg(cfg)
	if err != nil {
		return err
	}
	log.Println("Connected to Secrets store")

	db, err := dbFromCfg(cfg, secrets)
	if err != nil {
		return err
	}
	log.Println("Connected to DB")
	_ = db

	kv, err := kvFromCfg(cfg, secrets)
	if err != nil {
		return err
	}
	log.Println("Connected to KV cache")
	_ = kv

	_, _, worker, err := faktoryFromCfg(cfg)
	if err != nil {
		return err
	}
	worker.Concurrency = threads
	worker.ProcessStrictPriorityQueues(queues...)
	errCh := make(chan error)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		errCh <- worker.RunWithContext(ctx)
	}()
	go func() {
		sigCh := make(chan os.Signal, 2)
		signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)

		for {
			select {
			case <-ctx.Done():
				return
			case <-sigCh:
				cancel()
			}
		}
	}()
	log.Println("Connected to Faktory. Starting wurk")
	<-ctx.Done()
	return <-errCh
}
