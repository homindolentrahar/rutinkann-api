package controller

import (
	"com.homindolentrahar.rutinkann-api/helper"
	"com.homindolentrahar.rutinkann-api/model"
	"com.homindolentrahar.rutinkann-api/repository"
	"encoding/json"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type LogControllerImpl struct {
	Repository repository.LogRepository
	Database   *gorm.DB
}

func NewLogController(repository repository.LogRepository, database *gorm.DB) *LogControllerImpl {
	return &LogControllerImpl{
		Repository: repository,
		Database:   database,
	}
}

func (controller *LogControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request) {
	db := controller.Database.WithContext(request.Context())
	helper.PanicIfError(db.Error)

	logs, resultErr := controller.Repository.FindAll(db)
	response := helper.HandleErrorBaseResponse[[]model.Log](writer, http.StatusOK, &logs, resultErr, "Success getting all logs")

	encoder := json.NewEncoder(writer)
	encodeErr := encoder.Encode(response)
	helper.PanicIfError(encodeErr)
}

func (controller *LogControllerImpl) FindById(writer http.ResponseWriter, request *http.Request) {
	pathId := request.PathValue("id")
	id, err := strconv.Atoi(pathId)
	helper.PanicIfError(err)

	db := controller.Database.WithContext(request.Context())
	helper.PanicIfError(db.Error)

	log, resultErr := controller.Repository.FindById(db, id)
	response := helper.HandleErrorBaseResponse[model.Log](writer, http.StatusOK, log, resultErr, "Success getting log by ID")

	encoder := json.NewEncoder(writer)
	encodeErr := encoder.Encode(response)
	helper.PanicIfError(encodeErr)
}

func (controller *LogControllerImpl) Create(writer http.ResponseWriter, request *http.Request) {
	var reqBody model.Log
	decoder := json.NewDecoder(request.Body)
	decodeErr := decoder.Decode(&reqBody)
	helper.PanicIfError(decodeErr)

	db := controller.Database.WithContext(request.Context())
	helper.PanicIfError(db.Error)

	result, err := controller.Repository.Create(db, reqBody)
	response := helper.HandleErrorBaseResponse[model.Log](writer, http.StatusCreated, result, err, "Success creating log")

	encoder := json.NewEncoder(writer)
	encodeErr := encoder.Encode(response)
	helper.PanicIfError(encodeErr)
}

func (controller *LogControllerImpl) Update(writer http.ResponseWriter, request *http.Request) {
	pathId := request.PathValue("id")
	id, pathErr := strconv.Atoi(pathId)
	helper.PanicIfError(pathErr)

	var reqBody model.Log
	decoder := json.NewDecoder(request.Body)
	decodeErr := decoder.Decode(&reqBody)
	helper.PanicIfError(decodeErr)

	reqBody.Id = id

	db := controller.Database.WithContext(request.Context())
	helper.PanicIfError(db.Error)

	result, err := controller.Repository.Update(db, reqBody)
	response := helper.HandleErrorBaseResponse[[]model.Log](writer, http.StatusOK, &result, err, "Success updating log")

	encoder := json.NewEncoder(writer)
	encodeErr := encoder.Encode(response)
	helper.PanicIfError(encodeErr)
}

func (controller *LogControllerImpl) Delete(writer http.ResponseWriter, request *http.Request) {
	pathId := request.PathValue("id")
	id, pathErr := strconv.Atoi(pathId)
	helper.PanicIfError(pathErr)

	db := controller.Database.WithContext(request.Context())
	helper.PanicIfError(db.Error)

	err := controller.Repository.Delete(db, id)
	response := helper.HandleErrorBaseResponse[any](writer, http.StatusOK, nil, err, "Success deleting log")

	encoder := json.NewEncoder(writer)
	encodeErr := encoder.Encode(response)
	helper.PanicIfError(encodeErr)
}
