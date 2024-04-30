package repository

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"com.homindolentrahar.rutinkann-api/helper"
	"com.homindolentrahar.rutinkann-api/model"
	"com.homindolentrahar.rutinkann-api/web"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
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
	if err != nil || (user.Email != request.Email || helper.CompareEncryptedValue(request.Password, user.Password)) {
		return nil, "", err
	}

	var secretKey = []byte(os.Getenv("APP_SECRET_KEY"))
	userClaim := model.UserClaim{}
	token, tokenErr := createToken(string(secretKey), &userClaim)
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
	user.Password = fmt.Sprintf("%x", hashedPass[:])

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
	token, tokenErr := createToken(secretKey, &claim)
	if tokenErr != nil {
		return nil, "", tokenErr
	}

	return &user, token, nil
}

func createToken(secretKey string, userClaim *model.UserClaim) (string, error) {
	expiredTime := time.Now().Add(2 * time.Minute)
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ID:        strconv.Itoa(userClaim.ID),
		Issuer:    os.Getenv("APP_NAME"),
		Subject:   userClaim.Username,
		ExpiresAt: jwt.NewNumericDate(expiredTime),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	})

	return claims.SignedString([]byte(secretKey))
}
