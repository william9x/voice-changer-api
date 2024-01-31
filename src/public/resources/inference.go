package resources

import (
	"github.com/Braly-Ltd/voice-changer-api-core/entities"
	"github.com/hibiken/asynq"
	"github.com/vmihailenco/msgpack/v5"
)

// Inference ...
type Inference struct {
	TaskID       string `json:"task_id,omitempty"`
	TaskStatus   string `json:"task_status,omitempty"`
	Queue        string `json:"queue,omitempty"`
	Type         string `json:"type,omitempty"`
	MaxRetry     int    `json:"max_retry,omitempty"`
	Retried      int    `json:"retried,omitempty"`
	LastErr      string `json:"last_err,omitempty"`
	LastFailedAt int64  `json:"last_failed_at,omitempty"`
	Timeout      int64  `json:"timeout,omitempty"`
	Deadline     int64  `json:"deadline,omitempty"`

	SrcFileURL    string `json:"src_file_url,omitempty"`
	TargetFileURL string `json:"target_file_url,omitempty"`
	Model         string `json:"model,omitempty"`
	Transpose     int    `json:"transpose,omitempty"`
}

func NewFromTaskID(taskID string) *Inference {
	return &Inference{
		TaskID: taskID,
	}
}

func NewFromTaskInfo(info *asynq.TaskInfo) (*Inference, error) {
	var payload entities.VoiceChangePayload
	if err := msgpack.Unmarshal(info.Payload, &payload); err != nil {
		return nil, err
	}

	return &Inference{
		TaskID:       info.ID,
		TaskStatus:   info.State.String(),
		Queue:        info.Queue,
		Type:         info.Type,
		MaxRetry:     info.MaxRetry,
		Retried:      info.Retried,
		LastErr:      info.LastErr,
		LastFailedAt: info.LastFailedAt.UnixMilli(),
		Timeout:      info.Timeout.Milliseconds(),
		Deadline:     info.Deadline.UnixMilli(),

		SrcFileURL:    payload.SrcFileURL,
		TargetFileURL: payload.TargetFileURL,
		Model:         payload.Model,
		Transpose:     payload.Transpose,
	}, nil
}
