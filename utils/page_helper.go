package utils

type PageInfo struct {
	PageNo   int64       `json:"page_no"`
	PageSize int64       `json:"page_size"`
	Total    int64       `json:"total"`
	Records  interface{} `json:"records"`
}
