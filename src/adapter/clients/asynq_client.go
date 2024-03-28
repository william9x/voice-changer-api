package clients

import (
	"github.com/Braly-Ltd/voice-changer-api-adapter/properties"
	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
	"time"
)

func NewAsynqClient(props *properties.AsynqProperties) *asynq.Client {
	return asynq.NewClient(newCustomRedisClientOpt(props))
}

func NewAsynqInspector(props *properties.AsynqProperties) *asynq.Inspector {
	return asynq.NewInspector(newRedisClientOpt(props))
}

func newRedisClientOpt(props *properties.AsynqProperties) asynq.RedisClientOpt {
	return asynq.RedisClientOpt{
		Addr:         props.Addr,
		Username:     props.Username,
		Password:     props.Password,
		DB:           props.DB,
		PoolSize:     props.PoolSize,
		DialTimeout:  time.Minute,
		ReadTimeout:  -1,
		WriteTimeout: time.Minute,
	}
}

func newCustomRedisClientOpt(props *properties.AsynqProperties) CustomRedisClientOpt {
	return CustomRedisClientOpt{
		Addr:             props.Addr,
		Username:         props.Username,
		Password:         props.Password,
		DB:               props.DB,
		PoolSize:         props.PoolSize,
		DialTimeout:      time.Minute,
		ReadTimeout:      -1,
		WriteTimeout:     time.Minute,
		DisableIndentity: true,
	}
}

type CustomRedisClientOpt struct {
	Addr             string
	Username         string
	Password         string
	DB               int
	PoolSize         int
	DialTimeout      time.Duration
	ReadTimeout      time.Duration
	WriteTimeout     time.Duration
	DisableIndentity bool
}

func (opt CustomRedisClientOpt) MakeRedisClient() interface{} {
	return redis.NewClient(&redis.Options{
		Addr:             opt.Addr,
		Username:         opt.Username,
		Password:         opt.Password,
		DB:               opt.DB,
		PoolSize:         opt.PoolSize,
		DialTimeout:      opt.DialTimeout,
		ReadTimeout:      opt.ReadTimeout,
		WriteTimeout:     opt.WriteTimeout,
		DisableIndentity: opt.DisableIndentity,
	})
}
