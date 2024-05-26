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

type TaskController struct {
	service *services.TaskService
	// organizationService *services.OrganizationService //, organizationService *services.OrganizationService
}

func NewTaskController(
	service *services.TaskService) *TaskController {
	return &TaskController{service: service}
}

func (c *TaskController) CreateTask(ctx *gin.Context) {

	var projectRequest struct {
		Name    string `json:"name"`
		OwnerID uint   `json:"owner_id"`
		GoalID  uint   `json:"goal_id"`
	}

	if err := ctx.ShouldBindJSON(&projectRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	OwnerID := ctx.Keys["user_id"].(uint)
	// log.Printf("OwnerID: %v", OwnerID)
	owner, err := c.service.GetUserByID(OwnerID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid admin ID"})
		return
	}

	goal, err := c.service.GetGoalByID(projectRequest.GoalID)
	// goal.Owner = *owner
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Goal ID"})
		return
	}

	// project, _ := c.service.GetProjectByID(projectRequest.ProjectID)
	// if project == nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
	// 	return
	// }

	log.Printf("Task: %v", goal)

	task := models.Task{
		Name: projectRequest.Name,
		// ProjectID: projectRequest.ProjectID,
		OwnerID: OwnerID,
		// Project:   *project,
		GoalID: projectRequest.GoalID,
		Goal:   *goal,
		Owner:  *owner,
	}

	if err := c.service.CreateTask(&task); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, task)
}

func (c *TaskController) GetTaskByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID"})
		return
	}

	task, err := c.service.GetTaskByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, task)
}

func (c *TaskController) UpdateTask(ctx *gin.Context) {
	var task models.Task

	// Fetch the project ID from the request parameters
	idParam := ctx.Param("id")
	id, _ := strconv.ParseUint(idParam, 10, 32)

	// Fetch the existing project from the database
	existingTask, err := c.service.GetTaskByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Update the project's admin to the existing project's admin
	// project.Admin = existing_project.Admin

	// Bind the request body to the project struct
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Task: %v", task)
	owner, err := c.service.GetUserByID(existingTask.OwnerID)
	goal, err := c.service.GetGoalByID(existingTask.GoalID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Owner ID"})
		return
	}

	// existingProject, err := c.service.GetProjectByID(existingTask.ProjectID)
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
	// 	return
	// }

	task.Owner = *owner
	task.Goal = *goal
	// goal.Project = *existingProject

	// Update the project in the database
	if err := c.service.UpdateTask(&task); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, task)
}

func (c *TaskController) DeleteTask(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid project ID"})
		return
	}

	goal, err := c.service.GetTaskByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.DeleteTask(goal); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "goal deleted"})
}

func (c *TaskController) ListTasks(ctx *gin.Context) {
	projects, err := c.service.ListTasks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, projects)
}

// PatchTask partially updates an project's fields
func (c *TaskController) PatchTask(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid project ID"})
		return
	}

	log.Printf("Existing project id: %v", id)

	existingTask, err := c.service.GetTaskByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "project not found"})
		return
	}

	log.Printf("Existing project: %v", existingTask)

	var updatedFields map[string]interface{}
	if err := ctx.ShouldBindJSON(&updatedFields); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if err := c.service.PatchTask(existingTask, updatedFields); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, existingTask)
}
