package model

// TaskConfig setting
type Config struct {
	TaskID   int    `json:"task_id"`
	TaskName string `json:"task_name"`
	TaskSpec string `json:"task_spec"`
}
