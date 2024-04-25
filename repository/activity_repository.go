package repository

import (
	"gorm.io/gorm"

	"com.homindolentrahar.rutinkann-api/model"
)

type ActivityRepository interface {
	FindAll(database *gorm.DB) ([]model.Activity, error)
	FindById(database *gorm.DB, id int) (*model.Activity, error)
	Create(database *gorm.DB, activity model.Activity) (*model.Activity, error)
	Update(database *gorm.DB, activity model.Activity) ([]model.Activity, error)
	Delete(database *gorm.DB, id int) error
}
