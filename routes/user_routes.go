// routes/user_routes.go
package routes

import (
	"open-contribute/controllers"
	"open-contribute/middlewares"
	"open-contribute/repositories"
	"open-contribute/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupUserRoutes(router *gin.Engine, db *gorm.DB, jwtSecret string) {
	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	public := router.Group("/users")
	{
		public.POST("/register", userController.Register)
		public.POST("/login", userController.Login)
	}

	// Authenticated routes (logged-in users only)
	authProtected := router.Group("/users")
	authProtected.Use(middlewares.AuthMiddleware()) //userService
	{
		authProtected.POST("/logout", userController.Logout)
		authProtected.GET("/:id", middlewares.SelfOnlyMiddleware(userService), userController.GetUserByID)
		authProtected.PUT("/:id", middlewares.SelfOnlyMiddleware(userService), userController.UpdateUser)
		authProtected.DELETE("/:id", middlewares.SelfOnlyMiddleware(userService), userController.DeleteUser)
		authProtected.PATCH("/:id", middlewares.SelfOnlyMiddleware(userService), userController.PatchUser)
	}

	// Superuser-protected routes
	superuserProtected := router.Group("/users")
	superuserProtected.Use(middlewares.AuthMiddleware()) //userService
	superuserProtected.Use(middlewares.SuperuserCheckMiddleware(userService))
	{
		superuserProtected.GET("/", userController.GetUsers)
	}
}
