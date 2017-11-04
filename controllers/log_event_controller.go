package controllers

import (
	"github.com/datake914/monolog/services"
	"github.com/labstack/echo"
)

// LogEventController provides the methods for log event resource.
type LogEventController struct {
	baseController
}

// NewLogEventController is constructor of LogEventController.
func NewLogEventController(service services.Service) *LogEventController {
	return &LogEventController{
		baseController{service: service},
	}
}

// Index list the log events.
func (controller *LogEventController) Index(c echo.Context) error {
	// Bind request parameter.
	req := new(services.SearchLogEventRequest)
	if err := c.Bind(req); err != nil {
		return controller.render400(c)
	}
	// Execute serivce.
	res, err := controller.service.SearchLogEvents(req)
	if err != nil {
		return controller.handleError(c, err)
	}
	return controller.render200(c, res)
}
