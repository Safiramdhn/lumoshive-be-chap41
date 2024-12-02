package repository

import (
	"errors"
	"fmt"
	"lumoshive-be-chap41/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

// func (r *UserRepository) CreateUser(user models.User) error {
// 	return r.db.Create(&user).Error
// }

func (r *UserRepository) GetUser(id int) (models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	return user, err
}

func (r *UserRepository) UpdateUser(user models.User) error {
	return r.db.Save(&user).Error
}

// func (r *UserRepository) DeleteUser(id uint) error {
// 	return r.db.Delete(&models.User{}, id).Error
// }

func (r *UserRepository) GetUserRedeem(id int) (*models.User, error) {
	var user models.User
	err := r.db.Preload("Redeems").Where("id = ?", id).First(&user).Error
	return &user, err
}

func (r *UserRepository) GetUserUsage(id int) (*models.User, error) {
	// Defensive check for database initialization
	if r.db == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}

	var user models.User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // No user found, return nil without error
		}
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	// Fetch user usages and associated vouchers
	var usages []models.Usage
	err = r.db.Where("user_id = ?", id).Preload("Voucher").Find(&usages).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get usages for user ID %d: %w", id, err)
	}

	// Associate usages and vouchers with the user
	user.Usages = usages

	return &user, nil
}
