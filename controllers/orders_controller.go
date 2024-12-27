package controllers

import (
	"gdg-secondhand-marketplace-api/config"
	"gdg-secondhand-marketplace-api/middlewares"
	"gdg-secondhand-marketplace-api/models"

	"github.com/gofiber/fiber/v2"
)

func PostOrder(c *fiber.Ctx) error {
	var order models.Order
	if err := c.BodyParser(&order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse the body",
		})
	}

	userID := c.Locals("userID").(int) 

	if order.UserID == 0 || order.ItemID == 0{
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "You need to input the user_id and item_id",
		})
	}


	if order.UserID != userID && !middlewares.IsAdmin(c) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "You can only create orders for yourself",
		})
	}

	if err := config.DB.Create(&order).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error creating order",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(order)
}

func GetOrders(c *fiber.Ctx) error {
	userID := c.Locals("userID").(int)

	if !middlewares.IsAdmin(c) {
		var orders []models.Order
		if err := config.DB.Where("user_id = ?", userID).Find(&orders).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Error fetching orders",
			})
		}
		return c.JSON(orders)
	}

	var orders []models.Order
	if err := config.DB.Find(&orders).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error fetching orders",
		})
	}
	return c.JSON(orders)
}

func PutOrder(c *fiber.Ctx) error {
	orderID := c.Params("id")
	var order models.Order
	if err := config.DB.First(&order, orderID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Order not found",
		})
	}

	userID := c.Locals("userID").(int)

	if order.UserID != userID && !middlewares.IsAdmin(c) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "You can only modify your own orders or be an admin",
		})
	}

	if err := c.BodyParser(&order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse the body",
		})
	}

	if err := config.DB.Save(&order).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error updating order",
		})
	}
	return c.JSON(order)
}

func DeleteOrder(c *fiber.Ctx) error {
	orderID := c.Params("id")
	var order models.Order
	if err := config.DB.First(&order, orderID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Order not found",
		})
	}

	userID := c.Locals("userID").(int) 
	if order.UserID != userID {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "You can only delete your own orders",
		})
	}

	if err := config.DB.Delete(&order).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error deleting order",
		})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
