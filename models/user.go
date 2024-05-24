package models

import (
	"errors"
	"net/url"
	"regexp"
	"time"

	"gorm.io/gorm"
)

type User struct {
	BaseModel
	ActiveBaseModel
	CreatedUpdatedByBaseModel
	Username  string `gorm:"unique;not null"`
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	SuperUser bool
	// OrganizationID uint // Add this line
	// Organization   Organization `gorm:"foreignKey:OrganizationID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Add this line
}

func (u *User) BeforeSave(_ *gorm.DB) (err error) {
	// Validate username
	if err = validateUsername(u.Username); err != nil {
		return err
	}

	// Validate email
	if err = validateEmail(u.Email); err != nil {
		return err
	}

	// Validate password
	if err = validatePassword(u.Password); err != nil {
		return err
	}

	return nil
}

func validateUsername(username string) error {
	matched, err := regexp.MatchString("^[a-zA-Z0-9]+$", username)
	if err != nil {
		return err
	}
	if !matched {
		return errors.New("username must be alphanumeric")
	}
	return nil
}

func validateEmail(email string) error {
	matched, err := regexp.MatchString(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, email)
	if err != nil {
		return err
	}
	if !matched {
		return errors.New("invalid email format")
	}
	return nil
}

func validatePassword(password string) error {
	// log.Printf("Validating password for user %v", password)
	// Password should be at least 8 characters long
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}
	// Check for at least one number and one letter
	matched, err := regexp.MatchString(`[0-9]`, password)
	if err != nil {
		return err
	}
	if !matched {
		return errors.New("password must contain at least one number")
	}
	matched, err = regexp.MatchString(`[a-zA-Z]`, password)
	if err != nil {
		return err
	}
	if !matched {
		return errors.New("password must contain at least one letter")
	}
	return nil
}

type UserProfile struct {
	BaseModel
	ActiveBaseModel
	CreatedUpdatedByBaseModel
	UserID uint `gorm:"not null;uniqueIndex"`
	User   User `gorm:"foreignKey:UserID;references:ID"`

	FirstName          string
	MiddleName         string `json:"middle_name,omitempty"`
	LastName           string
	DateOfBirth        time.Time `json:"date_of_birth,omitempty"`
	PhoneNumber        string
	AadhaarNumber      string `json:"aadhaar_number,omitempty"` // Correct spelling
	Gender             string `json:"gender,omitempty"`
	SecondaryEmail     string `json:"secondary_email,omitempty"` // Standardize with field name
	LinkedIn           string `json:"linkedin,omitempty"`
	Twitter            string `json:"twitter,omitempty"`
	Instagram          string `json:"instagram,omitempty"`
	StreetAddress1     string `json:"street_address1"`
	StreetAddress2     string `json:"street_address2"`
	City               string
	State              string
	Country            string
	Pincode            string
	ProfilePicture     *url.URL `gorm:"type:varchar(255)"` // Using pointer for optional URL
	Language           string   `json:"language,omitempty"`
	Timezone           string   `json:"timezone,omitempty"`
	MarketingOptIn     bool     `json:"marketing_opt_in"`
	PasswordResetToken string   `gorm:"type:varchar(255)"`
	PhoneVerified      bool     `json:"phone_verified"`
	EmailVerified      bool     `json:"email_verified"`
	AddressVerified    bool     `json:"address_verified"`
	AadhaarVerified    bool     `json:"aadhaar_verified"`
}
