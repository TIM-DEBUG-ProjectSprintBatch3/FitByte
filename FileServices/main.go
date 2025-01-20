package main

import (
	"os"
	"os/signal"
	"syscall"

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

	// Handle graceful shutdown
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sig
		os.Exit(0)
	}()

	//? Auto Migrate
	migrations.Migrate()

	server := httpServer.HttpServer{}
	server.Listen()

}
