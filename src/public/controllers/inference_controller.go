package controllers

import (
	"fmt"
	"github.com/Braly-Ltd/voice-changer-api-core/entities"
	"github.com/Braly-Ltd/voice-changer-api-core/usecases"
	"github.com/Braly-Ltd/voice-changer-api-core/utils"
	"github.com/Braly-Ltd/voice-changer-api-public/properties"
	"github.com/Braly-Ltd/voice-changer-api-public/resources"
	"github.com/gin-gonic/gin"
	"github.com/golibs-starter/golib/exception"
	"github.com/golibs-starter/golib/log"
	"github.com/golibs-starter/golib/web/response"
	"strconv"
)

type InferenceController struct {
	modelProps              *properties.ModelProperties
	inferenceProps          *properties.InferenceProperties
	changeVoiceUseCase      usecases.ChangeVoiceUseCase
	getInferenceInfoUseCase usecases.GetInferenceInfoUseCase
}

func NewInferenceController(
	modelProps *properties.ModelProperties,
	inferenceProps *properties.InferenceProperties,
	changeVoiceUseCase usecases.ChangeVoiceUseCase,
) *InferenceController {
	return &InferenceController{
		modelProps:         modelProps,
		inferenceProps:     inferenceProps,
		changeVoiceUseCase: changeVoiceUseCase,
	}
}

// GetInfer
//
//	@ID				get-inference
//	@Summary 		Get status of an inference task
//	@Description
//	@Tags			InferenceController
//	@Accept			json
//	@Produce		json
//	@Param			transpose	formData		int				false	"Default: 0"
//	@Success		200		{object}	response.Response{data=resources.Inference}
//	@Success		400		{object}	response.Response
//	@Success		404		{object}	response.Response
//	@Failure		500		{object}	response.Response
//	@Router			/api/v1/infer/:id [get]
func (c *InferenceController) GetInfer(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		response.WriteError(ctx.Writer, exception.New(400, "ID invalid"))
		return
	}

	//task, err := c.getInferenceInfoUseCase.GetInferenceInfo(ctx, id)
	//if err != nil {
	//	log.Errorc(ctx, "%v", err)
	//	response.WriteError(ctx.Writer, exception.New(500, "Internal Server Error"))
	//	return
	//}

	response.Write(ctx.Writer, response.Ok(nil))
}

// CreateInfer
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
//	@Success		201		{object}	response.Response{data=resources.Inference}
//	@Failure		400		{object}	response.Response
//	@Failure		500		{object}	response.Response
//	@Router			/api/v1/infer [post]
func (c *InferenceController) CreateInfer(ctx *gin.Context) {
	model, exist := ctx.GetPostForm("model")
	if !exist {
		response.WriteError(ctx.Writer, exception.New(400, "Missing model"))
		return
	}
	if _, exist := c.modelProps.DataMap[model]; !exist {
		response.WriteError(ctx.Writer, exception.New(400, "Model not supported"))
		return
	}

	tranpose := 0
	tranposeStr, exist := ctx.GetPostForm("transpose")
	if exist {
		t, err := strconv.Atoi(tranposeStr)
		if err != nil {
			response.WriteError(ctx.Writer, exception.New(400, "Invalid transpose"))
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
	if !utils.ContainsString(c.inferenceProps.SupportedFiles, srcFile.Ext) {
		msg := fmt.Sprintf("File %s not supported. Supported file are: %v", srcFile.Ext, c.inferenceProps.SupportedFiles)
		response.WriteError(ctx.Writer, exception.New(400, msg))
		return
	}

	taskID, err := c.changeVoiceUseCase.CreateChangeVoiceTask(ctx, srcFile, model, tranpose)
	if err != nil {
		log.Errorc(ctx, "%v", err)
		response.WriteError(ctx.Writer, exception.New(500, "Internal Server Error"))
		return
	}

	response.Write(ctx.Writer, response.Created(resources.NewInferenceResource(taskID)))
}
