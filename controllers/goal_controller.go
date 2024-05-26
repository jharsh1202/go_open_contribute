// controllers/project_controller.go
package controllers

import (
	"log"
	"net/http"
	"open-contribute/models"
	"open-contribute/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GoalController struct {
	service *services.GoalService
	// organizationService *services.OrganizationService //, organizationService *services.OrganizationService
}

func NewGoalController(
	service *services.GoalService) *GoalController {
	return &GoalController{service: service}
}

func (c *GoalController) CreateGoal(ctx *gin.Context) {

	var projectRequest struct {
		Name      string `json:"name"`
		OwnerID   uint   `json:"owner_id"`
		ProjectID uint   `json:"project_id"`
	}

	if err := ctx.ShouldBindJSON(&projectRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//  log.Printf("keys: %v", ctx.Keys)
	//  log.Printf("keys: %v", projectRequest.OrganizationID)

	// //  log.Printf("params: %v", ctx.Params)

	OwnerID := ctx.Keys["user_id"].(uint)
	log.Printf("OwnerID: %v", OwnerID)
	owner, err := c.service.GetUserByID(OwnerID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid admin ID"})
		return
	}
	// //  log.Printf("Goal: %v", admin)
	//  log.Printf("Goal: %v", projectRequest.OrganizationID)

	// organization, _ := c.service.GetOrganizationByID(projectRequest.OrganizationID)

	// log.Printf("Goal: %v", c.service)
	project, _ := c.service.GetProjectByID(projectRequest.ProjectID)
	if project == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	log.Printf("Goal: %v", project)

	goal := models.Goal{
		Name:      projectRequest.Name,
		ProjectID: projectRequest.ProjectID,
		OwnerID:   OwnerID,
		Project:   *project,
		Owner:     *owner,
	}

	if err := c.service.CreateGoal(&goal); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, goal)
}

func (c *GoalController) GetGoalByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid goal ID"})
		return
	}

	goal, err := c.service.GetGoalByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, goal)
}

func (c *GoalController) UpdateGoal(ctx *gin.Context) {
	var goal models.Goal

	// Fetch the project ID from the request parameters
	idParam := ctx.Param("id")
	id, _ := strconv.ParseUint(idParam, 10, 32)

	// Fetch the existing project from the database
	existingGoal, err := c.service.GetGoalByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Update the project's admin to the existing project's admin
	// project.Admin = existing_project.Admin

	// Bind the request body to the project struct
	if err := ctx.ShouldBindJSON(&goal); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Goal: %v", existingGoal.OwnerID)
	owner, err := c.service.GetUserByID(existingGoal.OwnerID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Owner ID"})
		return
	}

	existingProject, err := c.service.GetProjectByID(existingGoal.ProjectID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	goal.Owner = *owner
	goal.Project = *existingProject

	// Update the project in the database
	if err := c.service.UpdateGoal(&goal); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, goal)
}

func (c *GoalController) DeleteGoal(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid project ID"})
		return
	}

	goal, err := c.service.GetGoalByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.DeleteGoal(goal); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "goal deleted"})
}

func (c *GoalController) ListGoals(ctx *gin.Context) {
	projects, err := c.service.ListGoals()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, projects)
}

// // PatchGoal partially updates an project's fields
func (c *GoalController) PatchGoal(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid project ID"})
		return
	}

	log.Printf("Existing project id: %v", id)

	existingGoal, err := c.service.GetGoalByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "project not found"})
		return
	}

	log.Printf("Existing project: %v", existingGoal)

	var updatedFields map[string]interface{}
	if err := ctx.ShouldBindJSON(&updatedFields); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if err := c.service.PatchGoal(existingGoal, updatedFields); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, existingGoal)
}
