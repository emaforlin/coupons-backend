package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type msgBaseResponse struct {
	IsError bool   `json:"is_error,omitempty"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func responseMsg(c echo.Context, message string, responseCode int, data ...any) error {
	var msg msgBaseResponse

	if len(data) > 0 {
		msg.Data = data
	}
	msg.Message = message
	return c.JSON(responseCode, msg)
}

func errorMsg(c echo.Context, message string, responseCode ...int) error {
	var code int
	if len(responseCode) < 1 {
		code = http.StatusBadRequest
	} else {
		code = responseCode[0]
	}

	return c.JSON(code, &msgBaseResponse{
		IsError: true,
		Message: message,
	})
}
