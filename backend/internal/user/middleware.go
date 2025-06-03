// Implements JWT authentication middleware.
package user

import (
	"log" // Added for logging
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			log.Println("[JWTMiddleware] Error: Failed to assert claims as jwt.MapClaims")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			return
		}
		log.Printf("[JWTMiddleware] Claims successfully parsed: %+v\n", claims)

		// Set user ID and role from claims into the context
		if userIDClaim, found := claims["user_id"]; found {
			log.Printf("[JWTMiddleware] Found 'user_id' claim: %v (type: %T)\n", userIDClaim, userIDClaim)
			c.Set("userID", userIDClaim)
			log.Printf("[JWTMiddleware] Context after setting userID: %+v\n", c.Keys)
		} else {
			log.Println("[JWTMiddleware] Error: 'user_id' claim not found in token")
		}

		if userRoleClaim, found := claims["role"]; found {
			log.Printf("[JWTMiddleware] Found 'role' claim: %v (type: %T)\n", userRoleClaim, userRoleClaim)
			c.Set("userRole", userRoleClaim)
			log.Printf("[JWTMiddleware] Context after setting userRole: %+v\n", c.Keys)
		} else {
			log.Println("[JWTMiddleware] Error: 'role' claim not found in token")
		}

		log.Println("[JWTMiddleware] Proceeding with c.Next()")
		c.Next()
	}
}

func RoleMiddleware(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("userRole")
		if !exists || role != requiredRole {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			return
		}
		c.Next()
	}
}
