package database

import (
	"context"
	"fmt"
	"project/internal/model"
	"time"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Open initializes a database connection and returns a GORM database instance.
func Open() (*gorm.DB, error) {

	dsn := "host=localhost user=postgres password=1234 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}
	return db, nil
}

// Connection initializes the database connection and performs migrations for relevant models.
func Connection() (*gorm.DB, error) {

	log.Info().Msg("database: Initializing database support")

	db, err := Open()
	if err != nil {
		return nil, fmt.Errorf("database: connection error: %w", err)
	}

	pg, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("database: failed to get database instance: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// Check the database connection status
	err = pg.PingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("database: connection error: %w", err)
	}

	// Perform model migrations
	err = db.Migrator().AutoMigrate(&model.User{})
	if err != nil {
		return nil, fmt.Errorf("database: auto migration failed for User model: %w", err)
	}

	err = db.Migrator().AutoMigrate(&model.Company{})
	if err != nil {
		return nil, fmt.Errorf("database: auto migration failed for Company model: %w", err)
	}

	err = db.Migrator().AutoMigrate(&model.Job{})
	if err != nil {
		return nil, fmt.Errorf("database: auto migration failed for Job model: %w", err)
	}
	return db, nil
}
