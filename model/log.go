package model

import "time"

type Log struct {
	Id         int       `json:"id" gorm:"primaryKey;column:id"`
	Desc       string    `json:"description" gorm:"column:description"`
	Count      int       `json:"count" gorm:"column:count"`
	CreatedAt  time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt  time.Time `json:"deleted_at" gorm:"index;column:deleted_at"`
	ActivityId int       `json:"activity_id" gorm:"column:activity_id"`
}
