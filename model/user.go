package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int       `json:"id" gorm:"primaryKey;column:id"`
	Name      string    `json:"name" gorm:"column:name" validate:"required"`
	Username  string    `json:"username" gorm:"column:username;unique" validate:"required,min=3"`
	Email     string    `json:"email" gorm:"column:email;unique" validate:"required,email"`
	Password  string    `json:"-" gorm:"column:password" validate:"required,min=8"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt time.Time `json:"-" gorm:"index;column:deleted_at"`
}

func (u *User) BeforeUpdate(_ *gorm.DB) error {
	//hashedPass, err := helper.EncryptValue(u.Password)
	//helper.PanicIfError(err)
	//u.Password = fmt.Sprintf("%x", hashedPass[:])
	return nil
}
