package models

type Project struct {
	BaseModel
	ActiveBaseModel
	CreatedUpdatedByBaseModel
	Name           string `gorm:"not null"`
	OrganizationID uint
	Organization   Organization
	Goals          []Goal
	AdminID        uint
	Admin          User
}
