package controller

import (
	"encoding/json"
	"math"
	"net/http"
	"strconv"

	"com.homindolentrahar.rutinkann-api/model"
	"com.homindolentrahar.rutinkann-api/repository"

	"com.homindolentrahar.rutinkann-api/helper"
)

type RoutineControllerImpl struct {
	Repository repository.RoutineRepository
}

func NewRoutineController(repository repository.RoutineRepository) *RoutineControllerImpl {
	return &RoutineControllerImpl{
		Repository: repository,
	}
}

func (controller *RoutineControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request) {
	pagination := helper.ParsePaginationFromRequest(request)
	activities, count, resultErr := controller.Repository.FindAll(pagination)
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

func (controller *RoutineControllerImpl) FindById(writer http.ResponseWriter, request *http.Request) {
	pathId := request.PathValue("id")
	id, err := strconv.Atoi(pathId)
	helper.PanicIfError(err)

	activity, resultErr := controller.Repository.FindById(id)
	response := helper.HandleErrorBaseResponse(writer, activity, resultErr, helper.BaseResponseConf{
		SuccessStatusCode: http.StatusOK,
		SuccessMessage:    "Success getting activity by ID",
	})

	encoder := json.NewEncoder(writer)
	encodeErr := encoder.Encode(response)
	helper.PanicIfError(encodeErr)
}

func (controller *RoutineControllerImpl) Create(writer http.ResponseWriter, request *http.Request) {
	var reqBody model.Routine
	decoder := json.NewDecoder(request.Body)
	decodeErr := decoder.Decode(&reqBody)
	helper.PanicIfError(decodeErr)

	activity, resultError := controller.Repository.Create(reqBody)
	response := helper.HandleErrorBaseResponse(writer, &activity, resultError, helper.BaseResponseConf{
		SuccessStatusCode: http.StatusCreated,
		SuccessMessage:    "Success creating activity",
	})

	encoder := json.NewEncoder(writer)
	encodeErr := encoder.Encode(response)
	helper.PanicIfError(encodeErr)
}

func (controller *RoutineControllerImpl) Update(writer http.ResponseWriter, request *http.Request) {
	pathId := request.PathValue("id")
	id, err := strconv.Atoi(pathId)
	helper.PanicIfError(err)

	var reqBody model.Routine
	decoder := json.NewDecoder(request.Body)
	decodeErr := decoder.Decode(&reqBody)
	helper.PanicIfError(decodeErr)

	reqBody.ID = id

	activity, resultErr := controller.Repository.Update(reqBody)
	response := helper.HandleErrorBaseResponse(writer, &activity, resultErr, helper.BaseResponseConf{
		SuccessStatusCode: http.StatusOK,
		SuccessMessage:    "Success updating activity",
	})

	encoder := json.NewEncoder(writer)
	encodeErr := encoder.Encode(response)
	helper.PanicIfError(encodeErr)
}

func (controller *RoutineControllerImpl) Delete(writer http.ResponseWriter, request *http.Request) {
	pathId := request.PathValue("id")
	id, convertIdErr := strconv.Atoi(pathId)
	helper.PanicIfError(convertIdErr)

	resultErr := controller.Repository.Delete(id)
	response := helper.HandleErrorBaseResponse[interface{}](writer, nil, resultErr, helper.BaseResponseConf{
		SuccessStatusCode: http.StatusOK,
		SuccessMessage:    "Success deleting activity",
	})

	encoder := json.NewEncoder(writer)
	encodeErr := encoder.Encode(response)
	helper.PanicIfError(encodeErr)
}
