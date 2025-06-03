package main

import (
	"database/sql"
	"log"
	"os"
	"path/filepath" // For absolute path logging

	"github.com/AshishBodhankar/clothing-store/backend/internal/middleware"
	"github.com/AshishBodhankar/clothing-store/backend/internal/product"
	"github.com/AshishBodhankar/clothing-store/backend/internal/user"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// Log current working directory
	wd, err := os.Getwd()
	if err != nil {
		log.Printf("Warning: Could not get working directory: %v", err)
	} else {
		log.Printf("Current working directory: %s", wd)
	}

	envPath := "backend/.env"
	absEnvPath, _ := filepath.Abs(envPath)
	log.Printf("Attempting to load .env file from: %s (absolute: %s)", envPath, absEnvPath)

	// Load environment variables from .env file
	err = godotenv.Load(envPath) // Use the variable
	if err != nil {
		// Log the error but don't make it fatal yet, to see if os.Getenv picks up anything
		log.Printf("Warning: Error loading .env file (%s): %v. Continuing to check os.Getenv...", envPath, err)
	} else {
		log.Printf(".env file loaded successfully from %s", envPath)
	}

	// Retrieve environment variables
	databaseURL := os.Getenv("DATABASE_URL")
	jwtSecret := os.Getenv("JWT_SECRET")

	log.Printf("Retrieved DATABASE_URL: '%s'", databaseURL)
	log.Printf("Retrieved JWT_SECRET: '%s'", jwtSecret)

	if databaseURL == "" || jwtSecret == "" {
		log.Fatalf("Critical: Environment variables DATABASE_URL ('%s') or JWT_SECRET ('%s') not set or empty. Exiting.", databaseURL, jwtSecret)
	} else {
		log.Println("DATABASE_URL and JWT_SECRET are set.")
	}

	// Initialize database connection
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()
	log.Println("Database connection opened successfully.")

	// Initialize Gin router
	r := gin.Default()

	// Apply global middleware
	r.Use(middleware.ErrorHandler())

	// Initialize and register user routes
	log.Printf("Initializing user repository...")
	userRepo := user.NewRepository(db)
	log.Printf("Initializing user service...")
	// userService no longer takes jwtSecret directly, it's read from env by GenerateJWT
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)
	log.Printf("Registering user routes with JWT_SECRET: '%s'", jwtSecret)
	userHandler.RegisterRoutes(r, jwtSecret) // Pass jwtSecret here
	log.Println("User routes registered.")

	// Register product routes
	// product.RegisterRoutes needs the jwtSecret for its own JWT middleware setup
	log.Printf("Registering product routes with JWT_SECRET: '%s'", jwtSecret)
	product.RegisterRoutes(r, db, jwtSecret)
	log.Println("Product routes registered.")

	// Start the server
	log.Println("Starting server on :8081...") // Changed port
	if err := r.Run(":8081"); err != nil {     // Changed port
		log.Fatalf("Failed to start server: %v", err)
	}
}
