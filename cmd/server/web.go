package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/eriner/burr/internal/ent"
	"github.com/eriner/burr/internal/kv"
	"github.com/eriner/burr/internal/secrets"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/spf13/cobra"
)

var webCmd = &cobra.Command{
	Use:     "web",
	Aliases: []string{"http"},
	RunE: func(cmd *cobra.Command, _ []string) error {
		f, err := cmd.Flags().GetString("config")
		if err != nil {
			return fmt.Errorf("failed to read config flag: %w", err)
		}
		cfg, err := ConfigFromFile(f)
		if err != nil {
			return fmt.Errorf("failed to read config file %q: %w", f, err)
		}
		return web(cfg)
	},
}

func web(cfg *Config) error {
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

	kv, err := kvFromCfg(cfg, secrets)
	if err != nil {
		return err
	}
	log.Println("Connected to KV cache")

	faktorySrv, faktorySrvClose, worker, err := faktoryFromCfg(cfg)
	if err != nil {
		return err
	}
	if faktorySrv != nil {
		go func() {
			// TODO: pass err to chan to handle or signal exit
			log.Println(faktorySrv.Run())
		}()
		defer faktorySrvClose()
		log.Printf("Faktory server listening at http://%s", faktorySrv.Options.Binding)
	}
	if worker != nil {
		go func() {
			// TODO: pass err to chan to handle or signal exit
			log.Println(worker.Run())
		}()
	}
	log.Println("Connected to Faktory")

	// TODO External S3

	// TODO Internal S3

	//
	// HTTP Server
	//
	return api(secrets, db, kv)
}

func api(secrets secrets.Store, db *ent.Client, kv kv.KV) error {
	r := chi.NewRouter()
	r.Use(middleware.RealIP)
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	// Implementations of existing APIs
	pleromaAPI(r)
	// TODO: mastodonAPI(r)

	server := &http.Server{
		Addr:         "127.0.0.1:3333",
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	httpErrs := make(chan error)
	go func() {
		httpErrs <- server.ListenAndServe()
	}()
	log.Printf("HTTP API server listening at %s...", server.Addr)
	if err := <-httpErrs; err != nil {
		return fmt.Errorf("HTTP server exited with error: %w", err)
	}
	return nil
}

func pleromaAPI(r *chi.Mux) {
	// TODO: implement plorama API
}
