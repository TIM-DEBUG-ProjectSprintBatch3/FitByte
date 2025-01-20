package main

import (
	"fmt"
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
	fmt.Printf("Load ENV\n")
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	fmt.Printf("DI Healthcheck\n")
	di.HealthCheck()

	// Handle graceful shutdown
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sig
		os.Exit(0)
	}()

	//? Auto Migrate
	fmt.Printf("Migrate\n")
	migrations.Migrate()

	fmt.Printf("Start Server\n")
	server := httpServer.HttpServer{}
	server.Listen()

}
