package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/R-Thibault/OrgaJobSearch/backend/config"
	"github.com/R-Thibault/OrgaJobSearch/backend/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware checks if the request contains a valid JWT token in the cookie.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
			c.Abort()
			return
		}

		// JWT key for validation
		var jwtKey = []byte(config.GetConfig("JWT_KEY"))

		// Parse the token
		token, err := jwt.ParseWithClaims(tokenString, &models.JWTToken{}, func(token *jwt.Token) (interface{}, error) {
			// Ensure the signing method is correct
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return jwtKey, nil
		})

		if err != nil {
			if ve, ok := err.(*jwt.ValidationError); ok && ve.Errors&jwt.ValidationErrorExpired != 0 {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Token expired"})
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			}
			c.Abort()
			return
		}

		// Access claims if the token is valid
		if claims, ok := token.Claims.(*models.JWTToken); ok && token.Valid {
			var bodyContent map[string]interface{}
			if err := json.Unmarshal([]byte(*claims.Body), &bodyContent); err != nil {

				c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to parse token body"})
				c.Abort()
				return
			}

			// Extract userUUID and userRole from the token body content
			userUUID, uuidExists := bodyContent["userUUID"].(string)
			userRoles, rolesExist := bodyContent["userRole"].([]interface{})

			if !uuidExists || !rolesExist {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token: required user data missing"})
				c.Abort()
				return
			}

			// Store userUUID and userRole in context for further use
			c.Set("userUUID", userUUID)
			c.Set("userRoles", userRoles)

		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Continue to the next handler
		c.Next()
	}
}

func RoleMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRoles, exists := c.Get("userRoles")
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{"error": "User role not found in context"})
			c.Abort()
			return
		}

		// Check if any role in userRoles matches allowedRoles
		rolesSlice, ok := userRoles.([]interface{})

		if !ok {
			c.JSON(http.StatusForbidden, gin.H{"error": "Invalid user role format"})
			c.Abort()
			return
		}
		for _, role := range rolesSlice {
			for _, allowedRole := range allowedRoles {
				if role == allowedRole {
					c.Next()
					return
				}
			}
		}
		c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient permission"})
		c.Abort()
	}
}
