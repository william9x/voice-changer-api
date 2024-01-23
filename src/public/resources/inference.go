package resources

// Inference ...
type Inference struct {
	TaskID    string `json:"task_id"`
	SourceURL string `json:"source_url"`
	TargetURL string `json:"target_url"`
	Status    string `json:"status"`
}
