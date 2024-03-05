package resources

import (
	"github.com/Braly-Ltd/voice-changer-api-core/entities"
	"github.com/Braly-Ltd/voice-changer-api-core/utils"
	"github.com/hibiken/asynq"
	"github.com/vmihailenco/msgpack/v5"
	"time"
)

// CreateInference ...
type CreateInference struct {
	ID       string `json:"id,omitempty"`
	Model    string `json:"model,omitempty"`
	Type     string `json:"type,omitempty"`
	Status   string `json:"status,omitempty"` // Status of the task. Values: active, pending, scheduled, retry, archived, completed
	MaxRetry int    `json:"max_retry"`
	Deadline string `json:"deadline,omitempty"` // Deadline for completing the task

	// @Deprecated
	TaskID string `json:"task_id,omitempty"`
}

// Inference ...
type Inference struct {
	ID           string `json:"id,omitempty"`
	Model        string `json:"model,omitempty"`
	Type         string `json:"type,omitempty"`
	Status       string `json:"status,omitempty"` // Status of the task. Values: active, pending, scheduled, retry, archived, completed
	MaxRetry     int    `json:"max_retry"`
	Deadline     string `json:"deadline"`
	Retried      int    `json:"retried"`
	LastErr      string `json:"last_err,omitempty"`
	LastFailedAt string `json:"last_failed_at,omitempty"`
	EnqueuedAt   string `json:"enqueued_at,omitempty"`
	CompletedAt  string `json:"completed_at,omitempty"`

	SrcFileURL    string `json:"src_file_url,omitempty"`
	TargetFileURL string `json:"target_file_url,omitempty"`
	Transpose     int    `json:"transpose"`

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

	var failedAt time.Time
	if info.LastFailedAt.UnixMilli() > 0 {
		failedAt = time.UnixMilli(info.LastFailedAt.UnixMilli())
	}

	return &Inference{
		ID:           utils.BuildInferenceKey(info.Queue, info.ID),
		Model:        payload.Model,
		Type:         info.Type,
		Status:       info.State.String(),
		MaxRetry:     info.MaxRetry,
		Deadline:     info.Deadline.Format(time.RFC3339),
		Retried:      info.Retried,
		LastErr:      info.LastErr,
		LastFailedAt: failedAt.Format(time.RFC3339),
		EnqueuedAt:   time.UnixMilli(payload.EnqueuedAt).Format(time.RFC3339),
		CompletedAt:  info.CompletedAt.Format(time.RFC3339),

		SrcFileURL:    payload.SrcFileURL,
		TargetFileURL: payload.TargetFileURL,
		Transpose:     payload.Transpose,

		// @Deprecated
		TaskID:     info.ID,
		TaskStatus: info.State.String(),
		Queue:      info.Queue,
	}, nil
}
