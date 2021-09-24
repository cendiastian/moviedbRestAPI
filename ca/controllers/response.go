package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Meta struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}
	Data interface{} `json:"data"`
}
type UpdateResponse struct {
	Meta struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}
}

func NewSuccesResponse(c echo.Context, data interface{}) error {
	response := Response{}
	response.Meta.Status = http.StatusOK
	response.Meta.Message = "Success"
	response.Data = data
	return c.JSON(http.StatusOK, response)
}
func UpdateSuccesResponse(c echo.Context, message string) error {
	response := UpdateResponse{}
	response.Meta.Status = http.StatusOK
	response.Meta.Message = message
	return c.JSON(http.StatusOK, response)
}

func NewErrorResponse(c echo.Context, status int, err error) error {
	response := Response{}
	response.Meta.Status = status
	response.Meta.Message = err.Error()
	response.Data = nil
	return c.JSON(status, response)
}
