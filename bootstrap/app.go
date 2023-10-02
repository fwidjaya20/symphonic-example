package bootstrap

import (
	"github.com/fwidjaya20/symphonic-example/config"
	"github.com/fwidjaya20/symphonic/foundation"
)

func Boot() {
	app := foundation.NewApplication()

	config.Boot()
	app.Boot()
}
