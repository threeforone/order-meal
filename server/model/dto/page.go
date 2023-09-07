package dto

type PageParam struct {
	Page     int `json:"page" form:"page" default:"1"`
	PageSize int `json:"pageSize" form:"pageSize" default:"10"`
}

type PageResp struct {
	Total int64 `json:"total"`
	Data  []any `json:"data"`
}
