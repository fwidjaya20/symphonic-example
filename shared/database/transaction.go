package database

import (
	"github.com/fwidjaya20/symphonic-skeleton/config"
	"github.com/fwidjaya20/symphonic-skeleton/shared/context"
	"github.com/labstack/echo/v4"
)

type TxCallback = func(e echo.Context) error

func RunInTransaction(c *context.SymphonicContext, callback TxCallback) error {
	tx := config.DB().Begin()

	defer func() {
		if r := recover(); nil != r {
			tx.Rollback()
		}
	}()

	c.Database = tx

	if err := callback(c); nil != err {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}
