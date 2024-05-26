package routes

import (
	"open-contribute/controllers"
	"open-contribute/middlewares"
	"open-contribute/repositories"
	"open-contribute/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupGoalRoutes(router *gin.Engine, db *gorm.DB, jwtSecret string) {
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
	// projectController := controllers.NewProjectController(&projectService)

	// Goal
	goalRepository := repositories.NewGoalRepository(db)
	goalService := services.NewGoalService(*goalRepository, projectService)
	goalController := controllers.NewGoalController(&goalService)

	// Public routes
	public := router.Group("/goals")
	{
		public.GET("/", goalController.ListGoals)
		public.GET("/:id", goalController.GetGoalByID)
	}

	// Admin-protected routes
	adminProtected := router.Group("/goals")
	// adminProtected.Use(middlewares.GoalAdminCheckMiddleware(userService, projectService))
	// adminProtected.Use(middlewares.OrgAdminCheckMiddleware(userService, organizationService))
	// adminProtected.Use(middlewares.SuperuserCheckMiddleware(userService))
	adminProtected.Use(middlewares.AuthMiddleware(userService))
	{
		adminProtected.POST("/", goalController.CreateGoal)
		adminProtected.PATCH("/:id", goalController.PatchGoal)
		adminProtected.PUT("/:id", goalController.UpdateGoal)
		adminProtected.DELETE("/:id", goalController.DeleteGoal)
	}

	// Admin-protected routes
	// projectAdminProtected := router.Group("/goals")

	// adminProtected.Use(middlewares.ProjectAdminCheckMiddleware(userService, projectService))
	// {
	// 	projectAdminProtected.POST("/", goalController.CreateProject)
	// }

}
