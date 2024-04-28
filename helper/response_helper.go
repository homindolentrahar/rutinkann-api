package helper

import (
	"com.homindolentrahar.rutinkann-api/web"
	"errors"
	"gorm.io/gorm"
	"net/http"
)

func HandleErrorBaseResponse[T interface{}](writer http.ResponseWriter, successStatusCode int, data *T, err error, successMsg string) web.BaseResponse[*T] {
	if err != nil {
		var statusCode int

		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			statusCode = http.StatusNotFound
		default:
			statusCode = http.StatusInternalServerError
		}

		writer.WriteHeader(statusCode)
		return web.BaseResponse[*T]{
			Status:  statusCode,
			Message: err.Error(),
			Data:    nil,
		}
	} else {
		writer.WriteHeader(successStatusCode)
		return web.BaseResponse[*T]{
			Status:  successStatusCode,
			Message: successMsg,
			Data:    data,
		}
	}
}
