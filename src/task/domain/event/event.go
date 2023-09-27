package event

import "github.com/labstack/echo/v4"

type TaskEvent interface {
	PublishTaskCreated(c echo.Context) error
}
