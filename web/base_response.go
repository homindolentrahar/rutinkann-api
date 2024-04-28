package web

type BaseResponse[T any] struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

type BasePaginationResponse[T any] struct {
	Status    int    `json:"status"`
	Message   string `json:"message"`
	Data      T      `json:"data"`
	Sort      string `json:"sort,omitempty"`
	Page      int    `json:"page,omitempty"`
	PageSize  int    `json:"page_size,omitempty"`
	Total     int64  `json:"total"`
	TotalPage int    `json:"total_page"`
}
