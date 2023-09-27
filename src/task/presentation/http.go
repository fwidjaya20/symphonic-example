package presentation

import (
	"errors"
	"net/http"

	"github.com/fwidjaya20/symphonic-skeleton/ioc"
	"github.com/fwidjaya20/symphonic-skeleton/shared/context"
	"github.com/fwidjaya20/symphonic-skeleton/shared/database"
	"github.com/fwidjaya20/symphonic-skeleton/shared/exception"
	"github.com/fwidjaya20/symphonic-skeleton/shared/vo"
	"github.com/fwidjaya20/symphonic-skeleton/src/task/application/public"
	"github.com/fwidjaya20/symphonic-skeleton/src/task/constant"
	"github.com/labstack/echo/v4"
)

func HttpVer1(e *echo.Echo) {
	e.POST("/todo", createTodo)
}

func createTodo(c echo.Context) error {
	var (
		err      error
		response *public.TaskResponse
		request  public.CreateTaskRequest
	)

	if err = c.Bind(&request); nil != err {
		return c.JSON(http.StatusBadRequest, vo.Reject{
			Code:    constant.ErrInvalidPayload,
			Message: err.Error(),
		})
	}

	err = database.RunInTransaction(c.(*context.SymphonicContext), func(e echo.Context) error {
		response, err = ioc.Injector().Task.Create.Execute(c, request)
		return err
	})
	if nil != err {
		var ex *exception.Exception
		errors.As(err, &ex)
		return c.JSON(constant.ToHttpStatusCode(ex.Code), vo.Reject{
			Code:     ex.Code,
			Message:  ex.Message,
			Metadata: nil,
		})
	}

	return c.JSON(http.StatusOK, vo.Resolve{
		Data:     response,
		Message:  "task has been created successfully",
		Metadata: nil,
	})
}
