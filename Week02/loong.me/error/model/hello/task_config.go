package hello

// TaskConfig setting
type TaskConfig struct {
	TaskID   int    `json:"task_id"`
	TaskName string `json:"task_name"`
	TaskSpec string `json:"task_spec"`
}
