package controllers

import (
	"time"
	
	"gdg-secondhand-marketplace-api/config"
	"gdg-secondhand-marketplace-api/models"
	
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("secret_key") // secret key for JWT signing

func Login(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).SendString("Invalid request")
	}

	var dbUser models.User
	if err := config.DB.Where("email = ?", user.Email).First(&dbUser).Error; err != nil {
		return c.Status(401).SendString("Invalid credentials")
	}

	if dbUser.Password != user.Password {
		return c.Status(401).SendString("Invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": dbUser.UserID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return c.Status(500).SendString("Error generating token")
	}

	return c.JSON(fiber.Map{
		"token": tokenString,
	})
}
