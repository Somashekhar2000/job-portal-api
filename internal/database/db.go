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

func Open() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=1234 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Connection() (*gorm.DB, error) {
	log.Info().Msg("main : Started : Initializing db support")
	db, err := Open()
	if err != nil {
		return nil, fmt.Errorf("connecting to db %w", err)
	}
	pg, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database instance %w ", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err = pg.PingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("database is not connected: %w ", err)
	}
	//db.Migrator().DropTable(&model.User{}, &model.Company{}, &model.Job{})
	err = db.Migrator().AutoMigrate(&model.User{})
	if err != nil {
		return nil, fmt.Errorf("auto migration failed: %w ", err)
	}
	err = db.Migrator().AutoMigrate(&model.Company{})
	if err != nil {
		return nil, fmt.Errorf("auto migration failed: %w ", err)
	}
	err = db.Migrator().AutoMigrate(&model.Job{})
	if err != nil {
		return nil, fmt.Errorf("auto migration failed: %w ", err)
	}
	return db, nil
}
