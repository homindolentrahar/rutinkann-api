package model

import "time"

type Routine struct {
	ID          int       `json:"id" gorm:"primaryKey;column:id"`
	Name        string    `json:"name" gorm:"column:name"`
	Desc        string    `json:"description" gorm:"column:description"`
	StreakCount int       `json:"streak_count" gorm:"column:streak_count"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt   time.Time `json:"deleted_at" gorm:"index;column:deleted_at"`
	Logs        []Log     `json:"logs" gorm:"-:migration;foreignKey:activity_id;references:id"`
}
