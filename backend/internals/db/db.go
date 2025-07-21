package db

import (
	"log"
	"os"
	"time"

	"github.com/Diaku49/FoodOrderSystem/backend/internals/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect() *gorm.DB {
	dsn := os.Getenv("DATABASE_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	sqlDb, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get generic database object: %v", err)
	}
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(80)
	sqlDb.SetConnMaxLifetime(time.Hour)

	// Checking the connection
	if err := sqlDb.Ping(); err != nil {
		log.Fatalf("database ping failed: %v", err)
	}

	return db
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.User{},
		&model.Video{},
		&model.VideoVariant{},
	)
}
