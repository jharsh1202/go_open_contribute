// services/user_service.go
package services

import (
	"errors"
	"log"
	"open-contribute/models"
	"open-contribute/repositories"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(username, email, password string) error
	LoginUser(username, password string) (*models.User, error)
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{userRepository: userRepo}
}

func (s *userService) RegisterUser(username, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &models.User{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
	}

	log.Printf("user: %+v", user)
	return s.userRepository.CreateUser(user)
}

func (s *userService) LoginUser(username, password string) (*models.User, error) {
	user, err := s.userRepository.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}