package middlewares

import (
	"net/http"
	"open-contribute/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ProjectAdminCheckMiddleware(userService services.UserService, projectService services.ProjectService) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetUint("user_id")

		projectIDStr := c.Param("id")
		projectID, err := strconv.ParseUint(projectIDStr, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
			c.Abort()
			return
		}

		project, err := projectService.GetProjectByID(uint(projectID))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Project not found"})
			c.Abort()
			return
		}

		user, err := userService.GetUserByID(userID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized - User Not Found"})
			c.Abort()
			return
		}

		if project.AdminID == user.ID || user.SuperUser || project.Organization.AdminID == user.ID {
			c.Next()
		} else {
			c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to access this organization"})
			c.Abort()
		}
	}
}
