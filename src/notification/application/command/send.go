package command

import (
	"fmt"

	"github.com/fwidjaya20/symphonic-skeleton/shared/exception"
	"github.com/fwidjaya20/symphonic-skeleton/src/notification/application/public"
	"github.com/fwidjaya20/symphonic-skeleton/src/notification/constant"
	"github.com/fwidjaya20/symphonic-skeleton/src/notification/domain/entity"
	"github.com/fwidjaya20/symphonic-skeleton/src/notification/infrastructure/cloud_messaging"
	"github.com/fwidjaya20/symphonic/facades"
	"github.com/labstack/echo/v4"
)

type SendHandler struct{}

func (h SendHandler) Execute(c echo.Context, request public.SendNotificationRequest) error {
	var (
		err error
	)

	client := cloud_messaging.GetCloudMessaging(request.Channel)
	if nil == client {
		err = fmt.Errorf("channel '%s' was not supported", request.Channel.ToString())
		facades.Logger().Errorf("[Notification] get cloud messaging client error: %v", err.Error())
		return exception.New(err, constant.ErrClientNotFound, err.Error(), nil)
	}

	if err = client.Send(entity.Message{
		Recipient: request.To,
		Subject:   request.Subject,
		Body:      request.Body,
	}); nil != err {
		facades.Logger().Errorf("[Notification] cloud messaging send new message error: %v", err.Error())
		return exception.New(err, constant.ErrSendMessage, err.Error(), nil)
	}

	return nil
}

func NewSendHandler() SendHandler {
	return SendHandler{}
}
