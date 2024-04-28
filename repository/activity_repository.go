package repository

import (
	"com.homindolentrahar.rutinkann-api/db"
	"com.homindolentrahar.rutinkann-api/model"
	"gorm.io/gorm"
)

type ActivityRepository interface {
	FindAll(database *gorm.DB, pagination *db.Pagination) ([]model.Activity, int64, error)
	FindById(database *gorm.DB, id int) (*model.Activity, error)
	Create(database *gorm.DB, activity model.Activity) (*model.Activity, error)
	Update(database *gorm.DB, activity model.Activity) ([]model.Activity, error)
	Delete(database *gorm.DB, id int) error
}
