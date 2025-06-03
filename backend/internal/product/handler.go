package product

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/AshishBodhankar/clothing-store/backend/internal/user"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, db *sql.DB, secretKey string) {
	// Group for routes requiring any authenticated user
	// Define the group path without a trailing slash if you want /products to be the endpoint
	authenticatedUserRoutes := r.Group("/products")
	authenticatedUserRoutes.Use(user.JWTMiddleware(secretKey))
	{
		// This will handle GET /products
		authenticatedUserRoutes.GET("", func(c *gin.Context) {
			products, err := ListProducts(c.Request.Context(), db)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, products)
		})

		// This will handle GET /products/:id
		authenticatedUserRoutes.GET("/:id", func(c *gin.Context) {
			idStr := c.Param("id")
			id, err := strconv.Atoi(idStr)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
				return
			}
			product, err := GetProduct(c.Request.Context(), db, id)
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
				return
			}
			c.JSON(http.StatusOK, product)
		})
	}

	// Group for admin-only protected routes (management actions)
	adminProtectedRoutes := r.Group("/products")
	adminProtectedRoutes.Use(user.JWTMiddleware(secretKey))
	adminProtectedRoutes.Use(user.RoleMiddleware(user.RoleAdmin))
	{
		// This will handle POST /products
		adminProtectedRoutes.POST("", func(c *gin.Context) {
			// Handler logic for creating a product
		})

		// This will handle PUT /products/:id
		adminProtectedRoutes.PUT("/:id", func(c *gin.Context) {
			// Handler logic for updating a product
		})

		adminProtectedRoutes.DELETE("/:id", func(c *gin.Context) {
			// Handler logic for deleting a product
		})
	}
}
