package user

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, db *sql.DB) {
	r.POST("/register", func(c *gin.Context) {
		// TODO: Add registration logic
	})

	r.POST("/login", func(c *gin.Context) {
		// TODO: Add login logic
	})
}
