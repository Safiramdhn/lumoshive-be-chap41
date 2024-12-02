package database

import (
	"fmt"
	"log"
	"lumoshive-be-chap41/models"
	"time"

	"gorm.io/gorm"
)

func SeedDatabase(db *gorm.DB) error {
	log.Println("Starting user seeding...")
	if err := seedUsers(db); err != nil {
		return fmt.Errorf("user seeding failed: %w", err)
	}

	log.Println("Starting voucher seeding...")
	if err := voucherSeeder(db); err != nil {
		return fmt.Errorf("voucher seeding failed: %w", err)
	}

	log.Println("Starting redeem seeding...")
	if err := redeemSeeder(db); err != nil {
		return fmt.Errorf("redeem seeding failed: %w", err)
	}

	log.Println("Database seeding completed successfully.")
	return nil
}

func voucherSeeder(db *gorm.DB) error {
	var count int64
	if err := db.Model(&models.Voucher{}).Count(&count).Error; err != nil {
		return fmt.Errorf("failed to count existing shipping data: %w", err)
	}

	if count > 0 {
		log.Println("Shipping data already initialized, skipping seeder.")
		return nil
	}
	// Seed Vouchers
	vouchers := []models.Voucher{
		{
			VoucherName:     "Free Shipping Java",
			VoucherCode:     "FREESHIPJAVA",
			VoucherType:     "ecommerce",
			Description:     "Free shipping for orders in Java island.",
			VoucherCategory: "free_shipping",
			DiscountAmount:  10.00,
			MinPurchase:     50000,
			PaymentMethod:   "credit_card",
			StartDate:       time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			EndDate:         time.Date(2024, 12, 31, 23, 59, 59, 0, time.UTC),
			ApplicableAreas: `["Jakarta", "Surabaya", "Bandung"]`,
			MinRatePoint:    50,
			Quantity:        100,
		},
		{
			VoucherName:     "10% Off Electronics",
			VoucherCode:     "ELEC10",
			VoucherType:     "ecommerce",
			Description:     "10% discount on electronics above Rp500,000.",
			VoucherCategory: "discount",
			DiscountAmount:  10.00,
			MinPurchase:     500000,
			PaymentMethod:   "debit_card",
			StartDate:       time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC),
			EndDate:         time.Date(2024, 11, 30, 23, 59, 59, 0, time.UTC),
			ApplicableAreas: `["Jakarta", "Medan"]`,
			MinRatePoint:    100,
			Quantity:        50,
		},
	}
	if err := db.Create(&vouchers).Error; err != nil {
		return err
	}
	return nil
}

func redeemSeeder(db *gorm.DB) error {
	// Fetch Voucher IDs
	var count int64
	if err := db.Model(&models.Redeem{}).Count(&count).Error; err != nil {
		return fmt.Errorf("failed to count existing shipping data: %w", err)
	}

	if count > 0 {
		log.Println("Shipping data already initialized, skipping seeder.")
		return nil
	}
	var voucher1, voucher2 models.Voucher
	if err := db.First(&voucher1, "voucher_code = ?", "ELEC10").Error; err != nil {
		return err
	}
	if err := db.First(&voucher2, "voucher_code = ?", "FREESHIPJAVA").Error; err != nil {
		return err
	}

	// Seed Redeems
	redeems := []models.Redeem{
		{
			UserID:      1,
			VoucherID:   voucher1.ID,
			VoucherCode: "ELEC10",
			RedeemDate:  time.Date(2024, 3, 15, 10, 30, 0, 0, time.UTC),
		},
		{
			UserID:      1,
			VoucherID:   voucher2.ID,
			VoucherCode: "FREESHIPJAVA",
			RedeemDate:  time.Date(2024, 4, 10, 12, 0, 0, 0, time.UTC),
		},
	}
	if err := db.Create(&redeems).Error; err != nil {
		return err
	}
	return nil
}

func seedUsers(db *gorm.DB) error {
	// Cek apakah data sudah ada
	var count int64
	if err := db.Model(&models.User{}).Count(&count).Error; err != nil {
		return fmt.Errorf("failed to check existing users: %w", err)
	}

	if count > 0 {
		log.Println("Users already exist, skipping seed.")
		return nil
	}

	users := []models.User{
		{
			Name:     "John Doe",
			Email:    "john.doe@example.com",
			Password: "password123",
			Points:   200,
		},
		{
			Name:     "Jane Smith",
			Email:    "jane.smith@example.com",
			Password: "securepassword",
			Points:   150,
		},
		{
			Name:     "Michael Johnson",
			Email:    "michael.johnson@example.com",
			Password: "michaelpassword",
			Points:   50,
		},
	}

	if err := db.Create(&users).Error; err != nil {
		return fmt.Errorf("error seeding users: %w", err)
	}

	return nil
}
