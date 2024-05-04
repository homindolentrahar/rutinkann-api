package db

import (
	"fmt"
	"os"
	"time"

	"com.homindolentrahar.rutinkann-api/model"

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
	dbUser := os.Getenv("DATABASE_USER")
	dbPassword := os.Getenv("DATABASE_PASSWORD")
	dbHost := os.Getenv("DATABASE_HOST")

	connection := fmt.Sprintf("postgres://%s:%s@%s/?sslmode=disable", dbUser, dbPassword, dbHost)

	return &PostgresStorage{
		Connection: connection,
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

	migrateErr := db.AutoMigrate(&model.Routine{}, &model.Log{}, &model.User{})
	if migrateErr != nil {
		return nil, migrateErr
	}

	return db, nil
}
