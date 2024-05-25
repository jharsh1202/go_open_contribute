// services/project_service.go
package services

import (
	"open-contribute/models"
	"open-contribute/repositories"
)

type ProjectService struct {
	repo                *repositories.ProjectRepository
	organizationService OrganizationService
}

func NewProjectService(repo repositories.ProjectRepository, organizationService OrganizationService) ProjectService {
	return ProjectService{
		repo:                &repo,
		organizationService: organizationService,
	}
}

func (s ProjectService) CheckUserExists(userID uint) (bool, error) {
	return s.organizationService.CheckUserExists(userID)
}

func (s *ProjectService) GetUserByID(userID uint) (*models.User, error) {
	//  log.Printf("User ID %v", userID)
	return s.organizationService.GetUserByID(userID)
}

func (s *ProjectService) GetOrganizationByID(userID uint) (*models.Organization, error) {
	//  log.Printf("User ID %v", userID)
	return s.organizationService.GetOrganizationByID(userID)
}

func (s *ProjectService) CreateProject(project *models.Project) error {
	// usr, err := s.organizationService.CheckUserExists(project.AdminID)
	// if !usr {
	// 	return errors.New("admin user does not exist")
	// }
	// if err != nil {
	// 	return errors.New("admin user does not exist")
	// }

	return s.repo.Create(project)
}

func (s *ProjectService) GetProjectByID(id uint) (*models.Project, error) {
	return s.repo.GetByID(id)
}

func (s *ProjectService) UpdateProject(project *models.Project) error { //, adminID uint
	return s.repo.Update(project)
}

func (s *ProjectService) PatchProject(existingOrg *models.Project, updatedFields map[string]interface{}) error {
	return s.repo.Patch(existingOrg, updatedFields)
}

func (s *ProjectService) DeleteProject(project *models.Project) error { //, adminID uint
	// if project.AdminID != adminID {
	// 	return errors.New("only admin can delete the project")
	// }
	return s.repo.Delete(project)
}

func (s *ProjectService) ListProjects() ([]models.Project, error) {
	return s.repo.List()
}
