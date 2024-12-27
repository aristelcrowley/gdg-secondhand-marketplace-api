package controllers

import (
	"gdg-secondhand-marketplace-api/config"
	"gdg-secondhand-marketplace-api/models"
	"gdg-secondhand-marketplace-api/middlewares"
	
	"github.com/gofiber/fiber/v2"
)

func PostCategory(c *fiber.Ctx) error {
	if !middlewares.IsAdmin(c) {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "You must be an admin to create a category",
		})
	}

	var category models.Category
	if err := c.BodyParser(&category); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse the body",
		})
	}

	if err := config.DB.Create(&category).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error creating category",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(category)
}

func GetCategories(c *fiber.Ctx) error {
	var categories []models.Category
	if err := config.DB.Find(&categories).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error fetching categories",
		})
	}
	return c.JSON(categories)
}

func PutCategory(c *fiber.Ctx) error {
	if !middlewares.IsAdmin(c) {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "You must be an admin to update a category",
		})
	}

	categoryID := c.Params("id")
	var category models.Category
	if err := config.DB.First(&category, categoryID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Category not found",
		})
	}

	if err := c.BodyParser(&category); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse the body",
		})
	}

	if err := config.DB.Save(&category).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error updating category",
		})
	}
	return c.JSON(category)
}

func DeleteCategory(c *fiber.Ctx) error {
	if !middlewares.IsAdmin(c) {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "You must be an admin to delete a category",
		})
	}

	categoryID := c.Params("id")
	var category models.Category
	if err := config.DB.Delete(&category, categoryID).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error deleting category",
		})
	}
	return c.SendString("Category deleted")
}
