package entities

import (
	"github.com/Braly-Ltd/voice-changer-api-core/constants"
	"github.com/vmihailenco/msgpack/v5"
)

type VoiceChangePayload struct {
	SrcFileName    string
	TargetFileName string
	Model          string
	Transpose      int

	queue constants.QueueType
}

func NewVoiceChangePayload(srcFileName, targetFileName, model string, transpose int, queue constants.QueueType) *VoiceChangePayload {
	return &VoiceChangePayload{
		SrcFileName:    srcFileName,
		TargetFileName: targetFileName,
		Model:          model,
		Transpose:      transpose,
		queue:          queue,
	}
}

func (p *VoiceChangePayload) Pack() ([]byte, error) {
	return msgpack.Marshal(p)
}

func (p *VoiceChangePayload) Queue() (constants.QueueType, error) {
	return p.queue, nil
}
