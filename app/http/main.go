package main

import (
	"fmt"

	"github.com/fwidjaya20/symphonic-skeleton/bootstrap"
	"github.com/fwidjaya20/symphonic-skeleton/bootstrap/http"
	SharedContext "github.com/fwidjaya20/symphonic-skeleton/shared/context"
	"github.com/fwidjaya20/symphonic/facades"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	bootstrap.Boot()

	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Gzip())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(middleware.SecureWithConfig(middleware.SecureConfig{
		Skipper:               nil,
		XSSProtection:         "1; mode=block",
		ContentTypeNosniff:    "nosniff",
		XFrameOptions:         "DENY",
		HSTSMaxAge:            31536000,
		HSTSExcludeSubdomains: true,
		ContentSecurityPolicy: "default-src 'self'; script-src 'self' 'unsafe-inline'; style-src 'self' 'unsafe-inline'; img-src *; font-src 'self'; frame-ancestors 'none'; base-uri 'self'; form-action 'self'",
		CSPReportOnly:         false,
		HSTSPreloadEnabled:    true,
		ReferrerPolicy:        "same-origin",
	}))

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			sc := &SharedContext.SymphonicContext{
				Context: c,
			}
			return next(sc)
		}
	})

	kernel := http.Kernel{}

	kernel.Routes(e)

	if err := e.Start(fmt.Sprintf(":%v", facades.Config().Get("app.port"))); nil != err {
		e.Logger.Fatal("shutting down the server")
	}
}
