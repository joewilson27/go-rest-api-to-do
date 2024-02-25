package task

const _TASK_NOT_FOUND_ = "Task not found"

type TaskAdd struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}
