// services/project_service.go
package services

import (
	"errors"
	"open-contribute/models"
	"open-contribute/repositories"
)

type TaskService struct {
	repo        *repositories.TaskRepository
	goalService GoalService
}

func NewTaskService(repo repositories.TaskRepository, goalService GoalService) TaskService {
	return TaskService{
		repo:        &repo,
		goalService: goalService,
	}
}

// func (s TaskService) CheckUserExists(userID uint) (bool, error) {
// 	return s.organizationService.CheckUserExists(userID)
// }

func (s *TaskService) GetUserByID(userID uint) (*models.User, error) {
	//  log.Printf("User ID %v", userID)
	return s.goalService.GetUserByID(userID)
}

func (s *TaskService) GetGoalByID(userID uint) (*models.Goal, error) {
	//  log.Printf("User ID %v", userID)
	return s.goalService.GetGoalByID(userID)
}

// func (s *TaskService) GetOrganizationByID(userID uint) (*models.Organization, error) {
// 	//  log.Printf("User ID %v", userID)
// 	return s.organizationService.GetOrganizationByID(userID)
// }

func (s *TaskService) GetProjectByID(projectID uint) (*models.Project, error) {
	// log.Printf("Project ID - %v", projectID)
	return s.goalService.GetProjectByID(projectID)
}

func (s *TaskService) CreateTask(project *models.Task) error {
	usr, err := s.goalService.CheckUserExists(project.OwnerID)
	if !usr {
		return errors.New("admin user does not exist")
	}
	if err != nil {
		return errors.New("admin user does not exist")
	}

	return s.repo.Create(project)
}

func (s *TaskService) GetTaskByID(id uint) (*models.Task, error) {
	return s.repo.GetByID(id)
}

func (s *TaskService) UpdateTask(task *models.Task) error { //, adminID uint
	return s.repo.Update(task)
}

func (s *TaskService) PatchTask(existingOrg *models.Task, updatedFields map[string]interface{}) error {
	return s.repo.Patch(existingOrg, updatedFields)
}

func (s *TaskService) DeleteTask(project *models.Task) error { //, adminID uint
	// if project.AdminID != adminID {
	// 	return errors.New("only admin can delete the project")
	// }
	return s.repo.Delete(project)
}

func (s *TaskService) ListTasks() ([]models.Task, error) {
	return s.repo.List()
}
