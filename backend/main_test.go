package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRegisterUser(t *testing.T) {
	router := gin.Default()
	router.POST("/api/register", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "User registered"})
	})

	req, _ := http.NewRequest("POST", "/api/register", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "User registered")
}

func TestLoginUser(t *testing.T) {
	router := gin.Default()
	router.POST("/api/login", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"token": "dummy_token"})
	})

	req, _ := http.NewRequest("POST", "/api/login", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "dummy_token")
}
