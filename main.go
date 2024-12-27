package main

import (
	"log"
	"os"

	"secondhand-marketplace-api/config"
	"secondhand-marketplace-api/routes"
	
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDatabase()

	routes.SetupRoutes(app)

	err = app.Listen(":3000")
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
