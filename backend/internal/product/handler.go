package product

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/AshishBodhankar/clothing-store/backend/internal/user"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, db *sql.DB, secretKey string) {
	// Public routes
	r.GET("/products", func(c *gin.Context) {
		products, err := ListProducts(c.Request.Context(), db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, products)
	})

	r.GET("/products/:id", func(c *gin.Context) {
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

	// Protected routes
	protected := r.Group("/products")
	protected.Use(user.JWTMiddleware(secretKey))
	protected.Use(user.RoleMiddleware(user.RoleAdmin))
	{
		protected.POST("/", func(c *gin.Context) {
			// Handler logic for creating a product
		})

		protected.PUT("/:id", func(c *gin.Context) {
			// Handler logic for updating a product
		})

		protected.DELETE("/:id", func(c *gin.Context) {
			// Handler logic for deleting a product
		})
	}
}
