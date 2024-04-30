package repository

import (
	"errors"
	"os"

	"com.homindolentrahar.rutinkann-api/helper"
	"com.homindolentrahar.rutinkann-api/model"
	"com.homindolentrahar.rutinkann-api/web"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type AuthRepositoryImpl struct {
	Validate *validator.Validate
}

func NewAuthRepositoryImpl(validate *validator.Validate) *AuthRepositoryImpl {
	return &AuthRepositoryImpl{
		Validate: validate,
	}
}

func (a AuthRepositoryImpl) SignIn(database *gorm.DB, request *web.SignInRequest) (*model.User, string, error) {
	var user model.User

	err := database.Where("email = ?", request.Email).First(&user).Error
	compareErr := helper.CompareEncryptedValue(request.Password, user.Password)

	if err != nil {
		return nil, "", err
	}

	if user.Email != request.Email || compareErr != nil {
		return nil, "", errors.New("invalid credential")
	}

	var secretKey = []byte(os.Getenv("APP_SECRET_KEY"))
	userClaim := model.UserClaim{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}
	token, tokenErr := helper.CreateToken(string(secretKey), &userClaim)
	if tokenErr != nil {
		return nil, "", tokenErr
	}

	return &user, token, err
}

func (a AuthRepositoryImpl) Register(database *gorm.DB, request *web.RegisterRequest) (*model.User, string, error) {
	user := model.User{
		Name:     request.Name,
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
	}

	validateErr := a.Validate.Struct(&user)
	if validateErr != nil {
		return nil, "", validateErr
	}

	hashedPass, hashedErr := helper.EncryptValue(user.Password)
	if hashedErr != nil {
		return nil, "", hashedErr
	}
	user.Password = hashedPass

	result := database.Create(&user)
	if result.Error != nil {
		return nil, "", result.Error
	}

	claim := model.UserClaim{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}

	secretKey := os.Getenv("APP_SECRET_KEY")
	token, tokenErr := helper.CreateToken(secretKey, &claim)
	if tokenErr != nil {
		return nil, "", tokenErr
	}

	return &user, token, nil
}
