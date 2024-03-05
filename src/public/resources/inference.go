package resources

import (
	"github.com/Braly-Ltd/voice-changer-api-core/entities"
	"github.com/hibiken/asynq"
	"github.com/vmihailenco/msgpack/v5"
)

// CreateInference ...
type CreateInference struct {
	ID        string `json:"id,omitempty"`
	Model     string `json:"model,omitempty"`
	Type      string `json:"type,omitempty"`
	Status    string `json:"status,omitempty"` // Status of the task. Values: active, pending, scheduled, retry, archived, completed
	MaxRetry  int    `json:"max_retry"`
	Deadline  string `json:"deadline,omitempty"` // Deadline for completing the task
	Retention string `json:"retention"`          // Retention in hours for how long to store the task info

	// @Deprecated
	TaskID string `json:"task_id,omitempty"`
}

// Inference ...
type Inference struct {
	ID            string `json:"id,omitempty"`
	Model         string `json:"model,omitempty"`
	Type          string `json:"type,omitempty"`
	Status        string `json:"status,omitempty"` // Status of the task. Values: active, pending, scheduled, retry, archived, completed
	MaxRetry      int    `json:"max_retry"`
	Deadline      int64  `json:"deadline"`
	Retried       int    `json:"retried"`
	LastErr       string `json:"last_err,omitempty"`
	LastFailedAt  int64  `json:"last_failed_at,omitempty"`
	SrcFileURL    string `json:"src_file_url,omitempty"`
	TargetFileURL string `json:"target_file_url,omitempty"`
	Transpose     int    `json:"transpose"`
	EnqueuedAt    string `json:"enqueued_at,omitempty"`
	CompletedAt   string `json:"completed_at,omitempty"`

	// @Deprecated
	TaskID     string `json:"task_id,omitempty"`
	TaskStatus string `json:"task_status,omitempty"`
	Queue      string `json:"queue,omitempty"`
}

func NewFromTaskInfo(info *asynq.TaskInfo) (*Inference, error) {
	var payload entities.VoiceChangePayload
	if err := msgpack.Unmarshal(info.Payload, &payload); err != nil {
		return nil, err
	}

	var failedAt int64 = 0
	if info.LastFailedAt.UnixMilli() > 0 {
		failedAt = info.LastFailedAt.UnixMilli()
	}
	return &Inference{
		TaskID:       info.ID,
		TaskStatus:   info.State.String(),
		Queue:        info.Queue,
		Type:         info.Type,
		MaxRetry:     info.MaxRetry,
		Retried:      info.Retried,
		LastErr:      info.LastErr,
		LastFailedAt: failedAt,
		Deadline:     info.Deadline.UnixMilli(),

		SrcFileURL:    payload.SrcFileURL,
		TargetFileURL: payload.TargetFileURL,
		Model:         payload.Model,
		Transpose:     payload.Transpose,
	}, nil
}
