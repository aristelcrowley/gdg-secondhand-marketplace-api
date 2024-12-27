package controllers

import (
	"strconv"
	"gdg-secondhand-marketplace-api/models"
	"gdg-secondhand-marketplace-api/config"
	"gdg-secondhand-marketplace-api/middlewares"

	"github.com/gofiber/fiber/v2"
)

func PostUser(c *fiber.Ctx) error {
	if !middlewares.IsAdmin(c) {
		return c.Status(fiber.StatusUnauthorized).SendString("You are not authorized to create users")
	}

	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).SendString("Invalid request")
	}

	if err := config.DB.Create(&user).Error; err != nil {
		return c.Status(500).SendString("Error creating user")
	}

	return c.Status(201).JSON(user)
}

func GetUsers(c *fiber.Ctx) error {
	if !middlewares.IsAdmin(c) {
		return c.Status(fiber.StatusUnauthorized).SendString("You are not authorized to view all users")
	}

	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		return c.Status(500).SendString("Error fetching users")
	}
	return c.JSON(users)
}

func GetUserByID(c *fiber.Ctx) error {
	
	id := c.Params("id")
	var user models.User

	userID := c.Locals("userID").(int)

	if id != strconv.Itoa(userID) && !middlewares.IsAdmin(c) {
		return c.Status(fiber.StatusUnauthorized).SendString("You are not authorized to view this user's profile")
	}
 
	if err := config.DB.First(&user, id).Error; err != nil {
		return c.Status(404).SendString("User not found")
	}
	return c.JSON(user)
}

func PutUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return c.Status(404).SendString("User not found")
	}

	userID := c.Locals("userID").(int)

	if id != strconv.Itoa(userID) && !middlewares.IsAdmin(c) {
		return c.Status(403).SendString("You are not authorized to update this user's profile")
	}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).SendString("Invalid request")
	}

	if err := config.DB.Model(&user).Where("user_id = ?", id).Updates(user).Error; err != nil {
		return c.Status(500).SendString("Error updating user")
	}

	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	if !middlewares.IsAdmin(c) {
		return c.Status(fiber.StatusUnauthorized).SendString("You are not authorized to delete users")
	}

	id := c.Params("id") 
	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		return c.Status(404).SendString("User not found")
	}

	return c.SendString("User deleted")
}
