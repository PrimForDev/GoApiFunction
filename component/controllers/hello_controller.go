package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// HelloHandler handles the /HelloFunction route
func HelloHandler(c *fiber.Ctx) error {
	message := "Pass a name in the query string for a personalized response.\n"
	name := c.Query("name")
	if name != "" {
		message = fmt.Sprintf("Hello, %s!\n", name)
	}
	return c.SendString(message)
}

// PostHandler handles POST requests to /HelloFunction
func PostHandler(c *fiber.Ctx) error {
	return c.SendString("POST request received!")
}
