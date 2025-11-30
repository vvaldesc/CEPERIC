package config

import (
	"fmt"
	"log"

	"github.com/ceperic/backend/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase(cfg *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBPort,
	)

	// Configurar logger
	logLevel := logger.Silent
	if cfg.Environment == "development" {
		logLevel = logger.Info
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	log.Println("✅ Database connected successfully")
	return db, nil
}

// AutoMigrate ejecuta las migraciones automáticas
func AutoMigrate(db *gorm.DB) error {
	log.Println("🔄 Running database migrations...")
	
	if err := db.AutoMigrate(
		&domain.User{},
		&domain.Document{},
	); err != nil {
		return err
	}
	
	log.Println("✅ Migrations completed successfully")
	return nil
}
