package models

type Goal struct {
	BaseModel
	ActiveBaseModel
	CreatedUpdatedByBaseModel
	Name      string `gorm:"not null"`
	ProjectID uint
	Project   Project
	Tasks     []Task
	OwnerID   uint
	Owner     User
}
