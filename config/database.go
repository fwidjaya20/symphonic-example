package config

import (
	"github.com/fwidjaya20/symphonic/facades"
	"github.com/golang-module/carbon"
)

func init() {
	config := facades.Config()

	config.Add("database", map[string]any{
		"connections": map[string]any{
			"postgresql": map[string]any{
				"driver":   "postgresql",
				"host":     config.Get("DB_HOST", "http://localhost"),
				"port":     config.Get("DB_PORT", 5432),
				"database": config.Get("DB_DATABASE", "forge"),
				"username": config.Get("DB_USERNAME", "postgres"),
				"password": config.Get("DB_PASSWORD", "postgres"),
			},
		},
		"default":  config.Get("DB_CONNECTION", "postgresql"),
		"timezone": carbon.UTC,
	})
}
