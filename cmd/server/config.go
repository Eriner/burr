package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server
	Secrets
	DB      map[string]any
	KV      map[string]any
	Faktory map[string]any
}

type Server struct {
	Title  string
	Domain string
}

type Secrets struct {
	Addr string
	Path string
}

func ConfigFromFile(path string) (*Config, error) {
	cfgf, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open YAML config file %q: %w", path, err)
	}
	defer cfgf.Close()
	var cfg *Config
	if err := yaml.NewDecoder(cfgf).Decode(&cfg); err != nil {
		return nil, fmt.Errorf("unable to parse YAML config file %q: %w", path, err)
	}
	log.Printf("Read configuration from %s", path)
	return cfg, nil
}
