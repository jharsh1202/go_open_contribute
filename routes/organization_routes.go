// routes/organization_routes.go
package routes

import (
	"open-contribute/controllers"
	"open-contribute/repositories"
	"open-contribute/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupOrganizationRoutes(router *gin.Engine, db *gorm.DB, jwtSecret string) {
	// userRepository := repositories.NewUserRepository(db)
	// userService := services.NewUserService(userRepository)
	organizationRepository := repositories.NewOrganizationRepository(db)
	organizationService := services.NewOrganizationService(organizationRepository)
	organizationController := controllers.NewOrganizationController(organizationService)

	adminProtected := router.Group("/organizations")
	// adminProtected.Use(middlewares.AuthMiddleware())
	// adminProtected.Use(middlewares.AdminCheckMiddleware(userService))
	{
		adminProtected.POST("/", organizationController.CreateOrganization)
		adminProtected.GET("/:id", organizationController.GetOrganizationByID)
		adminProtected.PUT("/:id", organizationController.UpdateOrganization)
		adminProtected.DELETE("/:id", organizationController.DeleteOrganization)
		adminProtected.GET("/", organizationController.ListOrganizations)
	}
}
