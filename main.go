package main

import (
	"github.com/JhonasMutton/izzi/internal/infra"
	"github.com/JhonasMutton/izzi/pkg"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	setupEnvVars()
	infra.SetupLogger()
}

func main() {
	app, e := setupApplication()
	if &app == nil || e != nil{
		infra.Logger.Fatal("Application setup error.")
	}
	app.SetupHealthCheck()
	routers := app.SetupRoutes()
	app.SetupServerApplication(routers)
}

func setupEnvVars() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file.")
	}
}

func setupApplication() (pkg.Application, error) {
	app, e := SetupApplication()
	if e != nil {
		infra.Logger.Fatal("Application setup error", "error", e.Error())
		panic(e)
	}
	return app, e
}
