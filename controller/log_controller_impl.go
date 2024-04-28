package controller

import (
	"com.homindolentrahar.rutinkann-api/helper"
	"com.homindolentrahar.rutinkann-api/model"
	"com.homindolentrahar.rutinkann-api/repository"
	"encoding/json"
	"gorm.io/gorm"
	"math"
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
	database := controller.Database.WithContext(request.Context())
	helper.PanicIfError(database.Error)

	pagination := helper.ParsePaginationFromRequest(request)
	logs, count, resultErr := controller.Repository.FindAll(database, pagination)
	response := helper.HandleErrorBasePaginationResponse[[]model.Log](writer, &logs, resultErr, helper.BasePaginationResponseConf{
		SuccessStatusCode: http.StatusOK,
		SuccessMessage:    "Success getting all logs",
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

func (controller *LogControllerImpl) FindById(writer http.ResponseWriter, request *http.Request) {
	pathId := request.PathValue("id")
	id, err := strconv.Atoi(pathId)
	helper.PanicIfError(err)

	db := controller.Database.WithContext(request.Context())
	helper.PanicIfError(db.Error)

	log, resultErr := controller.Repository.FindById(db, id)
	response := helper.HandleErrorBaseResponse[model.Log](writer, log, resultErr, helper.BaseResponseConf{
		SuccessStatusCode: http.StatusOK,
		SuccessMessage:    "Success getting log by ID",
	})

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
	response := helper.HandleErrorBaseResponse[model.Log](writer, result, err, helper.BaseResponseConf{
		SuccessStatusCode: http.StatusCreated,
		SuccessMessage:    "Success creating log",
	})

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
	response := helper.HandleErrorBaseResponse[[]model.Log](writer, &result, err, helper.BaseResponseConf{
		SuccessStatusCode: http.StatusOK,
		SuccessMessage:    "Success updating log",
	})

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
	response := helper.HandleErrorBaseResponse[any](writer, nil, err, helper.BaseResponseConf{
		SuccessStatusCode: http.StatusOK,
		SuccessMessage:    "Success deleting log",
	})

	encoder := json.NewEncoder(writer)
	encodeErr := encoder.Encode(response)
	helper.PanicIfError(encodeErr)
}
