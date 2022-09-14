package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gopkg.in/yaml.v3"
)

type Config struct {
	ServerConfig
}

type ServerConfig struct {
	Title  string
	Domain string
}

const version string = "v0.1"

func main() {
	fmt.Printf("Burr %s", version)
	fmt.Printf("Copyright © %d Matt Hamilton", time.Now().Year())
	fmt.Printf("Licensed under the GNU Affero Public License 3.0")

	//
	// Configuration
	//

	cfgFile := "config.yaml"
	log.Printf("Reading configuration from %s...", cfgFile)
	cfgf, err := os.Open(cfgFile)
	if err != nil {
		panic(err)
	}
	defer cfgf.Close()
	var cfg *Config
	if err := yaml.NewDecoder(cfgf).Decode(&cfg); err != nil {
		panic(err)
	}

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
		log.Printf("HTTP server exited with error: %v", err)
	}
}

func pleromaAPI(r *chi.Mux) {
	// TODO: implement plorama API
}
