package helper

import (
	"com.homindolentrahar.rutinkann-api/db"
	"net/http"
	"strconv"
)

func ParsePaginationFromRequest(request *http.Request) *db.Pagination {
	query := request.URL.Query()
	page, _ := strconv.Atoi(query.Get("page"))
	pageSize, _ := strconv.Atoi(query.Get("page_size"))

	return &db.Pagination{
		Page:     page,
		PageSize: pageSize,
		Sort:     query.Get("sort"),
	}
}
