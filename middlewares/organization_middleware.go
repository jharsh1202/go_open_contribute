package middlewares

import (
	"net/http"
	"open-contribute/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func OrgAdminCheckMiddleware(userService services.UserService, orgService services.OrganizationService) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetUint("user_id")
		orgIDStr := c.Param("id")
		orgID, err := strconv.ParseUint(orgIDStr, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid organization ID"})
			c.Abort()
			return
		}

		org, err := orgService.GetOrganizationByID(uint(orgID))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Organization not found"})
			c.Abort()
			return
		}

		user, err := userService.GetUserByID(userID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		if user.SuperUser || org.AdminID == userID {
			c.Next()
		} else {
			c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to access this organization"})
			c.Abort()
		}
	}
}
