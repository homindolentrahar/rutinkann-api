package db

import (
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type Pagination struct {
	Sort     string `json:"sort,omitempty;query:limit"`
	Page     int    `json:"page,omitempty;query:page"`
	PageSize int    `json:"page_size,omitempty;query:page_size"`
}

func Paginate(pagination *Pagination) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		var sort string

		if pagination.Page <= 0 {
			pagination.Page = 1
		}

		if pagination.PageSize <= 0 {
			pagination.PageSize = 15
		}

		offset := (pagination.Page - 1) * pagination.PageSize

		if pagination.Sort == "" {
			pagination.Sort = "-created_at"
		}

		if pagination.Sort[0] == '-' {
			sort = pagination.Sort[1:] + " desc"
		} else {
			sort = pagination.Sort
		}

		return db.Offset(offset).Limit(pagination.PageSize).Order(sort)
	}
}

func ParsePaginationFromRequest(request *http.Request) *Pagination {
	query := request.URL.Query()
	page, _ := strconv.Atoi(query.Get("page"))
	pageSize, _ := strconv.Atoi(query.Get("page_size"))

	return &Pagination{
		Page:     page,
		PageSize: pageSize,
		Sort:     query.Get("sort"),
	}
}
