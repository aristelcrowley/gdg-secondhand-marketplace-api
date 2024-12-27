package routes

import (
	"gdg-secondhand-marketplace-api/controllers"
	"gdg-secondhand-marketplace-api/middlewares"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/login", controllers.Login)

	protected := app.Group("/api", middlewares.CheckAuth)

	protected.Get("/users", controllers.GetUsers)  
	protected.Get("/user", controllers.GetUser)  
	protected.Put("/user", controllers.UpdateUser)
	protected.Delete("/user/:id", controllers.DeleteUser)  

	protected.Get("/categories", controllers.GetCategories)
	protected.Post("/categories", controllers.CreateCategory) 
	protected.Put("/categories/:id", controllers.UpdateCategory) 
	protected.Delete("/categories/:id", controllers.DeleteCategory) 

	
	protected.Get("/items", controllers.GetItems)
	protected.Post("/items", controllers.CreateItem)
	protected.Put("/items/:id", controllers.UpdateItem)  
	protected.Delete("/items/:id", controllers.DeleteItem


	protected.Get("/orders", controllers.GetOrders)
	protected.Post("/orders", controllers.CreateOrder)
	protected.Put("/orders/:id", controllers.UpdateOrder) 
	protected.Delete("/orders/:id", controllers.DeleteOrder) 
}
