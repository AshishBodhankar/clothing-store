package order

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, db *sql.DB) {
	r.POST("/orders", func(c *gin.Context) {
		// TODO: Add order placement logic
	})

	r.GET("/orders", func(c *gin.Context) {
		// TODO: Add order listing logic
	})
}
