package db

import (
	"com.homindolentrahar.rutinkann-api/model"
	"time"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage interface {
	Connect() (*gorm.DB, error)
}

type PostgresStorage struct {
	Connection string
	Driver     string
}

func NewPostgresStorage() *PostgresStorage {
	return &PostgresStorage{
		Connection: "postgres://postgres:root@localhost/rutinkann?sslmode=disable",
		Driver:     "postgres",
	}
}

func (ps *PostgresStorage) Connect() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(ps.Connection), &gorm.Config{})
	sqlDB, _ := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetConnMaxIdleTime(10 * time.Minute)

	migrateErr := db.AutoMigrate(&model.Activity{}, &model.Log{})
	if migrateErr != nil {
		return nil, migrateErr
	}

	return db, nil
}
