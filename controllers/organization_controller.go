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
	var organization models.Organization
	if err := ctx.ShouldBindJSON(&organization); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// userID := ctx.MustGet("userID").(uint)
	// organization.AdminID = userID
	if err := c.service.CreateOrganization(&organization); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, organization)
}

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

	// if err := ctx.ShouldBindJSON(&organization); err != nil {
	// 	log.Printf("UPDATEEEEE %v", organization)
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	log.Printf("UPDATEEEEE %v", organization)
	// userID := ctx.MustGet("userID").(uint)
	if err := c.service.UpdateOrganization(&organization); err != nil { //userID
		ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
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
	// userID := ctx.MustGet("userID").(uint)

	organization, err := c.service.GetOrganizationByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Optional: Ensure the user is authorized to delete the organization
	// if organization.AdminID != userID {
	// 	ctx.JSON(http.StatusForbidden, gin.H{"error": "only admin can delete the organization"})
	// 	return
	// }

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
