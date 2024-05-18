// routes/routes.go
package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	// // Assuming db is already initialized and passed here
	// var db *gorm.DB

	// Initialize and setup routes for various entities
	SetupUserRoutes(router, db)
	// Add more routes setup functions here for other entities
}
