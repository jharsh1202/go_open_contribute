package routes

import (
	"open-contribute/controllers"
	"open-contribute/middlewares"
	"open-contribute/repositories"
	"open-contribute/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupTaskRoutes(router *gin.Engine, db *gorm.DB, jwtSecret string) {
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

	// Task
	taskRepository := repositories.NewTaskRepository(db)
	taskService := services.NewTaskService(*taskRepository, goalService)
	taskController := controllers.NewTaskController(&taskService)

	// Public routes
	public := router.Group("/tasks")
	{
		public.GET("/", taskController.ListTasks)
		public.GET("/:id", taskController.GetTaskByID)
	}

	// Admin-protected routes
	adminProtected := router.Group("/tasks")
	// adminProtected.Use(middlewares.GoalAdminCheckMiddleware(userService, projectService))
	// adminProtected.Use(middlewares.OrgAdminCheckMiddleware(userService, organizationService))
	// adminProtected.Use(middlewares.SuperuserCheckMiddleware(userService))
	adminProtected.Use(middlewares.AuthMiddleware(userService))
	{
		adminProtected.POST("/", taskController.CreateTask)
		adminProtected.PATCH("/:id", taskController.PatchTask)
		adminProtected.PUT("/:id", taskController.UpdateTask)
		adminProtected.DELETE("/:id", taskController.DeleteTask)
	}

	// Admin-protected routes
	// projectAdminProtected := router.Group("/goals")

	// adminProtected.Use(middlewares.ProjectAdminCheckMiddleware(userService, projectService))
	// {
	// 	projectAdminProtected.POST("/", taskController.CreateProject)
	// }

}
