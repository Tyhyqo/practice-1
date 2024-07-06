package main

import (
	"log"
	"practice/app"
	"practice/controllers"
	"practice/migrations"
	"practice/repository"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := repository.OpenDB()
	migrations.OnAutoMigrate(db)

	myApp := app.NewApp()
	app.ConfigApp(myApp, db)

	controllers.InitControllers(myApp)
	app.RunApp(myApp)
}
