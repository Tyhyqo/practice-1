package app

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
	"practice/my_middleware"
)

type App struct {
	Server *echo.Echo
}

func NewApp() *App {
	app := &App{}
	app.Server = echo.New()
	return app
}

func RunApp(myApp *App) {
	myApp.Server.Logger.Fatal(myApp.Server.Start(":1323"))
}

func ConfigApp(myApp *App, db *gorm.DB) {
	myApp.Server.Use(middleware.Logger())
	myApp.Server.Use(middleware.Recover())
	myApp.Server.Use(my_middleware.DbMiddleware(db))
}
