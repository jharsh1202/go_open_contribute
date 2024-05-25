// services/organization_service.go
package services

import (
	"errors"
	"log"
	"open-contribute/models"
	"open-contribute/repositories"
)

type OrganizationService struct {
	repo        *repositories.OrganizationRepository
	userService UserService
}

func NewOrganizationService(repo repositories.OrganizationRepository, userService UserService) OrganizationService {
	return OrganizationService{
		repo:        &repo,
		userService: userService,
	}
}

func (s OrganizationService) CheckUserExists(userID uint) (bool, error) {
	return s.userService.CheckUserExists(userID)
}

func (s *OrganizationService) GetUserByID(userID uint) (*models.User, error) {
	//  log.Printf("User ID %v", userID)
	return s.userService.GetUserByID(userID)
}

func (s *OrganizationService) CreateOrganization(organization *models.Organization) error {
	usr, err := s.userService.CheckUserExists(organization.AdminID)
	if !usr {
		return errors.New("admin user does not exist")
	}
	if err != nil {
		return errors.New("admin user does not exist")
	}

	return s.repo.Create(organization)
}

func (s *OrganizationService) GetOrganizationByID(id uint) (*models.Organization, error) {
	// log.Printf("Organization ID %v", id)
	log.Printf("Organization ID %v", s)
	return s.repo.GetByID(id)
}

func (s *OrganizationService) UpdateOrganization(organization *models.Organization) error { //, adminID uint
	return s.repo.Update(organization)
}

func (s *OrganizationService) PatchOrganization(existingOrg *models.Organization, updatedFields map[string]interface{}) error {
	return s.repo.Patch(existingOrg, updatedFields)
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
