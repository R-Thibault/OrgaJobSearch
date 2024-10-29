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

		// The same key used for signing
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
			if ve, ok := err.(*jwt.ValidationError); ok {
				if ve.Errors&jwt.ValidationErrorExpired != 0 {
					c.JSON(http.StatusUnauthorized, gin.H{"error": "Token expired"})
					c.Abort()
					return
				} else {
					c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
					c.Abort()
					return
				}
			}
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(*models.JWTToken); ok && token.Valid {
			// Access the userUUID from the Body field
			if claims.Body == nil || *claims.Body == "" {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token: userUUID missing"})
				c.Abort()
				return
			}
			// Decode the JSON string in Body
			var bodyContent map[string]string
			if err := json.Unmarshal([]byte(*claims.Body), &bodyContent); err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to parse token body"})
				c.Abort()
				return
			}

			userUUID, uuidExists := bodyContent["userUUID"]
			userRole, roleExists := bodyContent["userRole"]
			if !uuidExists || !roleExists {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token: required user data missing"})
				c.Abort()
				return
			}

			// Store `userUUID` and `userRole` in context
			c.Set("userUUID", userUUID)
			c.Set("userRole", userRole)

		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Proceed to the next middleware/handler
		c.Next()
	}
}

func RoleMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("userRole")
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{"error": "User role not found in context"})
			c.Abort()
			return
		}
		userRolestr, ok := userRole.(string)
		if !ok {
			c.JSON(http.StatusForbidden, gin.H{"error": "Invalid user role format"})
			c.Abort()
			return
		}

		for _, role := range allowedRoles {
			if userRolestr == role {
				c.Next()
				return
			}
		}
		c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient permission"})
		c.Abort()
	}
}
