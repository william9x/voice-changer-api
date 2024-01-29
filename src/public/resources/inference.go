package resources

// Inference ...
type Inference struct {
	TaskID    string `json:"task_id"`
	SourceURL string `json:"source_url"`
	TargetURL string `json:"target_url"`
}

func NewInferenceResource(tid, sourceURL, targetURL string) *Inference {
	return &Inference{
		TaskID:    tid,
		SourceURL: sourceURL,
		TargetURL: targetURL,
	}
}
