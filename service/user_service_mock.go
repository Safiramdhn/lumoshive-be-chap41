package service

import (
	"lumoshive-be-chap41/models"

	"github.com/stretchr/testify/mock"
)

// MockUserService is a mock of the UserService interface
type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) GetUser(id int) (models.User, error) {
	args := m.Called(id)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *MockUserService) UpdateUser(user models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserService) GetUserRedeem(id int) (*models.User, error) {
	args := m.Called(id)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserService) GetUserUsage(id int) (*models.User, error) {
	args := m.Called(id)
	return args.Get(0).(*models.User), args.Error(1)
}
