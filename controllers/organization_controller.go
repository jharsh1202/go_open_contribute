// controllers/organization_controller.go
package controllers

import (
	"log"
	"net/http"
	"open-contribute/models"
	"open-contribute/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrganizationController struct {
	service *services.OrganizationService
}

func NewOrganizationController(service *services.OrganizationService) *OrganizationController {
	return &OrganizationController{service: service}
}

func (c *OrganizationController) CreateOrganization(ctx *gin.Context) {

	log.Printf("Create Organization")
	log.Printf("Request: %v", ctx.Request)

	var organizationRequest struct {
		Name    string `json:"name"`
		AdminID uint   `json:"admin_id"`
	}

	if err := ctx.ShouldBindJSON(&organizationRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	admin, err := c.service.GetUserByID(organizationRequest.AdminID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid admin ID"})
		return
	}

	organization := models.Organization{
		Name:    organizationRequest.Name,
		AdminID: organizationRequest.AdminID,
		Admin:   *admin,
	}

	log.Printf("Organization: %v", organization)

	if err := c.service.CreateOrganization(&organization); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, organization)
}

// func (c *OrganizationController) CreateOrganization(ctx *gin.Context) {
// 	var organization models.Organization
// 	if err := ctx.ShouldBindJSON(&organization); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	log.Printf("Organization: %v", organization)
// 	if err := c.service.CreateOrganization(&organization); err != nil {

// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, organization)
// }

func (c *OrganizationController) GetOrganizationByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid organization ID"})
		return
	}

	organization, err := c.service.GetOrganizationByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, organization)
}

func (c *OrganizationController) UpdateOrganization(ctx *gin.Context) {
	var organization models.Organization

	idParam := ctx.Param("id")
	id, _ := strconv.ParseUint(idParam, 10, 32)

	if err := ctx.ShouldBindJSON(&organization); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existingOrg, err := c.service.GetOrganizationByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	organization.ID = existingOrg.ID

	if err := c.service.UpdateOrganization(&organization); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, organization)
}

func (c *OrganizationController) DeleteOrganization(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid organization ID"})
		return
	}

	organization, err := c.service.GetOrganizationByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.DeleteOrganization(organization); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "organization deleted"})
}

func (c *OrganizationController) ListOrganizations(ctx *gin.Context) {
	organizations, err := c.service.ListOrganizations()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, organizations)
}

// PatchOrganization partially updates an organization's fields
func (c *OrganizationController) PatchOrganization(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid organization ID"})
		return
	}

	existingOrg, err := c.service.GetOrganizationByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "organization not found"})
		return
	}

	var updatedFields map[string]interface{}
	if err := ctx.ShouldBindJSON(&updatedFields); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if err := c.service.PatchOrganization(existingOrg, updatedFields); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, existingOrg)
}
