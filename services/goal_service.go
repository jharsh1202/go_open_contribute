// services/project_service.go
package services

import (
	"errors"
	"open-contribute/models"
	"open-contribute/repositories"
)

type GoalService struct {
	repo           *repositories.GoalRepository
	projectService ProjectService
}

func NewGoalService(repo repositories.GoalRepository, projectService ProjectService) GoalService {
	return GoalService{
		repo:           &repo,
		projectService: projectService,
	}
}

func (s GoalService) CheckUserExists(userID uint) (bool, error) {
	return s.projectService.CheckUserExists(userID)
}

func (s *GoalService) GetUserByID(userID uint) (*models.User, error) {
	//  log.Printf("User ID %v", userID)
	return s.projectService.GetUserByID(userID)
}

// func (s *GoalService) GetOrganizationByID(userID uint) (*models.Organization, error) {
// 	//  log.Printf("User ID %v", userID)
// 	return s.organizationService.GetOrganizationByID(userID)
// }

func (s *GoalService) GetProjectByID(projectID uint) (*models.Project, error) {
	// log.Printf("Project ID - %v", projectID)
	return s.projectService.GetProjectByID(projectID)
}

func (s *GoalService) CreateGoal(project *models.Goal) error {
	usr, err := s.projectService.CheckUserExists(project.OwnerID)
	if !usr {
		return errors.New("admin user does not exist")
	}
	if err != nil {
		return errors.New("admin user does not exist")
	}

	return s.repo.Create(project)
}

func (s *GoalService) GetGoalByID(id uint) (*models.Goal, error) {
	return s.repo.GetByID(id)
}

func (s *GoalService) UpdateGoal(project *models.Goal) error { //, adminID uint
	return s.repo.Update(project)
}

func (s *GoalService) PatchGoal(existingOrg *models.Goal, updatedFields map[string]interface{}) error {
	return s.repo.Patch(existingOrg, updatedFields)
}

func (s *GoalService) DeleteGoal(project *models.Goal) error { //, adminID uint
	// if project.AdminID != adminID {
	// 	return errors.New("only admin can delete the project")
	// }
	return s.repo.Delete(project)
}

func (s *GoalService) ListGoals() ([]models.Goal, error) {
	return s.repo.List()
}
