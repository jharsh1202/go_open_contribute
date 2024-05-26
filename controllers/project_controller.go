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

type ProjectController struct {
	service *services.ProjectService
	// organizationService *services.OrganizationService //, organizationService *services.OrganizationService
}

func NewProjectController(
	service *services.ProjectService) *ProjectController {
	return &ProjectController{service: service}
}

func (c *ProjectController) CreateProject(ctx *gin.Context) {

	var projectRequest struct {
		Name           string `json:"name"`
		AdminID        uint   `json:"admin_id"`
		OrganizationID uint   `json:"organization_id"`
	}

	if err := ctx.ShouldBindJSON(&projectRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//  log.Printf("keys: %v", ctx.Keys)
	//  log.Printf("keys: %v", projectRequest.OrganizationID)

	// //  log.Printf("params: %v", ctx.Params)

	AdminID := ctx.Keys["user_id"].(uint)
	admin, err := c.service.GetUserByID(AdminID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid admin ID"})
		return
	}
	// //  log.Printf("Project: %v", admin)
	//  log.Printf("Project: %v", projectRequest.OrganizationID)

	organization, _ := c.service.GetOrganizationByID(projectRequest.OrganizationID)

	// //  log.Printf("Project: %v", organization)

	project := models.Project{
		Name:           projectRequest.Name,
		OrganizationID: projectRequest.OrganizationID,
		Organization:   *organization,
		AdminID:        AdminID,
		Admin:          *admin,
	}

	if err := c.service.CreateProject(&project); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, project)
}

// func (c *ProjectController) CreateProject(ctx *gin.Context) {
// 	var project models.Project
// 	if err := ctx.ShouldBindJSON(&project); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	//  log.Printf("Project: %v", project)
// 	if err := c.service.CreateProject(&project); err != nil {

// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, project)
// }

func (c *ProjectController) GetProjectByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid project ID"})
		return
	}

	project, err := c.service.GetProjectByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, project)
}

func (c *ProjectController) UpdateProject(ctx *gin.Context) {
	var project models.Project

	// Fetch the project ID from the request parameters
	idParam := ctx.Param("id")
	id, _ := strconv.ParseUint(idParam, 10, 32)

	// Fetch the existing project from the database
	existingProject, err := c.service.GetProjectByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Update the project's admin to the existing project's admin
	// project.Admin = existing_project.Admin

	// Bind the request body to the project struct
	if err := ctx.ShouldBindJSON(&project); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	project.Admin = existingProject.Admin
	project.Organization = existingProject.Organization
	// project.Organization.Admin = existingProject.Organization.Admin

	// Update the project in the database
	if err := c.service.UpdateProject(&project); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, project)
}

func (c *ProjectController) DeleteProject(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid project ID"})
		return
	}

	project, err := c.service.GetProjectByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.DeleteProject(project); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "project deleted"})
}

func (c *ProjectController) ListProjects(ctx *gin.Context) {
	projects, err := c.service.ListProjects()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, projects)
}

// PatchProject partially updates an project's fields
func (c *ProjectController) PatchProject(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid project ID"})
		return
	}

	log.Printf("Existing project id: %v", id)

	existingProject, err := c.service.GetProjectByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "project not found"})
		return
	}

	log.Printf("Existing project: %v", existingProject)

	var updatedFields map[string]interface{}
	if err := ctx.ShouldBindJSON(&updatedFields); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if err := c.service.PatchProject(existingProject, updatedFields); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, existingProject)
}
