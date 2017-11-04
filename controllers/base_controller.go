package controllers

import (
	"net/http"

	"github.com/datake914/monolog/services"
	"github.com/labstack/echo"
)

// ResponseBody is http response body.
type ResponseBody struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type baseController struct {
	service services.Service
}

func (bc *baseController) render(c echo.Context, code int, body ResponseBody) error {
	return c.JSON(code, body)
}

func (bc *baseController) renderOK(c echo.Context, code int, data interface{}) error {
	return bc.render(c, code, ResponseBody{
		Success: true,
		Data:    data,
	})
}

func (bc *baseController) renderError(c echo.Context, code int, message string) error {
	return bc.render(c, code, ResponseBody{
		Success: false,
		Message: message,
	})
}

func (bc *baseController) render200(c echo.Context, data interface{}) error {
	return bc.renderOK(c, http.StatusOK, data)
}

func (bc *baseController) render400(c echo.Context) error {
	return bc.renderError(c, http.StatusBadRequest, "Bad request.")
}

func (bc *baseController) render401(c echo.Context) error {
	return bc.renderError(c, http.StatusUnauthorized, "Unauthorized.")
}

func (bc *baseController) render405(c echo.Context) error {
	return bc.renderError(c, http.StatusMethodNotAllowed, "Method not allowed.")
}

func (bc *baseController) render500(c echo.Context) error {
	return bc.renderError(c, http.StatusInternalServerError, "Internal server error.")
}

func (bc *baseController) handleError(c echo.Context, err error) error {
	if serviceErr, ok := err.(services.ServiceError); ok {
		if serviceErr.UnAuthorized() {
			return bc.render401(c)
		}
	}
	return bc.render500(c)
}
