// controllers/user_controller.go
package controllers

import (
	"log"
	"net/http"
	"open-contribute/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{userService: userService}
}

func (uc *UserController) Register(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
		IsAdmin  bool
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := uc.userService.RegisterUser(req.Username, req.Email, req.Password, req.IsAdmin); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("error: %v", req)

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func (uc *UserController) Login(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := uc.userService.LoginUser(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "user": user})
}

func (c *UserController) Logout(ctx *gin.Context) {
	// Optionally, you can implement server-side token invalidation logic here.
	// For now, we just send a response indicating successful logout.
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Successfully logged out",
	})
}
