package models

// Organization struct
type Organization struct {
	BaseModel
	ActiveBaseModel
	CreatedUpdatedByBaseModel
	Name    string `gorm:"unique;not null"`
	AdminID uint
	Admin   User `gorm:"foreignKey:AdminID"`
}
