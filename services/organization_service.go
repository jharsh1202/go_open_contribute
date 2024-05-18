// services/organization_service.go
package services

import (
	"log"
	"open-contribute/models"
	"open-contribute/repositories"
)

type OrganizationService struct {
	repo *repositories.OrganizationRepository
}

func NewOrganizationService(repo *repositories.OrganizationRepository) *OrganizationService {
	return &OrganizationService{repo: repo}
}

func (s *OrganizationService) CreateOrganization(organization *models.Organization) error {
	return s.repo.Create(organization)
}

func (s *OrganizationService) GetOrganizationByID(id uint) (*models.Organization, error) {
	return s.repo.GetByID(id)
}

func (s *OrganizationService) UpdateOrganization(organization *models.Organization) error { //, adminID uint
	// if organization.AdminID != adminID {
	// 	return errors.New("only admin can update the organization")
	// }
	log.Printf("%v", organization)

	return s.repo.Update(organization)
}

// func (s *OrganizationService) UpdateOrganization(id uint, updatedOrg *models.Organization) error {
// 	// Check if the organization with the given ID exists
// 	existingOrg, err := s.repo.GetByID(id)
// 	if err != nil {
// 		return err // Return error if organization not found
// 	}

// 	// Perform any necessary validation on the updated organization data
// 	// For example, ensure the name is not empty or validate other fields

// 	// Update the fields of the existing organization with the new values
// 	existingOrg.Name = updatedOrg.Name
// 	existingOrg.AdminID = updatedOrg.AdminID
// 	// Update other fields as needed

// 	// Save the updated organization back to the database
// 	if err := s.repo.Update(existingOrg); err != nil {
// 		return err // Return any database error that occurred during update
// 	}

// 	return nil // Return nil if update was successful
// }

func (s *OrganizationService) DeleteOrganization(organization *models.Organization) error { //, adminID uint
	// if organization.AdminID != adminID {
	// 	return errors.New("only admin can delete the organization")
	// }
	return s.repo.Delete(organization)
}

func (s *OrganizationService) ListOrganizations() ([]models.Organization, error) {
	return s.repo.List()
}
