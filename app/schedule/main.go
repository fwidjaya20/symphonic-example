package main

import (
	"github.com/fwidjaya20/symphonic-example/bootstrap"
	"github.com/fwidjaya20/symphonic/facades"
)

func main() {
	bootstrap.Boot()

	facades.App().GetSchedule().Run()

	select {}
}
