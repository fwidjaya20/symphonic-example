package bootstrap

import (
	"github.com/fwidjaya20/symphonic-skeleton/config"
	"github.com/fwidjaya20/symphonic/foundation"
)

func Boot() {
	app := foundation.NewApplication()

	config.Boot()
	app.Boot()
}
