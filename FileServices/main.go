package main

import (
	"github.com/TimDebug/FitByte/src/database/migrations"
	"github.com/TimDebug/FitByte/src/di"
	httpServer "github.com/TimDebug/FitByte/src/http"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	err := godotenv.Load()
	di.HealthCheck()
	if err != nil {
		panic(err)
	}

	//? Auto Migrate
	migrations.Migrate()

	server := httpServer.HttpServer{}
	server.Listen()

}
