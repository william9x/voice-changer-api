package controllers

import (
	"github.com/Braly-Ltd/voice-changer-api-public/properties"
	"github.com/Braly-Ltd/voice-changer-api-public/resources"
	"github.com/gin-gonic/gin"
	"github.com/golibs-starter/golib/web/response"
)

type ModelController struct {
	modelProps *properties.ModelProperties
}

func NewModelController(
	modelProps *properties.ModelProperties,
) *ModelController {
	return &ModelController{
		modelProps: modelProps,
	}
}

// GetModels
//
//	@ID				get-models
//	@Summary 		Get list supported models
//	@Description
//	@Tags			ModelController
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	response.Response{data=resources.Model}
//	@Failure		500		{object}	response.Response
//	@Router			/api/v1/models [get]
func (c *ModelController) GetModels(ctx *gin.Context) {
	models := make([]*resources.Model, 0, len(c.modelProps.Data))
	for _, model := range c.modelProps.Data {
		models = append(models, resources.NewModelResource(model.ID, model.Name, model.Category))
	}

	response.Write(ctx.Writer, response.Ok(models))
}
