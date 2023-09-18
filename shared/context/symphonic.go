package context

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type SymphonicContext struct {
	echo.Context
	Database *gorm.DB
}
