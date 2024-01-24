package controllers

import (
	"github.com/Braly-Ltd/voice-changer-api-core/entities"
	"github.com/Braly-Ltd/voice-changer-api-core/usecases"
	"github.com/gin-gonic/gin"
	"github.com/golibs-starter/golib/exception"
	"github.com/golibs-starter/golib/log"
	"github.com/golibs-starter/golib/web/response"
	"strconv"
)

type InferenceController struct {
	changeVoiceUseCase usecases.ChangeVoiceUseCase
}

func NewInferenceController(
	changeVoiceUseCase usecases.ChangeVoiceUseCase,
) *InferenceController {
	return &InferenceController{
		changeVoiceUseCase: changeVoiceUseCase,
	}
}

// Infer
//
//	@ID				create-inference
//	@Summary 		Change voice of an audio file to target voice
//	@Description
//	@Tags			InferenceController
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			file		formData		file			true	"Source voice"
//	@Param			model		formData		string			true	"Target voice"
//	@Param			transpose	formData		int				false	"Default: 0"
//	@Success		200		{object}	response.Response{data=resources.Inference}
//	@Failure		500		{object}	response.Response
//	@Router			/api/v1/infer [post]
func (c *InferenceController) Infer(ctx *gin.Context) {
	model, exist := ctx.GetPostForm("model")
	if !exist {
		response.WriteError(ctx.Writer, exception.New(400, "Missing model"))
		return
	}

	tranpose := 0
	tranposeStr, exist := ctx.GetPostForm("transpose")
	if exist {
		t, err := strconv.Atoi(tranposeStr)
		if err != nil {
			response.WriteError(ctx.Writer, exception.New(400, "Invalid file"))
			return
		}
		tranpose = t
	}

	file, fileHeader, err := ctx.Request.FormFile("file")
	if err != nil {
		response.WriteError(ctx.Writer, exception.New(400, "Invalid file"))
		return
	}
	defer file.Close()

	srcFile := entities.NewFile(fileHeader.Filename, fileHeader.Size, file)
	if err := c.changeVoiceUseCase.ChangeVoice(ctx, srcFile, model, tranpose); err != nil {
		log.Error(err)
		response.WriteError(ctx.Writer, exception.New(500, "Internal Server Error"))
		return
	}

	response.Write(ctx.Writer, response.Ok(nil))
}
