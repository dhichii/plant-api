package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type base struct {
	Code    int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type baseWithoutData struct {
	Code    int    `json:"status"`
	Message string `json:"message"`
}

type Reason struct {
	Reason interface{} `json:"reason"`
}

func CreateResponse(c echo.Context, httpCode int, data interface{}) error {
	response := &base{
		Code:    httpCode,
		Message: http.StatusText(httpCode),
		Data:    data,
	}
	return c.JSON(httpCode, *response)
}

func CreateWithoutDataResponse(c echo.Context, httpCode int) error {
	response := &baseWithoutData{
		Code:    httpCode,
		Message: http.StatusText(httpCode),
	}
	return c.JSON(httpCode, *response)
}
