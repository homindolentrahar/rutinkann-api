package model

import "time"

type LogModel struct {
	Id          int
	Desc        string
	Count       int
	CompletedAt time.Time
}
