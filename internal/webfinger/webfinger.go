package webfinger

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

// API implements webfinger
func API(r *chi.Mux) {
	r.Route("/.well-known/webfinger", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			params := r.URL.Query()
			resource := strings.Split(
				strings.TrimPrefix(params.Get("resource"), "acct:"),
				"@",
			)
			if len(resource) != 2 {
				// malformed user contained either no user, no domain, or multiple `@`s.
				panic("malformed webfinger resource")
			}
			username, domain := resource[0], resource[1]

			iriPath := "" // TODO

			wf := &webfinger{
				Subject: "acct:" + username + "@" + domain,
				Aliases: []string{
					"https://" + domain + iriPath,
				},
				Links: []link{
					{
						Rel:  "self",
						Type: "application/activity+json",
						Href: "https://" + domain + iriPath,
					},
				},
			}
			dat, err := json.Marshal(wf)
			if err != nil {
				panic(fmt.Errorf("failed to marshal webfinger json: %w", err))
			}
			w.Header().Set("Content-Type", "application/jrd+json")
			_, err = w.Write(dat)
			if err != nil {
				panic(fmt.Errorf("failed to write webfinger response: %w", err))
			}
		})
	})
}

type webfinger struct {
	Subject string   `json:"subject,omitempty"`
	Aliases []string `json:"aliases,omitempty"`
	Links   []link   `json:"links,omitempty"`
}

type link struct {
	Rel      string `json:"rel,omitempty"`
	Type     string `json:"type,omitempty"`
	Href     string `json:"href,omitempty"`
	Template string `json:"template,omitempty"`
}
