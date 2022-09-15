package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	database "github.com/eriner/burr/internal/db"
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
	//
	// Secrets
	//
	vaultToken := "root" // TODO: read token from stdin or an alternate then-unmounted file path
	if len(cfg.Vault.Addr) == 0 {
		return fmt.Errorf("vault.addr not specified")
	}
	u, err := url.Parse(cfg.Vault.Addr)
	if err != nil {
		return fmt.Errorf("invalid vault.addr value of %q: %w", cfg.Vault.Addr, err)
	}
	vaultClient := &http.Client{
		Timeout: 5 * time.Second,
	}
	switch u.Scheme {
	case "http":
		log.Println("WARNING: Vault address is not HTTPS")
		vaultClient.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}
	case "https":
		break
	default:
		return fmt.Errorf("vault.addr does not start with \"http://\" or \"https://\"")
	}
	vault, err := secrets.Vault(cfg.Vault.Addr, vaultToken, vaultClient)
	if err != nil {
		return fmt.Errorf("failed to connect to Vault: %w", err)
	}
	if err := vault.Init(); err != nil {
		return err
	}
	log.Printf("Connected to Vault")

	//
	// Queuing
	//
	//TODO

	//
	// Database
	//
	dbCfg := cfg.DB
	if dbCfg["driver"] == "postgres" || dbCfg["driver"] == "postgresql" {
		res, err := vault.ReadWithContext(context.Background(), cfg.Vault.BurrPath+"/db_password")
		if err != nil {
			return fmt.Errorf("failed to read postgres password from Vault: %w", err)
		}
		if len(res["password"].(string)) == 0 {
			return fmt.Errorf("Vault has %s/db_password entry, but no \"password\" data value: %w",
				cfg.Vault.BurrPath, err)
		}
		dbCfg["password"] = res["password"]
		res = nil // GC
	}
	db, err := database.Open(dbCfg)
	if err != nil {
		return fmt.Errorf("failed to connect to the database: %w", err)
	}
	defer db.Close()
	dbCfg = nil // GC
	log.Println("Connected to database")

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
	cache, err := kv.Open(nil)
	if err != nil {
		return fmt.Errorf("failed to connect to the redis worker cache: %w", err)
	}
	_ = cache // TODO will almost certainly be necessary in the near future
	log.Println("Connected to worker cache")

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
