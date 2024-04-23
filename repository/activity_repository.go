package repository

import (
	"context"
	"database/sql"

	"com.homindolentrahar.rutinkann-api/model"
)

type ActivityRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) ([]model.Activity, error)
	FindById(ctx context.Context, tx *sql.Tx, id int) (*model.Activity, error)
	Create(ctx context.Context, tx *sql.Tx, activity model.Activity) (*model.Activity, error)
	Update(ctx context.Context, tx *sql.Tx, activity model.Activity) (*model.Activity, error)
	Delete(ctx context.Context, tx *sql.Tx, id int) error
}
