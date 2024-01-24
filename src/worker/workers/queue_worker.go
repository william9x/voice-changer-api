package workers

import (
	"github.com/Braly-Ltd/voice-changer-api-core/constants"
	"github.com/Braly-Ltd/voice-changer-api-worker/handlers"
	"github.com/golibs-starter/golib/log"
	"github.com/hibiken/asynq"
	"go.uber.org/fx"
)

type QueueWorker struct {
	handlers []handlers.TaskHandler
}

func NewQueueWorker(
	handlers []handlers.TaskHandler,
) *QueueWorker {
	return &QueueWorker{handlers: handlers}
}

func ProvideQueueWorker() fx.Option {
	return fx.Provide(
		fx.Annotate(
			NewQueueWorker,
			fx.ParamTags(`group:"task_handlers"`),
		),
	)
}

func (w *QueueWorker) Start() {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr:     "localhost:6379",
			Password: "braly@123",
			DB:       1,
		},
		asynq.Config{
			Concurrency: 10,
			Queues: map[string]int{
				"level_2": 6,
				"default": 3,
			},
		},
	)

	// mux maps a type to a handler
	mux := asynq.NewServeMux()
	mux.HandleFunc(string(constants.TaskTypeInfer), w.handlers[0].Handle)

	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
