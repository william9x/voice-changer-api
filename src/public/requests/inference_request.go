package requests

import (
	"github.com/Braly-Ltd/voice-changer-api-core/entities"
	"mime/multipart"
)

// CreateInferenceRequest ...
type CreateInferenceRequest struct {
	Model     string                `form:"model,omitempty" binding:"notblank"`
	RawFile   *multipart.FileHeader `form:"file" binding:"required"`
	Type      string                `form:"type,omitempty,default=vc:rvc" binding:"tasktype"`
	Transpose int                   `form:"transpose,omitempty,default=0" binding:"min=-12,max=12"`

	SrcFile entities.File
}
