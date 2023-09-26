package config

import "github.com/fwidjaya20/symphonic/facades"

func init() {
	config := facades.Config()

	config.Add("queue", map[string]any{
		"connections": map[string]any{
			"redis": map[string]any{
				"host":     facades.Config().Get("database.connections.redis.host"),
				"port":     facades.Config().Get("database.connections.redis.port"),
				"database": facades.Config().Get("database.connections.redis.database"),
				"password": facades.Config().Get("database.connections.redis.password"),
			},
		},
		"default": "sync",
	})
}
