package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/primfordev/goapi/routes"
)

func main() {
	// Create a new Fiber app
	app := fiber.New()

	// Setup routes
	routes.SetupRoutes(app)

	// Set the address for listening
	listenAddr := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}
	log.Printf("Starting server on %s...", listenAddr)

	// Start the Fiber app
	if err := app.Listen(listenAddr); err != nil {
		log.Fatal(err)
	}
}
