package controllers

import (
	"gdg-secondhand-marketplace-api/config"
	"gdg-secondhand-marketplace-api/models"
	"gdg-secondhand-marketplace-api/middlewares"

	"github.com/gofiber/fiber/v2"
)

func PostItem(c *fiber.Ctx) error {
	var item models.Item
	if err := c.BodyParser(&item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse the body",
		})
	}

	userID := c.Locals("userID").(int) 

	if item.UserID == 0 || item.CategoryID == 0{
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "You need to input the user_id and category_id",
		})
	}

	if item.UserID != userID && !middlewares.IsAdmin(c) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "You can only create items for yourself",
		})
	}

	if err := config.DB.Create(&item).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error creating item",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(item)
}

func GetItems(c *fiber.Ctx) error {
	userID := c.Locals("userID").(int)

	if !middlewares.IsAdmin(c) {
		var items []models.Item
		if err := config.DB.Where("user_id = ?", userID).Find(&items).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Error fetching items",
			})
		}
		return c.JSON(items)
	}

	var items []models.Item
	if err := config.DB.Find(&items).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error fetching items",
		})
	}
	return c.JSON(items)
}

func PutItem(c *fiber.Ctx) error {
	itemID := c.Params("id")
	var item models.Item
	if err := config.DB.First(&item, itemID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Item not found",
		})
	}

	userID := c.Locals("userID").(int)

	if item.UserID != userID && !middlewares.IsAdmin(c) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "You can only modify your own items or be an admin",
		})
	}

	if err := c.BodyParser(&item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse the body",
		})
	}

	if err := config.DB.Save(&item).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error updating item",
		})
	}
	return c.JSON(item)
}

func DeleteItem(c *fiber.Ctx) error {
	itemID := c.Params("id")
	var item models.Item
	if err := config.DB.First(&item, itemID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Item not found",
		})
	}

	userID := c.Locals("userID").(int)

	if item.UserID != userID && !middlewares.IsAdmin(c) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "You can only delete your own items or be an admin",
		})
	}

	if err := config.DB.Delete(&item).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error deleting item",
		})
	}
	return c.SendString("Item deleted")
}
