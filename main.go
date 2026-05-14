package main

import (
	"log"

	"github.com/whitenhiemer/alertmind/internal/config"
	"github.com/whitenhiemer/alertmind/internal/server"
)

func main() {
	cfg := config.Load()

	if cfg.AnthropicAPIKey == "" {
		log.Fatal("ANTHROPIC_API_KEY is required")
	}

	log.Printf("alertmind starting on :%s (model=%s fetch_runbooks=%v)",
		cfg.Port, cfg.Model, cfg.FetchRunbooks)

	srv := server.New(cfg)
	log.Fatal(srv.Start())
}
