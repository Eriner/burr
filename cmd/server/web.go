package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/eriner/burr/internal/activitypub"
	"github.com/eriner/burr/internal/ent"
	"github.com/eriner/burr/internal/kv"
	"github.com/eriner/burr/internal/secrets"
	"github.com/eriner/burr/internal/webfinger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-fed/activity/pub"
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
	errCh := make(chan error) // anything sent over this chan indicates failure; log and exit
	listenhttp(errCh, secrets, db, kv, cfg.Domain)
	for err := range errCh {
		return fmt.Errorf("HTTP server exited with error: %w", err)
	}
	return nil
}

func listenhttp(errCh chan<- error, secrets secrets.Store, db *ent.Client, kv kv.KV, fqdn string) {
	r := chi.NewRouter()
	r.Use(middleware.RealIP)
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	// Plumbing APIs
	webfinger.API(r)
	hostMetaAPI(r, fqdn)
	activitypubAPI(r, db, fqdn)

	// Client (compatibility for apps, FEs) APIs
	mastodonAPI(r)
	// pleromaAPI(r)
	devUI(r)

	server := &http.Server{
		Addr:         "127.0.0.1:3333",
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	go func() {
		errCh <- server.ListenAndServe()
	}()
	log.Printf("HTTP API server listening at %s...", server.Addr)
}

//func pleromaAPI(r *chi.Mux) {
//	// TODO: implement plorama API
//}

func devUI(r *chi.Mux) {
	r.Route("/@{username}", func(r chi.Router) {
	})
}

func mastodonAPI(r *chi.Mux) {
	// TODO: implement mastodon API, precisely what ActivityPub C2S should prevent,
	// but hasn't due to problems in the spec (lack of identity, schema vocab, etc.)
}

func hostMetaAPI(r *chi.Mux, fqdn string) {
	r.Route("/.well-known/host-meta", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/xrd+xml")
			dat := `<?xml version="1.0" encoding="UTF-8"?>
<XRD xmlns="http://docs.oasis-open.org/ns/xri/xrd-1.0">
  <Link rel="lrdd" type="application/xrd+xml" template="https://` + fqdn + `/.well-known/webfinger?resource={uri}"/>
</XRD>`
			_, err := w.Write([]byte(dat))
			if err != nil {
				panic(fmt.Errorf("failed to write hostmeta response: %w", err))
			}
		})
	})
}

func activitypubAPI(r *chi.Mux, client *ent.Client, fqdn string) {
	s := &activitypub.Service{}
	d := &activitypub.APDB{
		Host: fqdn,
		DB:   client,
	}
	actor := pub.NewFederatingActor(s, s, d, s)
	asHandler := pub.NewActivityStreamsHandler(d, s)
	_ = asHandler

	r.Route("/actors/{username}", func(r chi.Router) {
		r.Get("/inbox", func(w http.ResponseWriter, r *http.Request) {
			handled, err := actor.GetInbox(r.Context(), w, r)
			if err != nil {
				panic(err)
			}
			if !handled {
				panic("that's weird...")
			}
		})
		r.Post("/inbox", func(w http.ResponseWriter, r *http.Request) {
			handled, err := actor.PostInbox(r.Context(), w, r)
			if err != nil {
				panic(err)
			}
			if !handled {
				panic("that's weird...")
			}
		})
		r.Get("/outbox", func(w http.ResponseWriter, r *http.Request) {
			handled, err := actor.GetOutbox(r.Context(), w, r)
			if err != nil {
				panic(err)
			}
			if !handled {
				panic("that's weird...")
			}
		})
		r.Post("/outbox", func(w http.ResponseWriter, r *http.Request) {
			// C2S only; not implemented
			/*
				handled, err := actor.PostOutbox(r.Context(), w, r)
				if err != nil {
					panic(err)
				}
				if !handled {
					panic("that's weird...")
				}
			*/
		})
	})
}
