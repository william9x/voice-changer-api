package workers

import (
	adapterProps "github.com/Braly-Ltd/voice-changer-api-adapter/properties"
	"github.com/Braly-Ltd/voice-changer-api-worker/handlers"
	"github.com/Braly-Ltd/voice-changer-api-worker/properties"
	"github.com/golibs-starter/golib/log"
	"github.com/hibiken/asynq"
)

type AsynqWorker struct {
	server *asynq.Server
	mux    *asynq.ServeMux
}

func NewAsynqWorker(
	handlers []handlers.TaskHandler,
	workerProps *properties.WorkerProperties,
	queueProps *adapterProps.AsynqProperties,
) *AsynqWorker {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr:     queueProps.Addr,
			Password: queueProps.Password,
			DB:       queueProps.DB,
		},
		asynq.Config{
			Concurrency:    workerProps.Concurrency,
			Queues:         queueProps.Queues,
			StrictPriority: workerProps.StrictPriority,
		},
	)

	mux := asynq.NewServeMux()
	for _, handler := range handlers {
		mux.HandleFunc(string(handler.Type()), handler.Handle)
	}

	return &AsynqWorker{
		server: srv,
		mux:    mux,
	}
}

func (w *AsynqWorker) Start() {
	log.Infof("starting asynq worker")
	if err := w.server.Run(w.mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}

func (w *AsynqWorker) Stop() {
	w.server.Stop()
	w.server.Shutdown()
}
