// repositories/project_repository.go
package repositories

import (
	"open-contribute/models"

	"gorm.io/gorm"
)

type ProjectRepository struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) *ProjectRepository {
	return &ProjectRepository{db: db}
}

func (r *ProjectRepository) Create(project *models.Project) error {
	return r.db.Create(project).Error
}

func (r *ProjectRepository) GetByID(id uint) (*models.Project, error) {
	//  log.Printf("ID %v", id)
	var project models.Project
	// if err := r.db.Preload("Admin").Preload("Members").First(&project, id).Error; err != nil {
	if err := r.db.Preload("Admin").First(&project, id).Error; err != nil {
		return nil, err
	}
	return &project, nil
}

func (r *ProjectRepository) Update(project *models.Project) error {
	// //  log.Printf("Project: - %v", project.Admin)
	return r.db.Save(&project).Error
}

func (r *ProjectRepository) Delete(project *models.Project) error {
	return r.db.Delete(project).Error
}

func (r *ProjectRepository) List() ([]models.Project, error) {
	var projects []models.Project
	// if err := r.db.Preload("Admin").Preload("Members").Find(&projects).Error; err != nil {
	if err := r.db.Find(&projects).Error; err != nil {
		return nil, err
	}
	return projects, nil
}

func (r *ProjectRepository) Patch(existingOrg *models.Project, updatedFields map[string]interface{}) error {
	return r.db.Model(existingOrg).Updates(updatedFields).Error
}
