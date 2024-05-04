package repository

import (
	"com.homindolentrahar.rutinkann-api/db"
	"com.homindolentrahar.rutinkann-api/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type RoutineRepositoryImpl struct {
	Database *gorm.DB
}

func NewRoutineRepository(database *gorm.DB) *RoutineRepositoryImpl {
	return &RoutineRepositoryImpl{
		Database: database,
	}
}

func (repo *RoutineRepositoryImpl) FindAll(pagination *db.Pagination) ([]model.Routine, int64, error) {
	var activities []model.Routine
	var count int64

	err := repo.Database.Scopes(db.Paginate(pagination)).Model(&model.Routine{}).Preload("Logs").Find(&activities).Error
	if err != nil {
		return nil, 0, err
	}

	repo.Database.Model(&model.Routine{}).Count(&count)

	return activities, count, nil
}

func (repo *RoutineRepositoryImpl) FindById(id int) (*model.Routine, error) {
	var activity model.Routine

	err := repo.Database.Model(&model.Routine{}).Preload("Logs").Where("routines.id = ?", id).First(&activity).Error
	if err != nil {
		return nil, err
	}

	return &activity, nil
}

func (repo *RoutineRepositoryImpl) Create(activity model.Routine) (*model.Routine, error) {
	activity.Logs = make([]model.Log, 0)
	result := repo.Database.Create(&activity)

	if result.Error != nil {
		return nil, result.Error
	}

	return &activity, nil
}

func (repo *RoutineRepositoryImpl) Update(activity model.Routine) ([]model.Routine, error) {
	var activities []model.Routine

	result := repo.Database.Model(&activities).Clauses(clause.Returning{}).Where("id = ?", activity.ID).Updates(&activity)

	if result.RowsAffected <= 0 {
		return nil, gorm.ErrRecordNotFound
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return activities, nil
}

func (repo *RoutineRepositoryImpl) Delete(id int) error {
	result := repo.Database.Delete(&model.Routine{}, id)

	if result.RowsAffected <= 0 {
		return gorm.ErrRecordNotFound
	}

	if result.Error != nil {
		return result.Error
	}

	return nil
}
