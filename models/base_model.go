package models

type BaseResponse struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

type BasePagination struct {
	Page int `json:"page" validate:"required"`
	Size int `json:"size" validate:"required"`
}

type BaseRequestGetDataHandler struct {
	Search     string         `json:"search"`
	Pagination BasePagination `json:"pagination"`
}

type BaseRequestGetDataRepository struct {
	Search     map[string]interface{} `json:"search"`
	Pagination BasePagination         `json:"pagination"`
}

type GetMostViewed struct {
	Data interface{} `json:"data,omitempty"`
}
