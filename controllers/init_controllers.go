package controllers

import (
	echoSwagger "github.com/swaggo/echo-swagger"
	"practice/app"
	_ "practice/docs"
	"practice/my_middleware"
)

func InitControllers(myApp *app.App) {
	myApp.Server.POST("/register", RegisterUser)
	myApp.Server.POST("/login", LoginUser)
	myApp.Server.GET("/protected", Protected, my_middleware.JWTMiddleware)
	myApp.Server.POST("/logout", LogoutUser)

	myApp.Server.GET("/swagger/*", echoSwagger.WrapHandler)
}
