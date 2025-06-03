// Defines HTTP handlers for user-related endpoints using Gin.
package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) RegisterRoutes(r *gin.Engine, jwtSecret string) {
	r.POST("/api/register", h.Register)
	r.POST("/api/login", h.Login)
	// Protected routes
	authorized := r.Group("/api")
	authorized.Use(JWTMiddleware(jwtSecret)) // Pass the secret here
	{
		authorized.GET("/profile", h.Profile)
	}
}

func (h *Handler) Register(c *gin.Context) {
	var req struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	user, err := h.Service.Register(c.Request.Context(), req.Name, req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (h *Handler) Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	token, err := h.Service.Login(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *Handler) Profile(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}

	role, exists := c.Get("userRole") // Changed "role" to "userRole"
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Role not found in context"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user_id": userID,
		"role":    role, // This can stay as "role" in the JSON response if you like
	})
}
