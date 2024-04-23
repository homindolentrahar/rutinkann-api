package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"com.homindolentrahar.rutinkann-api/helper"
	"com.homindolentrahar.rutinkann-api/model"
)

type ActivityRepositoryImpl struct{}

func NewActivityRepository() *ActivityRepositoryImpl {
	return &ActivityRepositoryImpl{}
}

func (repo *ActivityRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]model.Activity, error) {
	query := `SELECT * FROM activities ORDER BY id`
	rows, err := tx.QueryContext(ctx, query)
	helper.PanicIfError(err)

	if err != nil {
		return nil, errors.New("500: " + err.Error())
	}

	defer rows.Close()

	activities := make([]model.Activity, 0)
	for rows.Next() {
		activity := model.Activity{}
		err := rows.Scan(&activity.Id, &activity.Name, &activity.Desc, &activity.StreakCount, &activity.CreatedAt, &activity.UpdatedAt)
		activity.Logs = []model.LogModel{}

		activities = append(activities, activity)

		if err != nil {
			return nil, errors.New("500: " + err.Error())
		}
	}

	return activities, nil
}

func (repo *ActivityRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (*model.Activity, error) {
	query := `SELECT * FROM activities WHERE id=$1`
	row, err := tx.QueryContext(ctx, query, id)
	activity := model.Activity{}
	helper.PanicIfError(err)

	if err != nil {
		return nil, errors.New("500: " + err.Error())
	}

	defer row.Close()

	if row.Next() {
		err := row.Scan(&activity.Id, &activity.Name, &activity.Desc, &activity.StreakCount, &activity.CreatedAt, &activity.UpdatedAt)
		helper.PanicIfError(err)

		return &activity, nil
	} else {
		return nil, errors.New("404: activity not found")
	}
}

func (repo *ActivityRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, activity model.Activity) (*model.Activity, error) {
	query := `INSERT INTO activities (name, description, streak_count) VALUES ($1, $2, $3) RETURNING id`
	var id int
	err := tx.QueryRowContext(ctx, query, activity.Name, activity.Desc, activity.StreakCount).Scan(&id)
	helper.PanicIfError(err)

	if err != nil {
		return nil, errors.New("500: " + err.Error())
	}

	activity.Id = id
	activity.CreatedAt = time.Now()
	activity.UpdatedAt = time.Now()
	activity.Logs = []model.LogModel{}

	return &activity, nil
}

func (repo *ActivityRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, activity model.Activity) (*model.Activity, error) {
	currentTime := time.Now()
	updateQuery := `UPDATE activities SET name=$1, description=$2, streak_count=$3, updated_at=$4 WHERE id=$5`
	_, err := tx.ExecContext(ctx, updateQuery, activity.Name, activity.Desc, activity.StreakCount, currentTime, activity.Id)
	helper.PanicIfError(err)

	if err != nil {
		return nil, errors.New("500: " + err.Error())
	}

	activity.UpdatedAt = currentTime
	activity.Logs = []model.LogModel{}

	return &activity, nil
}

func (repo *ActivityRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	query := `DELETE FROM activities WHERE id=$1`
	_, err := tx.ExecContext(ctx, query, id)
	helper.PanicIfError(err)

	if err != nil {
		return err
	}

	return nil
}
