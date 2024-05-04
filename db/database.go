package db

import (
	"fmt"
	"os"
	"time"

	"com.homindolentrahar.rutinkann-api/helper"
	"com.homindolentrahar.rutinkann-api/model"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type StorageConfig struct {
	Connection string
	Driver     string
}

func ConnectStorage(config StorageConfig) *gorm.DB {
	db, err := gorm.Open(postgres.Open(config.Connection), &gorm.Config{})
	sqlDB, _ := db.DB()
	helper.PanicIfError(err)

	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetConnMaxIdleTime(10 * time.Minute)

	migrateErr := db.AutoMigrate(&model.Routine{}, &model.Log{}, &model.User{})
	helper.PanicIfError(migrateErr)

	return db
}

func ConnectPostgresStorage() *gorm.DB {
	dbUser := os.Getenv("DATABASE_USER")
	dbPassword := os.Getenv("DATABASE_PASSWORD")
	dbHost := os.Getenv("DATABASE_HOST")
	connection := fmt.Sprintf("postgres://%s:%s@%s/?sslmode=disable", dbUser, dbPassword, dbHost)
	config := StorageConfig{
		Connection: connection,
		Driver:     "postgres",
	}

	return ConnectStorage(config)
}
