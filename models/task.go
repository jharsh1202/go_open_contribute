package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Name   string `gorm:"not null"`
	GoalID uint
	Goal   Goal
}
