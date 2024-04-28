package controller

import (
	"com.homindolentrahar.rutinkann-api/model"
	"com.homindolentrahar.rutinkann-api/repository"
	"encoding/json"
	"gorm.io/gorm"
	"math"
	"net/http"
	"strconv"

	"com.homindolentrahar.rutinkann-api/helper"
)

type ActivityControllerImpl struct {
	Repository repository.ActivityRepository
	Database   *gorm.DB
}

func NewActivityController(repository repository.ActivityRepository, database *gorm.DB) *ActivityControllerImpl {
	return &ActivityControllerImpl{
		Repository: repository,
		Database:   database,
	}
}

func (controller *ActivityControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request) {
	db := controller.Database.WithContext(request.Context())
	helper.PanicIfError(db.Error)

	pagination := helper.ParsePaginationFromRequest(request)
	activities, count, resultErr := controller.Repository.FindAll(db, pagination)
	response := helper.HandleErrorBasePaginationResponse(writer, &activities, resultErr, helper.BasePaginationResponseConf{
		SuccessStatusCode: http.StatusOK,
		SuccessMessage:    "Success getting all activities",
		Sort:              pagination.Sort,
		Page:              pagination.Page,
		PageSize:          pagination.PageSize,
		Total:             count,
		TotalPage:         int(math.Ceil(float64(count) / float64(pagination.PageSize))),
	})

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

	activity, resultErr := controller.Repository.FindById(db, id)
	response := helper.HandleErrorBaseResponse(writer, activity, resultErr, helper.BaseResponseConf{
		SuccessStatusCode: http.StatusOK,
		SuccessMessage:    "Success getting activity by ID",
	})

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

	activity, resultError := controller.Repository.Create(database, reqBody)
	response := helper.HandleErrorBaseResponse(writer, &activity, resultError, helper.BaseResponseConf{
		SuccessStatusCode: http.StatusCreated,
		SuccessMessage:    "Success creating activity",
	})

	encoder := json.NewEncoder(writer)
	encodeErr := encoder.Encode(response)
	helper.PanicIfError(encodeErr)
}

func (controller *ActivityControllerImpl) Update(writer http.ResponseWriter, request *http.Request) {
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

	activity, resultErr := controller.Repository.Update(database, reqBody)
	response := helper.HandleErrorBaseResponse(writer, &activity, resultErr, helper.BaseResponseConf{
		SuccessStatusCode: http.StatusOK,
		SuccessMessage:    "Success updating activity",
	})

	encoder := json.NewEncoder(writer)
	encodeErr := encoder.Encode(response)
	helper.PanicIfError(encodeErr)
}

func (controller *ActivityControllerImpl) Delete(writer http.ResponseWriter, request *http.Request) {
	pathId := request.PathValue("id")
	id, convertIdErr := strconv.Atoi(pathId)
	helper.PanicIfError(convertIdErr)

	database := controller.Database.WithContext(request.Context())
	helper.PanicIfError(database.Error)

	resultErr := controller.Repository.Delete(database, id)
	response := helper.HandleErrorBaseResponse[interface{}](writer, nil, resultErr, helper.BaseResponseConf{
		SuccessStatusCode: http.StatusOK,
		SuccessMessage:    "Success deleting activity",
	})

	encoder := json.NewEncoder(writer)
	encodeErr := encoder.Encode(response)
	helper.PanicIfError(encodeErr)
}
