package main

import (
	"reflect"
	"strconv"
	"strings"

	static "github.com/Code-Hex/echo-static"
	"github.com/datake914/monolog/controllers"
	"github.com/datake914/monolog/services"
	assetfs "github.com/elazarl/go-bindata-assetfs"
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
	e.Use(static.ServeRoot("/", NewAssets("frontend/dist")))
	e.Use(static.ServeRoot("/login", NewAssets("frontend/dist")))
	e.Use(static.ServeRoot("/js", NewAssets("frontend/dist/static/js")))
	e.Use(static.ServeRoot("/css", NewAssets("frontend/dist/static/css")))
	e.Use(static.ServeRoot("/fonts", NewAssets("frontend/dist/static/fonts")))

	e.Binder = &CustomBinder{}

	s := createLogService(opts.Service())
	e.GET("/api/logs", controllers.NewLogController(s).Index)
	e.GET("/api/logevents", controllers.NewLogEventController(s).Index)

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

func NewAssets(root string) *assetfs.AssetFS {
	return &assetfs.AssetFS{
		Asset:     Asset,
		AssetDir:  AssetDir,
		AssetInfo: AssetInfo,
		Prefix:    root,
	}
}
