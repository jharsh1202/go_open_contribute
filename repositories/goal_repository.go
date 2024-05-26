// repositories/project_repository.go
package repositories

import (
	"open-contribute/models"

	"gorm.io/gorm"
)

type GoalRepository struct {
	db *gorm.DB
}

func NewGoalRepository(db *gorm.DB) *GoalRepository {
	return &GoalRepository{db: db}
}

func (r *GoalRepository) Create(project *models.Goal) error {
	return r.db.Create(project).Error
}

func (r *GoalRepository) GetByID(id uint) (*models.Goal, error) {
	//  log.Printf("ID %v", id)
	var project models.Goal
	// if err := r.db.Preload("Admin").Preload("Members").First(&project, id).Error; err != nil {
	if err := r.db.Preload("Owner").Preload("Project").First(&project, id).Error; err != nil {
		return nil, err
	}
	return &project, nil
}

func (r *GoalRepository) Update(goal *models.Goal) error {
	// //  log.Printf("Goal: - %v", project.Admin)
	return r.db.Save(&goal).Error
}

func (r *GoalRepository) Delete(project *models.Goal) error {
	return r.db.Delete(project).Error
}

func (r *GoalRepository) List() ([]models.Goal, error) {
	var goals []models.Goal
	// if err := r.db.Preload("Admin").Preload("Members").Find(&projects).Error; err != nil {
	if err := r.db.Find(&goals).Error; err != nil {
		return nil, err
	}
	return goals, nil
}

func (r *GoalRepository) Patch(existingOrg *models.Goal, updatedFields map[string]interface{}) error {
	return r.db.Model(existingOrg).Updates(updatedFields).Error
}
