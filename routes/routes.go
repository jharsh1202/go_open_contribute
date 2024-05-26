// routes/routes.go
package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB, jwtSecret string) {
	// Initialize and setup routes for various entities
	SetupUserRoutes(router, db, jwtSecret)
	SetupOrganizationRoutes(router, db, jwtSecret)
	SetupProjectRoutes(router, db, jwtSecret)
	SetupGoalRoutes(router, db, jwtSecret)
	// Add more routes setup functions here for other entities
}
