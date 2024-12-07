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
	// กำหนดข้อความเริ่มต้น
	message := "Pass an id in the query string for a personalized response.\n"
	action := c.Query("action")

	if action != "" {
		message = fmt.Sprintf("User action: %s", action)

		// ตรวจสอบ action ว่าเป็น "listuser" หรือไม่
		if action == "listuser" {
			// เชื่อมต่อกับ Azure SQL
			db, err := database.ConnectToAzureSQL()
			if err != nil {
				// หากเกิดข้อผิดพลาดในการเชื่อมต่อฐานข้อมูล ส่งข้อความที่เหมาะสม
				log.Println("Error connecting to database:", err)
				return c.Status(500).SendString(fmt.Sprintf("Connection error: %v", err))
			}

			// คุณสามารถใช้ db เพื่อดึงข้อมูลจากฐานข้อมูล
			// ตัวอย่างเช่น ดึงข้อมูลผู้ใช้จากฐานข้อมูล (ควรปรับตามโครงสร้างฐานข้อมูลของคุณ)
			var users []models.User // สมมติว่าคุณมีโครงสร้าง User ในฐานข้อมูล
			result := db.Find(&users)
			if result.Error != nil {
				return c.Status(500).SendString(fmt.Sprintf("Failed to retrieve users: %v", result.Error))
			}

			// ส่งข้อมูลผู้ใช้กลับไปยัง client
			return c.Status(200).JSON(users) // ส่งข้อมูลผู้ใช้ในรูปแบบ JSON

		}
	}

	// ส่งข้อความที่กำหนดไว้ในตอนแรก
	return c.SendString(message)
}
