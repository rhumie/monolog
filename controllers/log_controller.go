package controllers

import (
	"github.com/datake914/monolog/services"
	"github.com/labstack/echo"
)

// LogController provides the methods for log resource.
type LogController struct {
	baseController
}

// NewLogController is constructor of LogController.
func NewLogController(service services.Service) *LogController {
	return &LogController{
		baseController{service: service},
	}
}

// Index list the logs.
func (controller *LogController) Index(c echo.Context) error {
	// Bind request parameter.
	req := new(services.SearchLogRequest)
	if err := c.Bind(req); err != nil {
		return controller.render400(c)
	}
	// Execute serivce.
	res, err := controller.service.SearchLogs(req)
	if err != nil {
		return controller.handleError(c, err)
	}
	return controller.render200(c, res)
}
