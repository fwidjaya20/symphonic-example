package config

import (
	BootstrapEvent "github.com/fwidjaya20/symphonic-skeleton/bootstrap/event"
	BootstrapSchedule "github.com/fwidjaya20/symphonic-skeleton/bootstrap/schedule"
	"github.com/fwidjaya20/symphonic/contracts/foundation"
	"github.com/fwidjaya20/symphonic/database"
	"github.com/fwidjaya20/symphonic/event"
	"github.com/fwidjaya20/symphonic/facades"
	"github.com/fwidjaya20/symphonic/log"
	"github.com/fwidjaya20/symphonic/schedule"
	"github.com/golang-module/carbon"
)

func Boot() {
	// This function will fire all init() function under this directory
}

func init() {
	config := facades.Config()

	config.Add("app", map[string]any{
		"env":  config.Get("APP_ENV", "production"),
		"port": config.Get("APP_PORT", "8000"),
		"providers": []foundation.ServiceProvider{
			&event.ServiceProvider{},
			&database.ServiceProvider{},
			&log.ServiceProvider{},
			&schedule.ServiceProvider{},
			&BootstrapEvent.ServiceProvider{},
			&BootstrapSchedule.TaskSchedulerServiceProvider{},
		},
		"timezone": carbon.UTC,
	})
}
