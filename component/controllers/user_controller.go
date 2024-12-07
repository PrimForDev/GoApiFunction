package controllers

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/primfordev/goapi/database"
	"github.com/primfordev/goapi/models"
	"gorm.io/gorm"
)

// Global variable to store database connections
var db *gorm.DB

func IsNullOrEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

func GetUserHandler(c *fiber.Ctx) error {
	// Get id from Query Parameter
	userID := c.Query("id")

	// If id from Query Parameter is null
	if IsNullOrEmpty(userID) {
		var user models.User
		// Try to get the id from Body.
		if err := c.BodyParser(&user); err != nil {
			log.Println("No ID provided in Query or Body. Listing all users.")
		} else {
			// Convert user.ID (int) to string.
			userID = strconv.Itoa(user.ID)
		}
	}

	// If no id is specified (in both Query and Body)
	if IsNullOrEmpty(userID) {
		// Fetch all user entries from the database.
		var users []models.User
		db, err := database.ConnectToAzureSQL()
		if err != nil {
			log.Println("Error connecting to database:", err)
			return c.Status(500).SendString(fmt.Sprintf("Connection error: %v", err))
		}

		result := db.Find(&users)
		if result.Error != nil {
			return c.Status(500).SendString(fmt.Sprintf("Failed to retrieve users: %v", result.Error))
		}

		// Return all user data.
		return c.Status(200).JSON(users)
	}

	// If id exists, search for user in database by id.
	var user models.User
	db, err := database.ConnectToAzureSQL()
	if err != nil {
		log.Println("Error connecting to database:", err)
		return c.Status(500).SendString(fmt.Sprintf("Connection error: %v", err))
	}

	result := db.First(&user, userID)
	if result.Error != nil {
		// Check if the user is not found in the database.
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(404).SendString("User not found")
		}
		// Other cases such as database connection issues
		return c.Status(500).SendString(fmt.Sprintf("Failed to retrieve user: %v", result.Error))
	}

	// Return user data in JSON format.
	return c.Status(200).JSON(user)
}

func CreateUserHandler(c *fiber.Ctx) error {
	// Get values ​​from Query parameters
	name := c.Query("name", "")
	email := c.Query("email", "")
	firstName := c.Query("firstName", "")
	lastName := c.Query("lastName", "")
	role := c.Query("role", "")

	// Check the value with the IsNullOrEmpty function.
	if IsNullOrEmpty(name) || IsNullOrEmpty(email) || IsNullOrEmpty(firstName) || IsNullOrEmpty(lastName) || IsNullOrEmpty(role) {
		var user models.User
		if err := c.BodyParser(&user); err != nil {
			return c.Status(400).SendString(fmt.Sprintf("Failed to parse request body: %v", err))
		}
		// Use value from body if query is null
		name = user.Name
		email = user.Email
		firstName = user.FirstName
		lastName = user.LastName
		role = user.Role
	}

	// Create a new user.
	newUser := models.User{
		Name:      name,
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		Role:      role,
	}

	// Save data to database.
	result := db.Create(&newUser)
	if result.Error != nil {
		return c.Status(500).SendString(fmt.Sprintf("Failed to insert user: %v", result.Error))
	}

	return c.Status(201).JSON(newUser)
}

func UpdateUserHandler(c *fiber.Ctx) error {
	// Get id from Query Parameter
	userID := c.Query("id")

	// If id from Query Parameter is null
	if IsNullOrEmpty(userID) {
		var user models.User
		// Try to get the id from Body.
		if err := c.BodyParser(&user); err != nil {
			return c.Status(400).SendString(fmt.Sprintf("Failed to parse request body: %v", err))
		}
		// Check if id in Body is null.
		if user.ID != 0 {
			userID = strconv.Itoa(user.ID)
		}
	}

	// Check if id exists
	if IsNullOrEmpty(userID) {
		return c.Status(400).SendString("User ID is required")
	}
	log.Printf(`user id: %s`, userID)
	// Get new data from request body
	var updatedUser models.User
	if err := c.BodyParser(&updatedUser); err != nil {
		return c.Status(400).SendString(fmt.Sprintf("Failed to parse request body: %v", err))
	}

	// Connect database
	db, err := database.ConnectToAzureSQL()
	if err != nil {
		log.Println("Error connecting to database:", err)
		return c.Status(500).SendString(fmt.Sprintf("Connection error: %v", err))
	}

	// Find a user in the database using their id.
	var user models.User
	result := db.First(&user, userID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(404).SendString("User not found")
		}
		return c.Status(500).SendString(fmt.Sprintf("Failed to retrieve user: %v", result.Error))
	}

	// Update data only in fields that receive new values.
	if !IsNullOrEmpty(updatedUser.Name) {
		user.Name = updatedUser.Name
	}
	if !IsNullOrEmpty(updatedUser.Email) {
		user.Email = updatedUser.Email
	}
	if !IsNullOrEmpty(updatedUser.FirstName) {
		user.FirstName = updatedUser.FirstName
	}
	if !IsNullOrEmpty(updatedUser.LastName) {
		user.LastName = updatedUser.LastName
	}
	if !IsNullOrEmpty(updatedUser.Role) {
		user.Role = updatedUser.Role
	}

	// Save updated data in database.
	if err := db.Save(&user).Error; err != nil {
		return c.Status(500).SendString(fmt.Sprintf("Failed to update user: %v", err))
	}

	// Return the result.
	return c.Status(200).JSON(user)
}

func DeleteUserHandler(c *fiber.Ctx) error {
	// Get id from Query Parameter
	userID := c.Query("id")

	// If id from Query Parameter is null
	if IsNullOrEmpty(userID) {
		var user models.User
		// Try to get the id from Body.
		if err := c.BodyParser(&user); err != nil {
			return c.Status(400).SendString(fmt.Sprintf("Failed to parse request body: %v", err))
		}
		// Check if id in Body is null.
		if user.ID != 0 {
			userID = strconv.Itoa(user.ID)
		}
	}

	// Check if id exists
	if IsNullOrEmpty(userID) {
		return c.Status(400).SendString("User ID is required")
	}
	log.Printf("User ID received: %s", userID)
	fmt.Printf("User ID received: %s", userID)
	var user models.User
	result := db.First(&user, userID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(404).SendString("User not found")
		}
		return c.Status(500).SendString(fmt.Sprintf("Failed to retrieve user: %v", result.Error))
	}

	// Delete user from database
	if err := db.Delete(&user).Error; err != nil {
		return c.Status(500).SendString(fmt.Sprintf("Failed to delete user: %v", err))
	}

	// Return the result that the user was deleted.
	return c.Status(200).SendString("User deleted successfully")
}
