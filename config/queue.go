package config

import "github.com/fwidjaya20/symphonic/facades"

func init() {
	config := facades.Config()

	config.Add("queue", map[string]any{
		"connections": map[string]any{
			"redis": map[string]any{
				"host":       facades.Config().Get("repository.connections.redis.host"),
				"port":       facades.Config().Get("repository.connections.redis.port"),
				"repository": facades.Config().Get("repository.connections.redis.repository"),
				"password":   facades.Config().Get("repository.connections.redis.password"),
			},
		},
		"default": "sync",
	})
}
