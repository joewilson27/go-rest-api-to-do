package task

const _TASK_NOT_FOUND_ = "Task not found"

type TaskAdd struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

type PaginateResponse struct {
	TotalData   int64       `json:"totalData"`
	TotalPage   int         `json:"totalPage"`
	PageSize    int         `json:"pageSize"`
	CurrentPage int         `json:"currentPage"`
	Data        interface{} `json:"data"`
}

type DataResponse struct {
	TotalPage   int         `json:"totalPage"`
	PageSize    int         `json:"pageSize"`
	CurrentPage int         `json:"currentPage"`
	Data        interface{} `json:"data"`
}
