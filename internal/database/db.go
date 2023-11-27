package database

import (
	"context"
	"fmt"
	"job-portal-api/internal/model"
	"os"
	"time"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DatabaseConnection() (*gorm.DB, error) {

	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Info().Msg("error in opening database connection")
		return nil, fmt.Errorf("error in opening database connection : %w", err)
	}

	postgresDatabase, err := db.DB()
	if err != nil {
		log.Info().Msg("errorin getting database instance")
		return nil, fmt.Errorf("error in geting database object : %w", err)
	}

	context, cancle := context.WithTimeout(context.Background(), time.Second*5)
	defer cancle()

	err = postgresDatabase.PingContext(context)
	if err != nil {
		log.Info().Msg("dtabase connection not created")
		return nil, fmt.Errorf("database is not connected : %w", err)
	}

	//need auto migrate
	err = db.Migrator().AutoMigrate(&model.User{}, &model.Company{}, &model.Job{})
	if err != nil {
		log.Error().Err(err).Msg("error in creating tables")
		return nil, fmt.Errorf("error in creating tables : %w", err)
	}

	return db, nil
}
