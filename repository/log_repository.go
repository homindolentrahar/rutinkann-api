package repository

import (
	"com.homindolentrahar.rutinkann-api/db"
	"com.homindolentrahar.rutinkann-api/model"
)

type LogRepository interface {
	FindAll(pagination *db.Pagination) ([]model.Log, int64, error)
	FindById(id int) (*model.Log, error)
	Create(log model.Log) (*model.Log, error)
	Update(log model.Log) ([]model.Log, error)
	Delete(id int) error
}
