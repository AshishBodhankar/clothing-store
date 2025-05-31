package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/AshishBodhankar/clothing-store/backend/internal/middleware"
	"github.com/AshishBodhankar/clothing-store/backend/internal/product"
	"github.com/AshishBodhankar/clothing-store/backend/internal/user"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load environment variables
	databaseURL := os.Getenv("DATABASE_URL")
	jwtSecret := os.Getenv("JWT_SECRET")
	if databaseURL == "" || jwtSecret == "" {
		log.Fatal("Environment variables DATABASE_URL or JWT_SECRET not set")
	}

	// Initialize database connection
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize Gin router
	r := gin.Default()

	// Apply global middleware
	r.Use(middleware.ErrorHandler())

	// Initialize and register user routes
	userService := user.NewService(db, jwtSecret)
	userHandler := user.NewHandler(userService)
	userHandler.RegisterRoutes(r)

	// Register product routes
	product.RegisterRoutes(r, db)

	// Start the server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
