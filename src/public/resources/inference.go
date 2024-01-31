package resources

// Inference ...
type Inference struct {
	TaskID string `json:"task_id"`
}

func NewInferenceResource(taskID string) *Inference {
	return &Inference{
		TaskID: taskID,
	}
}
