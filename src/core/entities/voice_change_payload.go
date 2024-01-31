package entities

import "github.com/vmihailenco/msgpack/v5"

type VoiceChangePayload struct {
	SrcFileName    string
	SrcFileURL     string
	TargetFileName string
	TargetFileURL  string
	Model          string
	Transpose      int
}

func NewVoiceChangePayload(
	srcFileName, srcFileURL, targetFileName, targetFileURL, model string,
	transpose int,
) *VoiceChangePayload {
	return &VoiceChangePayload{
		SrcFileName:    srcFileName,
		SrcFileURL:     srcFileURL,
		TargetFileName: targetFileName,
		TargetFileURL:  targetFileURL,
		Model:          model,
		Transpose:      transpose,
	}
}

func (p *VoiceChangePayload) Packed() ([]byte, error) {
	return msgpack.Marshal(p)
}
