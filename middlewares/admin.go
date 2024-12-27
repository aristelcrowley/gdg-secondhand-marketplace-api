package middlewares

import "github.com/gofiber/fiber/v2"

func IsAdmin(c *fiber.Ctx) bool {
	role, ok := c.Locals("role").(string)
	if !ok || role == "" {
		return false
	}
	return role == "admin"
}
