package routes

import (
	"clothing-store/backend/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupProductRoutes(app *fiber.App) {
	api := app.Group("/products")
	api.Get("/", handlers.GetAllProducts)
	api.Get("/:id", handlers.GetProduct)
}
