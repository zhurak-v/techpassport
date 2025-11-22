package main

import (
	"log"

	"github.com/zhurak-v/techpassport/auth-service/src/infrastructure/application"
)

func main() {
	app, err := application.InitApplication()
	if err != nil {
		log.Fatalf("failed to initialize app: %v", err)
	}

	if app == nil {
		log.Fatal("Failed to initialize application")
	}

	if err := app.Run(":1111"); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
