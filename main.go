package main

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/datake914/monolog/controllers"
	"github.com/datake914/monolog/services"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

const (
	awsCloudWatchLogs string = "cloudwatchlogs"
	awsS3             string = "s3"
	elasticsearch     string = "es"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORS())
	e.Binder = &CustomBinder{}

	s := createLogService(opts.Service())
	e.GET("/logs", controllers.NewLogController(s).Index)
	e.GET("/logevents", controllers.NewLogEventController(s).Index)

	e.Logger.Fatal(e.Start(":" + strconv.Itoa(opts.Port())))
}

func createLogService(service string) services.Service {
	switch service {
	case awsCloudWatchLogs:
		return new(services.CloudwatchService)
	default:
		return new(services.CloudwatchService)
	}
}

type CustomBinder struct{}

func (cb *CustomBinder) Bind(i interface{}, c echo.Context) (err error) {
	// Use default binder.
	db := new(echo.DefaultBinder)
	if err = db.Bind(i, c); err != nil {
		return
	}
	// Bind X-ServiceToken header.
	token := strings.Split(c.Request().Header.Get("X-ServiceToken"), " ")
	if len(token) == 2 && token[0] == "Logs" {
		f := reflect.ValueOf(i).Elem().FieldByName("SessionToken")
		f.Set(reflect.ValueOf(token[1]))
	}
	return nil
}
