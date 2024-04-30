package repository

import (
	"com.homindolentrahar.rutinkann-api/model"
	"com.homindolentrahar.rutinkann-api/web"
	"gorm.io/gorm"
)

type AuthRepository interface {
	SignIn(database *gorm.DB, request *web.SignInRequest) (*model.User, string, error)
	Register(database *gorm.DB, request *web.RegisterRequest) (*model.User, string, error)
}
