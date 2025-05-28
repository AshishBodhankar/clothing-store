package main

import (
	"clothing-store/backend/routes"
	"clothing-store/backend/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	utils.InitDB()

	if utils.DB != nil {
		defer utils.DB.Close() // Close the DB connection pool when the app exits
	}

	app := fiber.New()

	app.Use(logger.New())

	// Serve static files (images, etc.)
	app.Static("/assets", "./frontend/assets") // maps /assets/* to frontend/assets/*

	// Setup routes
	routes.SetupProductRoutes(app)

	//start server
	app.Listen(":3000")
}
