package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/primfordev/goapi/controllers" // Update "yourusername/project" with your module path
)

// SetupRoutes sets up the API routes
func SetupRoutes(app *fiber.App) {
	// Create a route group for API endpoints
	api := app.Group("/api")

	// Define routes under the /api group
	api.Get("/HelloFunction", controllers.HelloHandler)
	api.Get("/UserFunction", controllers.UserHandler)
}
