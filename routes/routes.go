package routes

import (
	"gdg-secondhand-marketplace-api/controllers"
	"gdg-secondhand-marketplace-api/middlewares"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/login", controllers.Login)

	protected := app.Group("/api", middlewares.Protect)

	protected.Post("/users", controllers.PostUser)  
	protected.Get("/users", controllers.GetUsers)  
	protected.Get("/users/:id", controllers.GetUserByID)  
	protected.Put("/users/:id", controllers.PutUser)
	protected.Delete("/users/:id", controllers.DeleteUser)  

	protected.Post("/categories", controllers.PostCategory) 
	protected.Get("/categories", controllers.GetCategories)
	protected.Put("/categories/:id", controllers.PutCategory) 
	protected.Delete("/categories/:id", controllers.DeleteCategory) 

	protected.Post("/items", controllers.PostItem)
	protected.Get("/items", controllers.GetItems)
	protected.Put("/items/:id", controllers.PutItem)  
	protected.Delete("/items/:id", controllers.DeleteItem)

	protected.Post("/orders", controllers.PostOrder)
	protected.Get("/orders", controllers.GetOrders)
	protected.Put("/orders/:id", controllers.PutOrder) 
	protected.Delete("/orders/:id", controllers.DeleteOrder) 
}