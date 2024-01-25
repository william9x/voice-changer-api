package bootstrap

import (
	adapter "github.com/Braly-Ltd/voice-changer-api-adapter"
	"github.com/Braly-Ltd/voice-changer-api-adapter/clients"
	adapterProps "github.com/Braly-Ltd/voice-changer-api-adapter/properties"
	"github.com/Braly-Ltd/voice-changer-api-core/ports"
	"github.com/Braly-Ltd/voice-changer-api-worker/handlers"
	"github.com/Braly-Ltd/voice-changer-api-worker/properties"
	"github.com/Braly-Ltd/voice-changer-api-worker/routers"
	"github.com/Braly-Ltd/voice-changer-api-worker/workers"
	"github.com/golibs-starter/golib"
	golibgin "github.com/golibs-starter/golib-gin"
	"go.uber.org/fx"
)

func All() fx.Option {
	return fx.Options(
		golib.AppOpt(),
		golib.PropertiesOpt(),
		golib.LoggingOpt(),
		golib.EventOpt(),
		golib.BuildInfoOpt(Version, CommitHash, BuildTime),
		golib.ActuatorEndpointOpt(),
		golib.HttpRequestLogOpt(),

		// Provide all application properties
		golib.ProvideProps(adapterProps.NewMinIOProperties),
		golib.ProvideProps(adapterProps.NewAsynqProperties),
		golib.ProvideProps(properties.NewFileProperties),

		// Provide clients
		fx.Provide(clients.NewMinIOClient),
		fx.Provide(clients.NewAsynqClient),
		fx.Provide(clients.NewHTTPClient),

		// Provide port's implements
		fx.Provide(fx.Annotate(
			adapter.NewMinIOAdapter, fx.As(new(ports.ObjectStoragePort))),
		),
		fx.Provide(fx.Annotate(
			adapter.NewAsynqAdapter, fx.As(new(ports.TaskQueuePort))),
		),
		fx.Provide(fx.Annotate(
			adapter.NewSoVitsVCAdapter, fx.As(new(ports.InferencePort))),
		),

		// Provide task handlers
		handlers.ProvideHandler(handlers.NewVoiceChangeHandler),

		workers.ProvideQueueWorker(),

		// Provide use cases

		// Provide controllers, these controllers will be used
		// when register router was invoked

		// Provide gin http server auto config,
		// actuator endpoints and application routers
		golibgin.GinHttpServerOpt(),
		fx.Invoke(routers.RegisterGinRouters),

		// Graceful shutdown.
		// OnStop hooks will run in reverse order.
		golibgin.OnStopHttpServerOpt(),

		fx.Invoke(startQueueWorker),
	)
}

func startQueueWorker(queueWorker *workers.QueueWorker) {
	go queueWorker.Start()
}
