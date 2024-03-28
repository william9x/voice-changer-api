package clients

import (
	"github.com/Braly-Ltd/voice-changer-api-adapter/properties"
	"github.com/hibiken/asynq"
)

func NewAsynqClient(props *properties.AsynqProperties) *asynq.Client {
	return asynq.NewClient(newRedisClientOpt(props))
}

func NewAsynqInspector(props *properties.AsynqProperties) *asynq.Inspector {
	return asynq.NewInspector(newRedisClientOpt(props))
}

func newRedisClientOpt(props *properties.AsynqProperties) asynq.RedisClientOpt {
	return asynq.RedisClientOpt{
		Addr:        props.Addr,
		Username:    props.Username,
		Password:    props.Password,
		DB:          props.DB,
		PoolSize:    props.PoolSize,
		ReadTimeout: -1,
	}
}
