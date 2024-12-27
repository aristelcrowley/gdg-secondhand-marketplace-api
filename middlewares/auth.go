package middlewares

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Protect(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(401).SendString("Authorization token required")
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	jwtKey := []byte(os.Getenv("JWT_SECRET_KEY"))

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.ErrUnauthorized
		}
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return c.Status(401).SendString("Invalid token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := claims["sub"].(float64)
		c.Locals("userID", int(userID)) 
		c.Locals("role", claims["role"].(string)) 
	}

	return c.Next()
}
