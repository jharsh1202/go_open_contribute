// middlewares/auth_middleware.go
package middlewares

import (
	"net/http"
	"open-contribute/services"
	"open-contribute/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc { //userService services.UserService
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header"})
			c.Abort()
			return
		}

		token := parts[1]
		userID, err := utils.ParseJWT(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("user_id", userID)
		c.Next()
	}
}

func SuperuserCheckMiddleware(userService services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Keys["user_id"].(uint)
		user, err := userService.GetUserByID(userID)
		if err != nil || !user.SuperUser {
			c.JSON(http.StatusForbidden, gin.H{"error": "Superuser access required"})
			c.Abort()
			return
		}
		c.Next()
	}
}
