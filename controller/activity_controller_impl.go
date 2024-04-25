package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"com.homindolentrahar.rutinkann-api/model"
	"com.homindolentrahar.rutinkann-api/repository"
	"com.homindolentrahar.rutinkann-api/web"
	"gorm.io/gorm"

	"com.homindolentrahar.rutinkann-api/helper"
)

type ActivityControllerImpl struct {
	Repository repository.ActivityRepository
	Database   *gorm.DB
}

func NewActivityController(repository repository.ActivityRepository, database *gorm.DB) *ActivityControllerImpl {
	return &ActivityControllerImpl{Repository: repository, Database: database}
}

func (controller *ActivityControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request) {
	db := controller.Database.WithContext(request.Context())
	helper.PanicIfError(db.Error)

	var response web.BaseResponse[[]model.Activity]
	activities, resultErr := controller.Repository.FindAll(db)

	if resultErr != nil {
		var statusCode int

		switch {
		case strings.Contains(resultErr.Error(), "404"):
			writer.WriteHeader(404)
			statusCode = http.StatusNotFound
		case strings.Contains(resultErr.Error(), "500"):
			writer.WriteHeader(500)
			statusCode = http.StatusInternalServerError
		default:
			writer.WriteHeader(500)
			statusCode = http.StatusInternalServerError
		}

		response = web.BaseResponse[[]model.Activity]{
			Status:  statusCode,
			Message: resultErr.Error(),
			Data:    nil,
		}
	} else {
		writer.WriteHeader(http.StatusOK)
		response = web.BaseResponse[[]model.Activity]{
			Status:  http.StatusOK,
			Message: "Success getting all activities",
			Data:    activities,
		}
	}

	encoder := json.NewEncoder(writer)
	encodeErr := encoder.Encode(response)
	helper.PanicIfError(encodeErr)
}

func (controller *ActivityControllerImpl) FindById(writer http.ResponseWriter, request *http.Request) {
	pathId := request.PathValue("id")
	id, err := strconv.Atoi(pathId)
	helper.PanicIfError(err)

	db := controller.Database.WithContext(request.Context())
	helper.PanicIfError(db.Error)

	var response web.BaseResponse[*model.Activity]
	activity, resultErr := controller.Repository.FindById(db, id)

	if resultErr != nil {
		var statusCode int

		switch {
		case strings.Contains(resultErr.Error(), "not found"):
			writer.WriteHeader(404)
			statusCode = http.StatusNotFound
		default:
			writer.WriteHeader(500)
			statusCode = http.StatusInternalServerError
		}

		response = web.BaseResponse[*model.Activity]{
			Status:  statusCode,
			Message: resultErr.Error(),
			Data:    nil,
		}
	} else {
		writer.WriteHeader(http.StatusOK)
		response = web.BaseResponse[*model.Activity]{
			Status:  http.StatusOK,
			Message: "Success getting activity by id",
			Data:    activity,
		}
	}

	encoder := json.NewEncoder(writer)
	encodeErr := encoder.Encode(response)
	helper.PanicIfError(encodeErr)
}

func (controller *ActivityControllerImpl) Create(writer http.ResponseWriter, request *http.Request) {

	var reqBody model.Activity
	decoder := json.NewDecoder(request.Body)
	decodeErr := decoder.Decode(&reqBody)
	helper.PanicIfError(decodeErr)

	database := controller.Database.WithContext(request.Context())
	helper.PanicIfError(database.Error)

	var response web.BaseResponse[*model.Activity]
	activity, resultError := controller.Repository.Create(database, reqBody)

	if resultError != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		response = web.BaseResponse[*model.Activity]{
			Status:  http.StatusInternalServerError,
			Message: resultError.Error(),
			Data:    nil,
		}
	} else {
		writer.WriteHeader(http.StatusCreated)
		response = web.BaseResponse[*model.Activity]{
			Status:  http.StatusCreated,
			Message: "Success creating new activity",
			Data:    activity,
		}
	}

	encoder := json.NewEncoder(writer)
	encodeErr := encoder.Encode(response)
	helper.PanicIfError(encodeErr)
}

func (controller *ActivityControllerImpl) Update(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content-Type", "application/json")

	pathId := request.PathValue("id")
	id, err := strconv.Atoi(pathId)
	helper.PanicIfError(err)

	var reqBody model.Activity
	decoder := json.NewDecoder(request.Body)
	decodeErr := decoder.Decode(&reqBody)
	helper.PanicIfError(decodeErr)

	reqBody.ID = id

	database := controller.Database.WithContext(request.Context())
	helper.PanicIfError(database.Error)

	var response web.BaseResponse[[]model.Activity]
	activity, resultErr := controller.Repository.Update(database, reqBody)

	if resultErr != nil {
		var statusCode int

		switch {
		case strings.Contains(resultErr.Error(), "404"):
			writer.WriteHeader(404)
			statusCode = http.StatusNotFound
		case strings.Contains(resultErr.Error(), "500"):
			writer.WriteHeader(500)
			statusCode = http.StatusInternalServerError
		default:
			writer.WriteHeader(500)
			statusCode = http.StatusInternalServerError
		}

		writer.WriteHeader(http.StatusInternalServerError)
		response = web.BaseResponse[[]model.Activity]{
			Status:  statusCode,
			Message: resultErr.Error(),
			Data:    nil,
		}
	} else {
		writer.WriteHeader(http.StatusOK)
		response = web.BaseResponse[[]model.Activity]{
			Status:  http.StatusOK,
			Message: "Success updating activity",
			Data:    activity,
		}
	}

	encoder := json.NewEncoder(writer)
	encodeErr := encoder.Encode(response)
	helper.PanicIfError(encodeErr)
}

func (controller *ActivityControllerImpl) Delete(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content-Type", "application/json")

	pathId := request.PathValue("id")
	id, convertIdErr := strconv.Atoi(pathId)
	helper.PanicIfError(convertIdErr)

	database := controller.Database.WithContext(request.Context())
	helper.PanicIfError(database.Error)

	var response web.BaseResponse[interface{}]
	resultErr := controller.Repository.Delete(database, id)

	if resultErr != nil {
		var statusCode int

		switch {
		case strings.Contains(resultErr.Error(), "404"):
			writer.WriteHeader(404)
			statusCode = http.StatusNotFound
		case strings.Contains(resultErr.Error(), "500"):
			writer.WriteHeader(500)
			statusCode = http.StatusInternalServerError
		default:
			writer.WriteHeader(500)
			statusCode = http.StatusInternalServerError
		}

		response = web.BaseResponse[interface{}]{
			Status:  statusCode,
			Message: resultErr.Error(),
		}
	} else {
		writer.WriteHeader(http.StatusOK)
		response = web.BaseResponse[interface{}]{
			Status:  http.StatusOK,
			Message: "Success deleting activity",
		}
	}

	encoder := json.NewEncoder(writer)
	encodeErr := encoder.Encode(response)
	helper.PanicIfError(encodeErr)
}
