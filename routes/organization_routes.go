// routes/organization_routes.go
package routes

import (
	"open-contribute/controllers"
	"open-contribute/middlewares"
	"open-contribute/repositories"
	"open-contribute/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupOrganizationRoutes(router *gin.Engine, db *gorm.DB, jwtSecret string) {
	// Initialize repositories and services

	// User
	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)

	// Organization
	organizationRepository := repositories.NewOrganizationRepository(db)
	organizationService := services.NewOrganizationService(organizationRepository, userService)
	organizationController := controllers.NewOrganizationController(organizationService)

	// Public routes
	public := router.Group("/organizations")
	{
		public.GET("/", organizationController.ListOrganizations)
		public.GET("/:id", organizationController.GetOrganizationByID)
	}

	// Admin-protected routes
	adminProtected := router.Group("/organizations")
	adminProtected.Use(middlewares.AuthMiddleware()) // userService Pass the userService to AuthMiddleware
	adminProtected.Use(middlewares.AdminCheckMiddleware(userService))
	{
		adminProtected.POST("/", organizationController.CreateOrganization)
		adminProtected.PATCH("/:id", organizationController.PatchOrganization)
		adminProtected.PUT("/:id", organizationController.UpdateOrganization)
		adminProtected.DELETE("/:id", organizationController.DeleteOrganization)
	}
}
