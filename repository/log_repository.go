package repository

import (
	"com.homindolentrahar.rutinkann-api/db"
	"com.homindolentrahar.rutinkann-api/model"
	"gorm.io/gorm"
)

type LogRepository interface {
	FindAll(database *gorm.DB, pagination *db.Pagination) ([]model.Log, int64, error)
	FindById(database *gorm.DB, id int) (*model.Log, error)
	Create(database *gorm.DB, log model.Log) (*model.Log, error)
	Update(database *gorm.DB, log model.Log) ([]model.Log, error)
	Delete(database *gorm.DB, id int) error
}
