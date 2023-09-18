package http

import (
	"net/http"

	"github.com/fwidjaya20/symphonic/facades"
	"github.com/labstack/echo/v4"
)

type Kernel struct {
}

func (kernel *Kernel) Routes(e *echo.Echo) {
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, facades.Config().Inspect())
	})
}
