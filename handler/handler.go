package handler

type WebResponse struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type WebResponseWithPagination struct {
	Code    int         `json:"code"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Page    *PageInfo   `json:"page,omitempty"`
}

type PageInfo struct {
	CurrentPage int `json:"currentPage"`
	TotalPage   int `json:"totalPage"`
	TotalRows   int `json:"totalRows"`
	PerPage     int `json:"perPage"`
}
