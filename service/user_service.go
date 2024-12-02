package service

import (
	"lumoshive-be-chap41/models"
	"lumoshive-be-chap41/repository"
)

type UserService interface {
	GetUser(id int) (models.User, error)
	UpdateUser(user models.User) error
	GetUserRedeem(id int) (*models.User, error)
	GetUserUsage(id int) (*models.User, error)
	// CreateUser(user models.User) error
	// DeleteUser(id uint) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo}
}

// func (s *userService) CreateUser(user models.User) error {
// 	return s.repo.CreateUser(user)
// }

func (s *userService) GetUser(id int) (models.User, error) {
	return s.repo.GetUser(id)
}

func (s *userService) UpdateUser(user models.User) error {
	return s.repo.UpdateUser(user)
}

// func (s *userService) DeleteUser(id uint) error {
// 	return s.repo.DeleteUser(id)
// }

func (s *userService) GetUserRedeem(id int) (*models.User, error) {
	return s.repo.GetUserRedeem(id)
}

func (s *userService) GetUserUsage(id int) (*models.User, error) {
	return s.repo.GetUserUsage(id)
}
