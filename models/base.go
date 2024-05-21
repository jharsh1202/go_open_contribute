// models/base_model.go

package models

import (
	"time"

	"gorm.io/gorm"
)

// BaseModel for common fields
type BaseModel struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// ActiveModel for active/inactive status
type ActiveBaseModel struct {
	IsActive bool `gorm:"default:true"`
}

// AuditableModel for created by and updated by fields
type CreatedUpdatedByBaseModel struct {
	CreatedBy uint
	UpdatedBy uint
}
