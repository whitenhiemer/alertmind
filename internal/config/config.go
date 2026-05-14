package config

import "os"

type Config struct {
	AnthropicAPIKey   string
	SlackWebhookURL   string
	DiscordWebhookURL string
	Port              string
	Model             string
	FetchRunbooks     bool
}

func Load() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	model := os.Getenv("ALERTMIND_MODEL")
	if model == "" {
		model = "claude-haiku-4-5-20251001"
	}
	return &Config{
		AnthropicAPIKey:   os.Getenv("ANTHROPIC_API_KEY"),
		SlackWebhookURL:   os.Getenv("SLACK_WEBHOOK_URL"),
		DiscordWebhookURL: os.Getenv("DISCORD_WEBHOOK_URL"),
		Port:              port,
		Model:             model,
		FetchRunbooks:     os.Getenv("FETCH_RUNBOOKS") != "false",
	}
}
