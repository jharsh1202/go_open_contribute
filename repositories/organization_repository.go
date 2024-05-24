// repositories/organization_repository.go
package repositories

import (
	"log"
	"open-contribute/models"

	"gorm.io/gorm"
)

type OrganizationRepository struct {
	db *gorm.DB
}

func NewOrganizationRepository(db *gorm.DB) *OrganizationRepository {
	return &OrganizationRepository{db: db}
}

func (r *OrganizationRepository) Create(organization *models.Organization) error {
	return r.db.Create(organization).Error
}

func (r *OrganizationRepository) GetByID(id uint) (*models.Organization, error) {
	log.Printf("ID %v", id)
	var organization models.Organization
	// if err := r.db.Preload("Admin").Preload("Members").First(&organization, id).Error; err != nil {
	if err := r.db.First(&organization, id).Error; err != nil {
		return nil, err
	}
	return &organization, nil
}

func (r *OrganizationRepository) Update(organization *models.Organization) error {
	return r.db.Save(&organization).Error
}

func (r *OrganizationRepository) Delete(organization *models.Organization) error {
	return r.db.Delete(organization).Error
}

func (r *OrganizationRepository) List() ([]models.Organization, error) {
	var organizations []models.Organization
	// if err := r.db.Preload("Admin").Preload("Members").Find(&organizations).Error; err != nil {
	if err := r.db.Find(&organizations).Error; err != nil {
		return nil, err
	}
	return organizations, nil
}

func (r *OrganizationRepository) Patch(existingOrg *models.Organization, updatedFields map[string]interface{}) error {
	return r.db.Model(existingOrg).Updates(updatedFields).Error
}
