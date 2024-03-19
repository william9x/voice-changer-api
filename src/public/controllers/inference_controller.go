package controllers

import (
	"fmt"
	"github.com/Braly-Ltd/voice-changer-api-core/constants"
	"github.com/Braly-Ltd/voice-changer-api-core/entities"
	"github.com/Braly-Ltd/voice-changer-api-core/utils"
	"github.com/Braly-Ltd/voice-changer-api-public/properties"
	"github.com/Braly-Ltd/voice-changer-api-public/requests"
	"github.com/Braly-Ltd/voice-changer-api-public/resources"
	"github.com/Braly-Ltd/voice-changer-api-public/services"
	"github.com/gin-gonic/gin"
	"github.com/golibs-starter/golib/exception"
	"github.com/golibs-starter/golib/log"
	"github.com/golibs-starter/golib/web/response"
)

type InferenceController struct {
	modelProps       *properties.ModelProperties
	inferenceProps   *properties.InferenceProperties
	inferenceService *services.InferenceService
	downloadService  *services.DownloadService
}

func NewInferenceController(
	modelProps *properties.ModelProperties,
	inferenceProps *properties.InferenceProperties,
	inferenceService *services.InferenceService,
	downloadService *services.DownloadService,
) *InferenceController {
	return &InferenceController{
		modelProps:       modelProps,
		inferenceProps:   inferenceProps,
		inferenceService: inferenceService,
		downloadService:  downloadService,
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
//	@Param			id		path    	string     true        "Task ID"
//	@Success		200		{object}	response.Response{data=resources.Inference}
//	@Success		400		{object}	response.Response
//	@Success		404		{object}	response.Response
//	@Failure		500		{object}	response.Response
//	@Router			/api/v1/infer/{id} [get]
func (c *InferenceController) GetInfer(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		response.WriteError(ctx.Writer, exception.New(400, "ID invalid"))
		return
	}

	queueId, inferId := utils.ExtractInferenceKey(ctx.Param("id"))
	if queueId == "" || inferId == "" {
		response.WriteError(ctx.Writer, exception.New(400, "Invalid infer ID"))
		return
	}

	inferInfo, err := c.inferenceService.GetInferenceInfo(ctx, queueId, inferId)
	if err != nil {
		response.WriteError(ctx.Writer, exception.New(404, "Task not found"))
		return
	}

	resp, err := resources.NewFromTaskInfo(inferInfo)
	if err != nil {
		log.Errorc(ctx, "new task info resource error: %v", err)
		response.WriteError(ctx.Writer, exception.New(500, "Internal Server Error"))
		return
	}

	response.Write(ctx.Writer, response.Ok(resp))
}

// CreateInfer
//
//	@ID				create-inference
//	@Summary 		Change voice of an audio file to target voice
//	@Description
//	@Tags			InferenceController
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			file				formData		file			false	"Source voice"
//	@Param			source_provider		formData		string			false	"Source provider" default(youtube) Enums(youtube)
//	@Param			source_url			formData		string			false	"Source URL"
//	@Param			model				formData		string			true	"Target voice" default(trump)
//	@Param			type				formData		string			true	"Task's type" default(vc:rvc) Enums(vc:rvc, aic)
//	@Param			transpose			formData		int				false	"Transpose" default(0) minimum(-12) maximum(12)
//	@Success		201		{object}	response.Response{data=resources.Infer}
//	@Failure		400		{object}	response.Response
//	@Failure		500		{object}	response.Response
//	@Router			/api/v1/infer [post]
func (c *InferenceController) CreateInfer(ctx *gin.Context) {
	var req requests.CreateInferenceRequest
	if err := ctx.ShouldBind(&req); err != nil {
		response.WriteError(ctx.Writer, exception.New(40000, err.Error()))
		return
	}

	if _, exist := c.modelProps.DataMap[req.Model]; !exist {
		response.WriteError(ctx.Writer, exception.New(40000, "Model not supported"))
		return
	}

	if req.Type == string(constants.TaskTypeVoiceChangeRVC) {
		if req.RawFile == nil {
			response.WriteError(ctx.Writer, exception.New(40000, "file required"))
			return
		}

		file, err := entities.NewFile(req.RawFile)
		if err != nil {
			response.WriteError(ctx.Writer, exception.New(40000, "RawFile invalid"))
			return
		}
		if !utils.ContainsString(c.inferenceProps.SupportedFiles, file.Ext) {
			msg := fmt.Sprintf("RawFile %s not supported. Supported rawFile are: %v", req.SrcFile.Ext, c.inferenceProps.SupportedFiles)
			response.WriteError(ctx.Writer, exception.New(40000, msg))
			return
		}

		req.SrcFile = file
	}

	if req.Type == string(constants.TaskTypeAICover) {
		if req.SrcURL == "" {
			response.WriteError(ctx.Writer, exception.New(40000, "source URL required"))
			return
		}

		file, err := c.downloadService.Download(ctx, req.SrcProvider, req.SrcURL)
		if err != nil {
			response.WriteError(ctx.Writer, exception.New(40000, "Download file error"))
			return
		}

		req.SrcFile = file
	}

	resp, err := c.inferenceService.CreateInference(ctx, req)
	if err != nil {
		log.Errorc(ctx, "%v", err)
		response.WriteError(ctx.Writer, exception.New(500, "Internal Server Error"))
		return
	}

	response.Write(ctx.Writer, response.Created(resp))
}
