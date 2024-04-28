package helper

import (
	"com.homindolentrahar.rutinkann-api/web"
	"errors"
	"gorm.io/gorm"
	"net/http"
)

type BaseResponseConf struct {
	SuccessStatusCode int    `default:"200"`
	SuccessMessage    string `default:"success executing request"`
	ErrorMessage      string
}

type BasePaginationResponseConf struct {
	SuccessStatusCode int    `default:"200"`
	SuccessMessage    string `default:"success executing request"`
	ErrorMessage      string
	Sort              string `default:""`
	Page              int    `default:"1"`
	PageSize          int    `default:"10"`
	Total             int64
	TotalPage         int
}

func HandleErrorBasePaginationResponse[T any](writer http.ResponseWriter, data *T, err error, conf BasePaginationResponseConf) web.BasePaginationResponse[*T] {
	if err != nil {
		var statusCode int
		var message string

		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			statusCode = http.StatusNotFound
		default:
			statusCode = http.StatusInternalServerError
		}

		if conf.ErrorMessage == "" {
			message = err.Error()
		} else {
			message = conf.ErrorMessage
		}

		writer.WriteHeader(statusCode)
		return web.BasePaginationResponse[*T]{
			Status:    statusCode,
			Message:   message,
			Data:      nil,
			Sort:      conf.Sort,
			Page:      conf.Page,
			PageSize:  conf.PageSize,
			Total:     conf.Total,
			TotalPage: conf.TotalPage,
		}
	} else {
		writer.WriteHeader(conf.SuccessStatusCode)
		return web.BasePaginationResponse[*T]{
			Status:    conf.SuccessStatusCode,
			Message:   conf.SuccessMessage,
			Data:      data,
			Sort:      conf.Sort,
			Page:      conf.Page,
			PageSize:  conf.PageSize,
			Total:     conf.Total,
			TotalPage: conf.TotalPage,
		}
	}
}

func HandleErrorBaseResponse[T any](writer http.ResponseWriter, data *T, err error, conf BaseResponseConf) web.BaseResponse[*T] {
	if err != nil {
		var statusCode int
		var message string

		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			statusCode = http.StatusNotFound
		default:
			statusCode = http.StatusInternalServerError
		}

		if conf.ErrorMessage == "" {
			message = err.Error()
		} else {
			message = conf.ErrorMessage
		}

		writer.WriteHeader(statusCode)
		return web.BaseResponse[*T]{
			Status:  statusCode,
			Message: message,
			Data:    nil,
		}
	} else {
		writer.WriteHeader(conf.SuccessStatusCode)
		return web.BaseResponse[*T]{
			Status:  conf.SuccessStatusCode,
			Message: conf.SuccessMessage,
			Data:    data,
		}
	}
}
