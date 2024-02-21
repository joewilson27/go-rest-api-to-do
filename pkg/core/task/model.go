package task

type Task struct {
	ID     uint   `gorm:"primaryKey;autoincrement" json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}
