package main

import (
	"clothing-store/backend/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())

	// Serve static files (images, etc.)
	app.Static("/assets", "./frontend/assets") // maps /assets/* to frontend/assets/*

	// Setup routes
	routes.SetupProductRoutes(app)

	//start server
	app.Listen(":3000")
}
