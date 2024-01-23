package clients

import (
	"github.com/Braly-Ltd/voice-changer-api-adapter/properties"
	"github.com/hibiken/asynq"
)

func NewAsynqClient(props *properties.AsynqProperties) *asynq.Client {
	return asynq.NewClient(asynq.RedisClientOpt{
		Addr:     props.Addr,
		Username: props.Username,
		Password: props.Password,
		DB:       props.DB,
		PoolSize: props.PoolSize,
	})
}
