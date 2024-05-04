package repository

import (
	"com.homindolentrahar.rutinkann-api/db"
	"com.homindolentrahar.rutinkann-api/model"
)

type RoutineRepository interface {
	FindAll(pagination *db.Pagination) ([]model.Routine, int64, error)
	FindById(id int) (*model.Routine, error)
	Create(activity model.Routine) (*model.Routine, error)
	Update(activity model.Routine) ([]model.Routine, error)
	Delete(id int) error
}
