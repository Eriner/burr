package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var workerCmd = &cobra.Command{
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
		return worker(cfg)
	},
}

func worker(cfg *Config) error {
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

	//
	// Queuing
	//
	//TODO

	//
	// External S3
	//
	//TODO

	//
	// Internal S3
	//
	//TODO

	panic("TODO")
}
