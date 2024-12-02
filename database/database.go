package database

import (
	"fmt"
	"log"
	"lumoshive-be-chap41/config"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDB(cfg config.Config) (*gorm.DB, error) {
	// Validate configuration
	if cfg.DBHost == "" || cfg.DBPort == "" || cfg.DBUser == "" || cfg.DBName == "" || cfg.DBPassword == "" {
		return nil, fmt.Errorf("database configuration is incomplete")
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Enable color output
		},
	)

	// Open the connection to the database
	db, err := gorm.Open(postgres.Open(makePostgresString(cfg)), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	// Call Migrate function to auto-migrate database schemas
	if err := Migrate(db); err != nil {
		return nil, fmt.Errorf("failed to migrate database: %v", err)
	}

	// Seed initial database data
	if err := SeedDatabase(db); err != nil {
		return nil, fmt.Errorf("failed to seed database: %v", err)
	}

	fmt.Println("Database connected successfully")
	return db, nil
}

// makePostgresString creates the PostgreSQL DSN (Data Source Name)
func makePostgresString(cfg config.Config) string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBName, cfg.DBPassword)
}
