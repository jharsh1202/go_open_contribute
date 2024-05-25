package models

type Task struct {
	BaseModel
	ActiveBaseModel
	CreatedUpdatedByBaseModel
	Name    string `gorm:"not null"`
	GoalID  uint
	Goal    Goal
	OwnerID uint
	Owner   User
}
