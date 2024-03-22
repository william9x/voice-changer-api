package requests

import (
	"github.com/Braly-Ltd/voice-changer-api-core/constants"
	"github.com/Braly-Ltd/voice-changer-api-core/entities"
	"mime/multipart"
)

// CreateInferenceRequest ...
type CreateInferenceRequest struct {
	Type      string `form:"type,omitempty,default=vc:rvc" binding:"tasktype"`
	Model     string `form:"model,omitempty,default=trump" binding:"notblank"`
	Transpose int    `form:"transpose,omitempty,default=0" binding:"min=-12,max=12"`

	RawFile     *multipart.FileHeader `form:"file"`
	SrcURL      string                `form:"source_url,omitempty"`
	SrcProvider string                `form:"source_provider,omitempty"`

	QueueID uint8 `form:"queue_id,omitempty"`

	Queue   constants.QueueType `form:"queue,omitempty"`
	SrcFile entities.File
}
