package controller

import (
	"encoding/json"
	"math"
	"net/http"
	"strconv"

	"com.homindolentrahar.rutinkann-api/db"
	"com.homindolentrahar.rutinkann-api/helper"
	"com.homindolentrahar.rutinkann-api/model"
	"com.homindolentrahar.rutinkann-api/repository"
)

type LogControllerImpl struct {
	Repository repository.LogRepository
}

func NewLogController(repository repository.LogRepository) *LogControllerImpl {
	return &LogControllerImpl{
		Repository: repository,
	}
}

func (controller *LogControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request) {
	pagination := db.ParsePaginationFromRequest(request)
	logs, count, resultErr := controller.Repository.FindAll(pagination)
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

	log, resultErr := controller.Repository.FindById(id)
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

	result, err := controller.Repository.Create(reqBody)
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

	result, err := controller.Repository.Update(reqBody)
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

	err := controller.Repository.Delete(id)
	response := helper.HandleErrorBaseResponse[any](writer, nil, err, helper.BaseResponseConf{
		SuccessStatusCode: http.StatusOK,
		SuccessMessage:    "Success deleting log",
	})

	encoder := json.NewEncoder(writer)
	encodeErr := encoder.Encode(response)
	helper.PanicIfError(encodeErr)
}
