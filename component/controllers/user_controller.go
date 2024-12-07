package controllers

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/primfordev/goapi/database"
	"github.com/primfordev/goapi/models"
)

// UserHandler handles routes with dynamic user ID
func UserHandler(c *fiber.Ctx) error {
	// Set default message
	message := "Pass an id in the query string for a personalized response.\n"
	action := c.Query("action")

	if action != "" {
		message = fmt.Sprintf("User action: %s", action)

		// Check if the action is "listuser"
		if action == "listuser" {
			// Connect to Azure SQL
			db, err := database.ConnectToAzureSQL()
			if err != nil {
				// If a database connection error occurs, send an appropriate message.
				log.Println("Error connecting to database:", err)
				return c.Status(500).SendString(fmt.Sprintf("Connection error: %v", err))
			}

			// For example, fetch user data from a database (this should be customized according to your database structure).
			var users []models.User // Suppose you have a User structure in your database.
			result := db.Find(&users)
			if result.Error != nil {
				return c.Status(500).SendString(fmt.Sprintf("Failed to retrieve users: %v", result.Error))
			}

			// Send user data back to the client.
			return c.Status(200).JSON(users) // Submit user data in JSON format

		}
	}

	// Send the initially specified message
	return c.SendString(message)
}
