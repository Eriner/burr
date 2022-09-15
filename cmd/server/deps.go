package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	fserver "github.com/contribsys/faktory/server"
	"github.com/contribsys/faktory/webui"
	fworker "github.com/contribsys/faktory_worker_go"
	database "github.com/eriner/burr/internal/db"
	"github.com/eriner/burr/internal/ent"
	"github.com/eriner/burr/internal/faktory"
	"github.com/eriner/burr/internal/kv"
	"github.com/eriner/burr/internal/secrets"
)

func secretsFromCfg(cfg *Config) (secrets.Store, error) {
	if cfg == nil {
		return nil, fmt.Errorf("incomplete arguments")
	}
	token := "root" // TODO: read token from stdin or an alternate then-unmounted file path
	if len(cfg.Secrets.Addr) == 0 {
		return nil, fmt.Errorf("secrets.addr is required but was not provided")
	}
	u, err := url.Parse(cfg.Secrets.Addr)
	if err != nil {
		return nil, fmt.Errorf("invalid secrets.addr value of %q: %w", cfg.Secrets.Addr, err)
	}
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	switch u.Scheme {
	case "http":
		log.Println("WARNING: address is not HTTPS")
		client.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}
	case "https":
		break
	default:
		return nil, fmt.Errorf("vault.addr does not start with \"http://\" or \"https://\"")
	}
	vault, err := secrets.Open(map[string]any{
		"addr":   cfg.Secrets.Addr,
		"token":  token,
		"client": client,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Secrets provider: %w", err)
	}
	return vault, nil
}

func dbFromCfg(cfg *Config, secrets secrets.Store) (*ent.Client, error) {
	if cfg == nil || secrets == nil {
		return nil, fmt.Errorf("incomplete arguments")
	}
	c := cfg.DB
	if c != nil {
		c = make(map[string]any)
		c["driver"] = "sqlite"
	}
	switch c["driver"] {
	case "postgres", "postgresql":
		user, ok := c["user"].(string)
		if !ok {
			user = "postgres"
		}
		dbName, ok := c["database"].(string)
		if !ok {
			dbName = "burr"
		}
		addr, ok := c["addr"].(string)
		if !ok {
			addr = "127.0.0.1:5432"
		}
		x := strings.Split(addr, ":")
		var host, port string
		if len(x) == 1 {
			host = addr
			port = "5432"

		} else if len(x) == 2 {
			host = x[0]
			port = x[2]
		}

		res, err := secrets.ReadWithContext(context.Background(), cfg.Secrets.Path+"/db_password")
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve postgres password: %w", err)
		}
		if len(res["password"].(string)) == 0 {
			return nil, fmt.Errorf("Secrets has %s/db_password entry, but no \"password\" value: %w",
				cfg.Secrets.Path, err)
		}
		password, ok := res["password"].(string)
		if !ok {
			return nil, fmt.Errorf("failed to retrieve database password from secrets store")
		}
		res = nil // GC
		return database.Postgres(host, port, user, password, dbName)
	default: // "sqlite"
		return database.Sqlite()

	}
}

func kvFromCfg(cfg *Config, secrets secrets.Store) (store kv.KV, err error) {
	if cfg == nil || secrets == nil {
		return nil, fmt.Errorf("incomplete arguments")
	}
	c := cfg.KV
	if c == nil {
		c = make(map[string]any)
		c["driver"] = "builtin"
	}
	switch c["driver"] {
	case "redis":
		if len(c) == 0 {
			return nil, fmt.Errorf("redis requires configuration")
		}
		addr, ok := c["addr"].(string)
		if !ok {
			return nil, fmt.Errorf("redis configuration must include \"addr\"")
		}
		password, ok := c["password"].(string)
		if !ok {
			password = ""
		}
		return kv.Redis(addr, password)
	case "builtin":
		return kv.Builtin()
	default:
		return nil, fmt.Errorf("unknown KV driver %q", c["driver"])
	}
}

func faktoryFromCfg(cfg *Config) (srv *fserver.Server, srvClose func(), worker *fworker.Manager, err error) {
	if cfg == nil {
		return nil, nil, nil, fmt.Errorf("incomplete arguments")
	}
	c := cfg.Faktory
	if c == nil {
		c = make(map[string]any)
		c["bind"] = "http://127.0.0.1:7419"
	}
	bind, ok := c["bind"].(string)
	if !ok {
		bind = ""
	}
	addr, ok := c["addr"].(string)
	if !ok {
		addr = bind
	}
	dir, ok := c["dir"].(string)
	if !ok {
		td, err := ioutil.TempDir("", "")
		if err != nil {
			return nil, nil, nil, fmt.Errorf("failed to create temporary directory for faktory: %w", err)
		}
		dir = td
	}
	ui, ok := c["ui"].(string)
	if !ok {
		ui = ""
	}
	if len(ui) == 0 {
		ui = "http://127.0.0.1:7420"
	}
	ui = strings.TrimPrefix(ui, "https://")
	ui = strings.TrimPrefix(ui, "http://")

	if len(bind) > 0 {
		bind = strings.TrimPrefix(bind, "https://")
		bind = strings.TrimPrefix(bind, "http://")
		srv, srvClose, err = faktory.LocalServer(bind, dir)
		if err != nil {
			return nil, nil, nil, fmt.Errorf("failed to bind and boot local Faktory server: %w", err)
		}
		srv.Register(webui.Subsystem(ui))
	}
	if len(addr) > 0 {
		os.Setenv("FAKTORY_URL", addr) // TODO: is there really not a better way to do this?
		worker = fworker.NewManager()
	}
	return srv, srvClose, worker, nil
}
