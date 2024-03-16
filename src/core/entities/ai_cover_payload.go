package entities

import "github.com/vmihailenco/msgpack/v5"

type AICoverPayload struct {
	SrcURL         string
	SrcProvider    string
	ModelID        string
	TargetFileName string
	TargetFileURL  string
	EnqueuedAt     int64
}

func (p *AICoverPayload) Packed() ([]byte, error) {
	return msgpack.Marshal(p)
}
