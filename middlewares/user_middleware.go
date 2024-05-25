package middlewares

import (
	"net/http"
	"open-contribute/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SelfOnlyMiddleware(userService services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get authenticated user ID from context (set by AuthMiddleware)
		authUserID, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Assert the type to uint
		authUserIDUint, ok := authUserID.(uint)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID"})
			c.Abort()
			return
		}

		// Fetch the authenticated user details
		authUser, err := userService.GetUserByID(authUserIDUint)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Get the user ID from the request parameters
		userIDStr := c.Param("id")
		if userIDStr == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
			c.Abort()
			return
		}

		userID, err := strconv.ParseUint(userIDStr, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			c.Abort()
			return
		}

		// Allow if the user is a super user or if they are accessing their own details
		if authUser.SuperUser || authUserIDUint == uint(userID) {
			c.Next()
		} else {
			c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to access this resource"})
			c.Abort()
		}
	}
}
