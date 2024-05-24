// services/user_service.go
package services

import (
	"errors"
	"open-contribute/models"
	"open-contribute/repositories"
	"open-contribute/utils"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(username, email, password string, superUser bool) error
	LoginUser(username, password string) (*models.User, *string, error)
	GetUserByID(id uint) (*models.User, error)
	GetUsers() ([]models.User, error)
	CheckUserExists(id uint) (bool, error)
	UpdateUser(user *models.User) error
	PatchUser(existingUser *models.User, updatedFields map[string]interface{}) error
	DeleteUser(user *models.User) error
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{userRepository: userRepo}
}

func (s *userService) RegisterUser(username, email, password string, superUser bool) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &models.User{
		Username:  username,
		Email:     email,
		Password:  string(hashedPassword),
		SuperUser: superUser,
	}

	return s.userRepository.CreateUser(user)
}

func (s *userService) LoginUser(username, password string) (*models.User, *string, error) {
	user, err := s.userRepository.GetUserByUsername(username)
	if err != nil {
		return nil, nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, nil, errors.New("invalid credentials")
	}

	jwt, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return nil, nil, err
	}

	return user, &jwt, nil
}

func (s *userService) GetUserByID(id uint) (*models.User, error) {
	return s.userRepository.GetUserByID(id)
}

func (s *userService) GetUsers() ([]models.User, error) {
	return s.userRepository.GetUsers()
}

func (s *userService) CheckUserExists(id uint) (bool, error) {
	user, err := s.userRepository.GetUserByID(id)
	if err != nil {
		return false, err
	}

	return user != nil, nil
}

func (s *userService) UpdateUser(user *models.User) error { //, adminID uint
	return s.userRepository.UpdateUser(user)
}

func (s *userService) DeleteUser(user *models.User) error { //, adminID uint
	return s.userRepository.DeleteUser(user)
}

func (s *userService) PatchUser(existingOrg *models.User, updatedFields map[string]interface{}) error {
	return s.userRepository.PatchUser(existingOrg, updatedFields)
}
