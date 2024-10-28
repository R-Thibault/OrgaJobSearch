package middleware

import (
	"fmt"
	"net/http"

	"github.com/R-Thibault/OrgaJobSearch/backend/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware checks if the request contains a valid JWT token in the cookie.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtKey := []byte(config.GetConfig("JWT_KEY"))

		// Extract token from cookie
		cookie, err := c.Cookie("token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required: token missing"})
			c.Abort()
			return
		}

		// Parse the JWT token
		token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtKey, nil
		})

		if err != nil {
			// Check if the error is due to an expired token or other reasons
			if ve, ok := err.(*jwt.ValidationError); ok {
				if ve.Errors&jwt.ValidationErrorExpired != 0 {
					c.JSON(http.StatusUnauthorized, gin.H{"error": "Token expired"})
				} else {
					c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
				}
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			}
			c.Abort()
			return
		}

		// Check if the token is valid and extract claims
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Extract userUUID and role from the token claims
			if userUUID, ok := claims["userUUID"].(string); ok {
				c.Set("userUUID", userUUID) // Store userUUID in the context
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token: userUUID missing"})
				c.Abort()
				return
			}

			if role, ok := claims["role"].(string); ok {
				c.Set("role", role) // Store role in the context
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token: role missing"})
				c.Abort()
				return
			}
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Proceed to the next middleware/handler
		c.Next()
	}
}
