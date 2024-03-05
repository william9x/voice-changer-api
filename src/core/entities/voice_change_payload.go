package entities

import "github.com/vmihailenco/msgpack/v5"

type VoiceChangePayload struct {
	Model          string
	Transpose      int
	SrcFileName    string
	SrcFileURL     string
	TargetFileName string
	TargetFileURL  string
	EnqueuedAt     int64
}

func (p *VoiceChangePayload) Packed() ([]byte, error) {
	return msgpack.Marshal(p)
}
