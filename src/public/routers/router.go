package routers

import (
	"github.com/Braly-Ltd/voice-changer-api-public/controllers"
	"github.com/Braly-Ltd/voice-changer-api-public/docs"
	"github.com/Braly-Ltd/voice-changer-api-public/properties"
	"github.com/gin-gonic/gin"
	"github.com/golibs-starter/golib"
	"github.com/golibs-starter/golib/web/actuator"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"go.uber.org/fx"
)

type RegisterRoutersIn struct {
	fx.In
	App          *golib.App
	Engine       *gin.Engine
	SwaggerProps *properties.SwaggerProperties

	Actuator            *actuator.Endpoint
	InferenceController *controllers.InferenceController
	ModelController     *controllers.ModelController
}

func RegisterGinRouters(p RegisterRoutersIn) {
	group := p.Engine.Group(p.App.Path())
	group.GET("/actuator/health", gin.WrapF(p.Actuator.Health))
	group.GET("/actuator/info", gin.WrapF(p.Actuator.Info))

	if p.SwaggerProps.Enabled {
		docs.SwaggerInfo.BasePath = p.App.Path()
		group.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	apiGroup := group.Group("/api")

	// Model APIs
	apiGroup.GET("/v1/models", p.ModelController.GetModels)

	// Inference APIs
	apiGroup.GET("/v1/infer/:id", p.InferenceController.GetInfer)
	apiGroup.POST("/v1/infer", p.InferenceController.CreateInfer)
}
