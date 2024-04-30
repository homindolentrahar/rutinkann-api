package controller

import (
	"encoding/json"
	"net/http"

	"com.homindolentrahar.rutinkann-api/helper"
	"com.homindolentrahar.rutinkann-api/repository"
	"com.homindolentrahar.rutinkann-api/web"
	"gorm.io/gorm"
)

type AuthControllerImpl struct {
	Repository repository.AuthRepository
	Database   *gorm.DB
}

func NewAuthControllerImpl(repository repository.AuthRepository, database *gorm.DB) *AuthControllerImpl {
	return &AuthControllerImpl{Repository: repository, Database: database}
}

func (a AuthControllerImpl) SignIn(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (a AuthControllerImpl) Register(writer http.ResponseWriter, request *http.Request) {
	var reqBody web.RegisterRequest
	decoder := json.NewDecoder(request.Body)
	decodeErr := decoder.Decode(&reqBody)
	helper.PanicIfError(decodeErr)

	database := a.Database.WithContext(request.Context())
	helper.PanicIfError(database.Error)

	user, token, err := a.Repository.Register(database, &reqBody)
	authResponse := web.AuthResponse{
		User:  user,
		Token: token,
	}
	response := helper.HandleBaseAuthResponse(writer, &authResponse, err, helper.BaseResponseConf{
		SuccessStatusCode: http.StatusCreated,
		SuccessMessage:    "Success registering user",
	})

	encoder := json.NewEncoder(writer)
	encodeErr := encoder.Encode(response)
	helper.PanicIfError(encodeErr)
}
