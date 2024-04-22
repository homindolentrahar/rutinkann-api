package model

import "time"

type Activity struct {
	Id          int        `json:"id" db:"id"`
	Name        string     `json:"name" db:"name"`
	Desc        string     `json:"description" db:"description"`
	StreakCount int        `json:"streak_count" db:"streak_count"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`
	Logs        []LogModel `json:"logs"`
}
