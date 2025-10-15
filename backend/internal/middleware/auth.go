package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Claims represents JWT custom claims
type Claims struct {
	UserID string `json:"user_id"` // Используем snake_case как в генерации токена
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// AuthMiddleware validates JWT token
func AuthMiddleware(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"error": gin.H{
					"code":    "UNAUTHORIZED",
					"message": "Authorization header is required",
				},
			})
			c.Abort()
			return
		}

		// Extract token from "Bearer <token>"
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"error": gin.H{
					"code":    "UNAUTHORIZED",
					"message": "Invalid authorization format",
				},
			})
			c.Abort()
			return
		}

		tokenString := parts[1]

		// Parse and validate token
		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecret), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"error": gin.H{
					"code":    "UNAUTHORIZED",
					"message": "Invalid or expired token",
				},
			})
			c.Abort()
			return
		}

		// Extract claims
		if claims, ok := token.Claims.(*Claims); ok {
			c.Set("userId", claims.UserID)
			c.Set("role", claims.Role)
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"error": gin.H{
					"code":    "UNAUTHORIZED",
					"message": "Invalid token claims",
				},
			})
			c.Abort()
			return
		}
	}
}

// AdminMiddleware checks if user has admin role
func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists || role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"error": gin.H{
					"code":    "FORBIDDEN",
					"message": "Admin access required",
				},
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

// OptionalAuthMiddleware validates JWT token if present, but allows guests
func OptionalAuthMiddleware(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		// If no auth header, allow guest access
		if authHeader == "" {
			// Generate a temporary session ID for guest
			sessionID := c.GetHeader("X-Session-ID")
			if sessionID != "" {
				c.Set("session_id", sessionID)
			}
			c.Next()
			return
		}

		// Extract token from "Bearer <token>"
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.Next() // Allow to proceed even with invalid format
			return
		}

		tokenString := parts[1]

		// Parse and validate token
		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecret), nil
		})

		if err != nil || !token.Valid {
			// Token invalid, but allow guest access
			c.Next()
			return
		}

		// Extract claims
		if claims, ok := token.Claims.(*Claims); ok {
			c.Set("user_id", claims.UserID)
			c.Set("role", claims.Role)
		}

		c.Next()
	}
}
