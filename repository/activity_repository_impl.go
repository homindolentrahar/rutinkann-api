package repository

import (
	"com.homindolentrahar.rutinkann-api/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ActivityRepositoryImpl struct{}

func NewActivityRepository() *ActivityRepositoryImpl {
	return &ActivityRepositoryImpl{}
}

func (repo *ActivityRepositoryImpl) FindAll(database *gorm.DB) ([]model.Activity, error) {
	var activities []model.Activity

	err := database.Model(&model.Activity{}).Preload("Logs").Find(&activities).Error
	if err != nil {
		return nil, err
	}

	return activities, nil
}

func (repo *ActivityRepositoryImpl) FindById(database *gorm.DB, id int) (*model.Activity, error) {
	var activity model.Activity

	err := database.Model(&model.Activity{}).Preload("Logs").Where("activities.id = ?", id).First(&activity).Error
	if err != nil {
		return nil, err
	}

	return &activity, nil
}

func (repo *ActivityRepositoryImpl) Create(database *gorm.DB, activity model.Activity) (*model.Activity, error) {
	activity.Logs = make([]model.Log, 0)
	result := database.Create(&activity)

	if result.Error != nil {
		return nil, result.Error
	}

	return &activity, nil
}

func (repo *ActivityRepositoryImpl) Update(database *gorm.DB, activity model.Activity) ([]model.Activity, error) {
	var activities []model.Activity

	result := database.Model(&activities).Clauses(clause.Returning{}).Where("id = ?", activity.ID).Updates(&activity)

	if result.RowsAffected <= 0 {
		return nil, gorm.ErrRecordNotFound
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return activities, nil
}

func (repo *ActivityRepositoryImpl) Delete(database *gorm.DB, id int) error {
	result := database.Delete(&model.Activity{}, id)

	if result.RowsAffected <= 0 {
		return gorm.ErrRecordNotFound
	}

	if result.Error != nil {
		return result.Error
	}

	return nil
}
