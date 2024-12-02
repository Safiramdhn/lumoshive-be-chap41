package database

import (
	"fmt"
	"log"
	"lumoshive-be-chap41/models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.Exec(`
		DO $$ BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'voucher_type') THEN
				CREATE TYPE voucher_type AS ENUM ('ecommerce', 'redeem_point');
			END IF;
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'voucher_category') THEN
				CREATE TYPE voucher_category AS ENUM ('free_shipping', 'discount');
			END IF;
		END $$;
	`).Error
	if err != nil {
		return err
	}

	if err := db.Exec(`CREATE TABLE IF NOT EXISTS migrations (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) UNIQUE,
		applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`).Error; err != nil {
		return fmt.Errorf("failed to create migrations table: %w", err)
	}

	// Define migrations
	models := []struct {
		name  string
		model interface{}
	}{
		{"voucher", models.Voucher{}},
		{"redeem", models.Redeem{}},
		{"usage", models.Usage{}},
		{"user", models.User{}},
	}

	for _, migration := range models {
		var count int64
		err := db.Raw("SELECT COUNT(1) FROM migrations WHERE name = ?", migration.name).Scan(&count).Error
		if err != nil {
			return fmt.Errorf("failed to check migration status for %s: %w", migration.name, err)
		}

		if count > 0 {
			log.Printf("Migration '%s' already applied, skipping.", migration.name)
			continue
		}

		// Run migration
		if err := db.AutoMigrate(migration.model); err != nil {
			return fmt.Errorf("failed to migrate model %T: %w", migration.model, err)
		}

		// Record migration as applied
		if err := db.Exec("INSERT INTO migrations (name) VALUES (?)", migration.name).Error; err != nil {
			return fmt.Errorf("failed to record migration %s: %w", migration.name, err)
		}

		log.Printf("Migration '%s' applied successfully.", migration.name)
	}

	return nil
}
