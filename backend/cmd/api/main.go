package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	"github.com/AshishBodhankar/clothing-store/backend/internal/middleware"
	"github.com/AshishBodhankar/clothing-store/backend/internal/product"
	"github.com/AshishBodhankar/clothing-store/backend/internal/user"
)

func main() {
	// Load environment variables or configuration
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		log.Fatal("JWT_SECRET environment variable not set")
	}

	// Initialize database connection
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize Gin router
	r := gin.Default()

	// Apply global middleware
	r.Use(middleware.ErrorHandler())

	// Register routes
	user.RegisterRoutes(r, db, secretKey)
	product.RegisterRoutes(r, db, secretKey)

	// Start the server
	if err := r.Run(); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
