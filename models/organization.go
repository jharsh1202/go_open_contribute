package models

type Organization struct {
	BaseModel
	ActiveBaseModel
	CreatedUpdatedByBaseModel
	Name    string `gorm:"unique;not null"`
	AdminID uint   //Change later
	// Admin   User   `gorm:"foreignKey:AdminID"`
	// Admins  []User `gorm:"hasOne:Admin"`
	// Members []User `gorm:"many2many:user_organizations;"`
}
