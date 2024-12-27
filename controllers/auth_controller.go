package controllers

import (
	"os"
	"time"

	"gdg-secondhand-marketplace-api/models"
	"gdg-secondhand-marketplace-api/config"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Login(c *fiber.Ctx) error {
	var creds models.User
	if err := c.BodyParser(&creds); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse the body: " + err.Error(),
		})
	}

	var user models.User
	if err := config.DB.Where("email = ?", creds.Email).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	if user.Password != creds.Password {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	claims := jwt.MapClaims{
		"sub":  user.ID,         
		"role": user.Role,        
		"exp":  time.Now().Add(time.Hour * 24).Unix(), 
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY"))) 
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error generating token",
		})
	}

	return c.JSON(fiber.Map{
		"token": tokenString,
	})
}
