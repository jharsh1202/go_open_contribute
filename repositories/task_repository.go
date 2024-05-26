// repositories/project_repository.go
package repositories

import (
	"open-contribute/models"

	"gorm.io/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) Create(project *models.Task) error {
	return r.db.Create(project).Error
}

func (r *TaskRepository) GetByID(id uint) (*models.Task, error) {
	//  log.Printf("ID %v", id)
	var project models.Task
	// if err := r.db.Preload("Admin").Preload("Members").First(&project, id).Error; err != nil {
	if err := r.db.Preload("Owner").Preload("Goal").First(&project, id).Error; err != nil {
		return nil, err
	}
	return &project, nil
}

func (r *TaskRepository) Update(goal *models.Task) error {
	// //  log.Printf("Task: - %v", project.Admin)
	return r.db.Save(&goal).Error
}

func (r *TaskRepository) Delete(project *models.Task) error {
	return r.db.Delete(project).Error
}

func (r *TaskRepository) List() ([]models.Task, error) {
	var goals []models.Task
	if err := r.db.Preload("Owner").Preload("Goal").Find(&goals).Error; err != nil {
		// if err := r.db.Find(&goals).Error; err != nil {
		return nil, err
	}
	return goals, nil
}

func (r *TaskRepository) Patch(existingOrg *models.Task, updatedFields map[string]interface{}) error {
	return r.db.Model(existingOrg).Updates(updatedFields).Error
}
