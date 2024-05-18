package models

import "gorm.io/gorm"

type Goal struct {
	gorm.Model
	Name      string `gorm:"not null"`
	ProjectID uint
	Project   Project
	Tasks     []Task
}
