package config

import (
	"github.com/fwidjaya20/symphonic/facades"
	"github.com/golang-module/carbon"
)

func init() {
	config := facades.Config()

	config.Add("repository", map[string]any{
		"connections": map[string]any{
			"postgresql": map[string]any{
				"driver":     "postgresql",
				"host":       config.Get("DB_HOST", "localhost"),
				"port":       config.Get("DB_PORT", 5432),
				"repository": config.Get("DB_DATABASE", "forge"),
				"username":   config.Get("DB_USERNAME", "postgres"),
				"password":   config.Get("DB_PASSWORD", "postgres"),
			},
			"redis": map[string]any{
				"driver":     "redis",
				"host":       facades.Config().Get("REDIS_HOST", "localhost"),
				"port":       facades.Config().Get("REDIS_PORT", 6379),
				"repository": facades.Config().Get("REDIS_DATABASE", "0"),
				"password":   facades.Config().Get("REDIS_PASSWORD", ""),
			},
		},
		"default":  config.Get("DB_CONNECTION", "postgresql"),
		"dir":      "./repository",
		"timezone": carbon.UTC,
	})
}
