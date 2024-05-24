// routes/user_routes.go
package routes

import (
	"open-contribute/controllers"
	"open-contribute/repositories"
	"open-contribute/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupUserRoutes(router *gin.Engine, db *gorm.DB, jwtSecret string) {
	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/register", userController.Register)
		userRoutes.POST("/login", userController.Login)
		userRoutes.POST("/logout", userController.Logout)
		userRoutes.GET("/:id", userController.GetUserByID)

		userRoutes.GET("/", userController.GetUsers)
		// userRoutes.POST("/", userController.CreateUser)
		userRoutes.PUT("/:id", userController.UpdateUser)
		userRoutes.DELETE("/:id", userController.DeleteUser)
		userRoutes.PATCH("/:id", userController.PatchUser)
	}
}
