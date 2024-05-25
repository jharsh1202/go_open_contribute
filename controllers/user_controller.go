// controllers/user_controller.go
package controllers

import (
	"log"
	"net/http"
	"open-contribute/models"
	"open-contribute/services"
	"strconv"

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
		Username  string `json:"username" binding:"required"`
		Email     string `json:"email" binding:"required"`
		Phone     string `json:"phone" binding:"required"`
		Password  string `json:"password" binding:"required"`
		SuperUser bool
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := uc.userService.RegisterUser(req.Username, req.Email, req.Password, req.SuperUser, req.Phone); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

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

	user, token, err := uc.userService.LoginUser(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	userResponse := map[string]interface{}{
		"user":  user,
		"token": token,
	}
	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "data": userResponse})
}

func (uc *UserController) Logout(ctx *gin.Context) {
	// Optionally, you can implement server-side token invalidation logic here.
	// For now, we just send a response indicating successful logout.
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Successfully logged out",
	})
}

func (uc *UserController) GetUserByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid User ID"})
		return
	}

	usr, err := uc.userService.GetUserByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, usr)
}

func (uc *UserController) GetUsers(ctx *gin.Context) {
	log.Printf("GetUsers")
	var users []models.User

	users, err := uc.userService.GetUsers()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	var user models.User

	idParam := ctx.Param("id")
	id, _ := strconv.ParseUint(idParam, 10, 32)

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existingUser, err := uc.userService.GetUserByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	user.ID = existingUser.ID

	if err := uc.userService.UpdateUser(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (c *UserController) DeleteUser(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	user, err := c.userService.GetUserByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := c.userService.DeleteUser(user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "user deleted"})
}

func (c *UserController) PatchUser(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	existingOrg, err := c.userService.GetUserByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	var updatedFields map[string]interface{}
	if err := ctx.ShouldBindJSON(&updatedFields); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if err := c.userService.PatchUser(existingOrg, updatedFields); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, existingOrg)
}
