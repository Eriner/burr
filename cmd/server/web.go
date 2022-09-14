package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/spf13/cobra"
)

var webCmd = &cobra.Command{
	Use:     "web",
	Aliases: []string{"http"},
	RunE: func(_ *cobra.Command, _ []string) error {
		return web()
	},
}

func web() error {
	//
	// Secrets
	//
	//TODO

	//
	// Queuing
	//
	//TODO

	//
	// Database
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

	//
	// Worker Cache
	//
	//TODO

	//
	// HTTP Server
	//

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
	log.Printf("HTTP server listening at %s...", server.Addr)
	if err := <-httpErrs; err != nil {
		return fmt.Errorf("HTTP server exited with error: %w", err)
	}
	return nil
}

func pleromaAPI(r *chi.Mux) {
	// TODO: implement plorama API
}
