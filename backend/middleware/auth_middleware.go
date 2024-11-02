package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/R-Thibault/OrgaJobSearch/backend/config"
	"github.com/R-Thibault/OrgaJobSearch/backend/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware checks if the request contains a valid JWT token in the cookie.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Attempt to retrieve the token from cookies
		tokenString, err := c.Cookie("token")
		if err != nil {
			log.Println("No token provided")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
			c.Abort()
			return
		}
		log.Printf("Token retrieved: %s", tokenString)

		// JWT key for validation
		var jwtKey = []byte(config.GetConfig("JWT_KEY"))
		log.Println("JWT key retrieved from config")

		// Parse the token
		token, err := jwt.ParseWithClaims(tokenString, &models.JWTToken{}, func(token *jwt.Token) (interface{}, error) {
			// Ensure the signing method is correct
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				log.Printf("Unexpected signing method: %v", token.Header["alg"])
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return jwtKey, nil
		})

		if err != nil {
			if ve, ok := err.(*jwt.ValidationError); ok && ve.Errors&jwt.ValidationErrorExpired != 0 {
				log.Println("Token expired")
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Token expired"})
			} else {
				log.Printf("Invalid token: %v", err)
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			}
			c.Abort()
			return
		}

		// Access claims if the token is valid
		if claims, ok := token.Claims.(*models.JWTToken); ok && token.Valid {
			log.Printf("Token claims validated: %+v", claims)

			// Parse body content from token claims
			var bodyContent map[string]interface{}
			if err := json.Unmarshal([]byte(*claims.Body), &bodyContent); err != nil {
				log.Println("Failed to parse token body")
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to parse token body"})
				c.Abort()
				return
			}
			log.Printf("Token body content: %+v", bodyContent)

			// Extract userUUID and userRole from the token body content
			userUUID, uuidExists := bodyContent["userUUID"].(string)
			userRoles, rolesExist := bodyContent["userRole"].([]interface{})
			if !uuidExists || !rolesExist {
				log.Println("Invalid token: required user data missing")
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token: required user data missing"})
				c.Abort()
				return
			}
			log.Printf("User UUID: %s, User Roles: %v", userUUID, userRoles)

			// Store userUUID and userRole in context for further use
			c.Set("userUUID", userUUID)
			c.Set("userRoles", userRoles)
		} else {
			log.Println("Invalid token: token claims are not valid")
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
			log.Println("User role not found in context")
			c.JSON(http.StatusForbidden, gin.H{"error": "User role not found in context"})
			c.Abort()
			return
		}
		log.Printf("User roles found in context: %v", userRoles)

		// Check if any role in userRoles matches allowedRoles
		rolesSlice, ok := userRoles.([]interface{})
		if !ok {
			log.Println("Invalid user role format")
			c.JSON(http.StatusForbidden, gin.H{"error": "Invalid user role format"})
			c.Abort()
			return
		}

		// Check role permissions
		for _, role := range rolesSlice {
			for _, allowedRole := range allowedRoles {
				if role == allowedRole {
					log.Printf("User role %v is allowed", role)
					c.Next()
					return
				}
			}
		}
		log.Printf("User has insufficient permission. Allowed roles: %v, User roles: %v", allowedRoles, rolesSlice)
		c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient permission"})
		c.Abort()
	}
}
