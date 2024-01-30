package bootstrap

import (
	"context"
	"errors"
	adapter "github.com/Braly-Ltd/voice-changer-api-adapter"
	"github.com/Braly-Ltd/voice-changer-api-adapter/clients"
	adapterProps "github.com/Braly-Ltd/voice-changer-api-adapter/properties"
	"github.com/Braly-Ltd/voice-changer-api-core/ports"
	"github.com/Braly-Ltd/voice-changer-api-core/usecases"
	"github.com/Braly-Ltd/voice-changer-api-public/controllers"
	"github.com/Braly-Ltd/voice-changer-api-public/properties"
	"github.com/Braly-Ltd/voice-changer-api-public/routers"
	"github.com/golibs-starter/golib"
	golibgin "github.com/golibs-starter/golib-gin"
	"github.com/golibs-starter/golib/log"
	"go.uber.org/fx"
	"net/http"
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
		golib.ProvideProps(properties.NewSwaggerProperties),
		golib.ProvideProps(properties.NewTLSProperties),
		golib.ProvideProps(adapterProps.NewMinIOProperties),
		golib.ProvideProps(adapterProps.NewAsynqProperties),

		// Provide clients
		fx.Provide(clients.NewMinIOClient),
		fx.Provide(clients.NewAsynqClient),

		// Provide port's implements
		fx.Provide(fx.Annotate(
			adapter.NewMinIOAdapter, fx.As(new(ports.ObjectStoragePort))),
		),
		fx.Provide(fx.Annotate(
			adapter.NewAsynqAdapter, fx.As(new(ports.TaskQueuePort))),
		),

		// Provide use cases
		fx.Provide(fx.Annotate(
			usecases.NewChangeVoiceUseCaseImpl, fx.As(new(usecases.ChangeVoiceUseCase))),
		),

		// Provide controllers, these controllers will be used
		// when register router was invoked
		fx.Provide(controllers.NewInferenceController),

		// Provide gin http server auto config,
		// actuator endpoints and application routers
		GinHttpServerOpt(),
		fx.Invoke(routers.RegisterGinRouters),

		// Graceful shutdown.
		// OnStop hooks will run in reverse order.
		golibgin.OnStopHttpServerOpt(),
	)
}

func GinHttpServerOpt() fx.Option {
	return fx.Options(
		fx.Provide(golibgin.NewGinEngine),
		fx.Provide(golibgin.NewHTTPServer),
		fx.Invoke(golibgin.RegisterHandlers),
		fx.Invoke(OnStartHttpsServerHook),
	)
}

func OnStartHttpsServerHook(lc fx.Lifecycle, app *golib.App, httpServer *http.Server, tls *properties.TLSProperties) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Infof("Application will be served at %s. Service name: %s, service path: %s",
				httpServer.Addr, app.Name(), app.Path())
			go func() {
				if tls.Enabled {
					if err := httpServer.ListenAndServeTLS(tls.CertFile, tls.KeyFile); err != nil && !errors.Is(err, http.ErrServerClosed) {
						log.Errorf("Could not serve HTTP request at %s, error [%v]", httpServer.Addr, err)
					}
				} else {
					if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
						log.Errorf("Could not serve HTTP request at %s, error [%v]", httpServer.Addr, err)
					}
				}
				log.Infof("Stopped HTTP Server %s", httpServer.Addr)
			}()
			return nil
		},
	})
}
