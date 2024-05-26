package routes

import (
	"open-contribute/controllers"
	"open-contribute/middlewares"
	"open-contribute/repositories"
	"open-contribute/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupProjectRoutes(router *gin.Engine, db *gorm.DB, jwtSecret string) {
	// Initialize repositories and services

	// User
	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)

	// Organization
	organizationRepository := repositories.NewOrganizationRepository(db)
	organizationService := services.NewOrganizationService(*organizationRepository, userService)

	// Project
	projectRepository := repositories.NewProjectRepository(db)
	projectService := services.NewProjectService(*projectRepository, organizationService)
	projectController := controllers.NewProjectController(&projectService)

	// Public routes
	public := router.Group("/projects")
	{
		public.GET("/", projectController.ListProjects)
		public.GET("/:id", projectController.GetProjectByID)
	}

	// Admin-protected routes
	adminProtected := router.Group("/projects")
	// adminProtected.Use(middlewares.ProjectAdminCheckMiddleware(userService, projectService))
	// adminProtected.Use(middlewares.OrgAdminCheckMiddleware(userService, organizationService))
	// adminProtected.Use(middlewares.SuperuserCheckMiddleware(userService))
	adminProtected.Use(middlewares.AuthMiddleware(userService))
	{
		adminProtected.POST("/", projectController.CreateProject)
		adminProtected.PATCH("/:id", projectController.PatchProject)
		adminProtected.PUT("/:id", projectController.UpdateProject)
		adminProtected.DELETE("/:id", projectController.DeleteProject)
	}

	// Admin-protected routes
	// projectAdminProtected := router.Group("/projects")

	// adminProtected.Use(middlewares.ProjectAdminCheckMiddleware(userService, projectService))
	// {
	// 	projectAdminProtected.POST("/", projectController.CreateProject)
	// }

}
