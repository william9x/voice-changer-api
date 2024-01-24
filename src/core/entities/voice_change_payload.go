package entities

import (
	"github.com/Braly-Ltd/voice-changer-api-core/constants"
	"github.com/vmihailenco/msgpack/v5"
)

type VoiceChangeTask struct {
	SrcFileName    string
	TargetFileName string
	Model          string
	Transpose      int

	TaskID    string
	TaskType  constants.TaskType
	TaskQueue constants.QueueType
}

func NewVoiceChangeTask(
	tID string,
	srcFileName, targetFileName, model string,
	transpose int,
	tType constants.TaskType,
	tQueue constants.QueueType,
) *VoiceChangeTask {
	return &VoiceChangeTask{
		SrcFileName:    srcFileName,
		TargetFileName: targetFileName,
		Model:          model,
		Transpose:      transpose,

		TaskID:    tID,
		TaskType:  tType,
		TaskQueue: tQueue,
	}
}

func (p *VoiceChangeTask) ID() string {
	return p.TaskID
}

func (p *VoiceChangeTask) Pack() ([]byte, error) {
	return msgpack.Marshal(p)
}

func (p *VoiceChangeTask) Type() constants.TaskType {
	return p.TaskType
}

func (p *VoiceChangeTask) Queue() constants.QueueType {
	return p.TaskQueue
}
