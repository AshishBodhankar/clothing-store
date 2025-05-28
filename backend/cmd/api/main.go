// Root: clothing-store/backend/cmd/api/main.go
package main

import (
	"clothing-store/backend/internal/common"
	"clothing-store/backend/internal/product"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := common.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}

	r := gin.Default()

	// Serve static files
	r.Static("/assets", "./frontend/assets")

	product.RegisterRoutes(r, db)

	// Future: user.RegisterRoutes(r, db)
	// Future: order.RegisterRoutes(r, db)

	r.Run(":8080")
}
