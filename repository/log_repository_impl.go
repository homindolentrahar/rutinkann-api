package repository

import (
	"com.homindolentrahar.rutinkann-api/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type LogRepositoryImpl struct{}

func NewLogRepository() *LogRepositoryImpl {
	return &LogRepositoryImpl{}
}

func (repo *LogRepositoryImpl) FindAll(database *gorm.DB) ([]model.Log, error) {
	var logs []model.Log

	err := database.Order("created_at desc").Find(&logs).Error
	if err != nil {
		return nil, err
	}

	return logs, nil
}

func (repo *LogRepositoryImpl) FindById(database *gorm.DB, id int) (*model.Log, error) {
	var log model.Log

	err := database.First(&log, id).Error
	if err != nil {
		return nil, err
	}

	return &log, nil
}

func (repo *LogRepositoryImpl) Create(database *gorm.DB, log model.Log) (*model.Log, error) {
	result := database.Create(&log)

	if result.Error != nil {
		return nil, result.Error
	}

	return &log, nil
}

func (repo *LogRepositoryImpl) Update(database *gorm.DB, log model.Log) ([]model.Log, error) {
	var logs []model.Log

	result := database.Model(&logs).Clauses(clause.Returning{}).Where("id = ?", log.Id).Updates(&log)

	if result.RowsAffected <= 0 {
		return nil, gorm.ErrRecordNotFound
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return logs, nil
}

func (repo *LogRepositoryImpl) Delete(database *gorm.DB, id int) error {
	result := database.Delete(&model.Log{}, id)

	if result.RowsAffected <= 0 {
		return gorm.ErrRecordNotFound
	}

	if result.Error != nil {
		return result.Error
	}

	return nil
}
