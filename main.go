package main

import (
	"log"

	"gdg-secondhand-marketplace-api/config"
	"gdg-secondhand-marketplace-api/routes"
	
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
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
