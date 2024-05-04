package repository

import (
	"com.homindolentrahar.rutinkann-api/db"
	"com.homindolentrahar.rutinkann-api/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type LogRepositoryImpl struct {
	Database *gorm.DB
}

func NewLogRepository(database *gorm.DB) *LogRepositoryImpl {
	return &LogRepositoryImpl{Database: database}
}

func (repo *LogRepositoryImpl) FindAll(pagination *db.Pagination) ([]model.Log, int64, error) {
	var logs []model.Log
	var count int64

	err := repo.Database.Scopes(db.Paginate(pagination)).Find(&logs).Error
	if err != nil {
		return nil, 0, err
	}
	repo.Database.Model(&model.Log{}).Count(&count)

	return logs, count, nil
}

func (repo *LogRepositoryImpl) FindById(id int) (*model.Log, error) {
	var log model.Log

	err := repo.Database.First(&log, id).Error
	if err != nil {
		return nil, err
	}

	return &log, nil
}

func (repo *LogRepositoryImpl) Create(log model.Log) (*model.Log, error) {
	result := repo.Database.Create(&log)

	if result.Error != nil {
		return nil, result.Error
	}

	return &log, nil
}

func (repo *LogRepositoryImpl) Update(log model.Log) ([]model.Log, error) {
	var logs []model.Log

	result := repo.Database.Model(&logs).Clauses(clause.Returning{}).Where("id = ?", log.Id).Updates(&log)

	if result.RowsAffected <= 0 {
		return nil, gorm.ErrRecordNotFound
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return logs, nil
}

func (repo *LogRepositoryImpl) Delete(id int) error {
	result := repo.Database.Delete(&model.Log{}, id)

	if result.RowsAffected <= 0 {
		return gorm.ErrRecordNotFound
	}

	if result.Error != nil {
		return result.Error
	}

	return nil
}
