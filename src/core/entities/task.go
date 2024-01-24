package entities

import "github.com/Braly-Ltd/voice-changer-api-core/constants"

type Task interface {
	Pack() ([]byte, error)
	Type() constants.TaskType
	Queue() constants.QueueType
}
