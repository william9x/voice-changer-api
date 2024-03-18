package entities

type InferenceCommand struct {
	ModelPath string `json:"model_path"`
	IndexPath string `json:"index_path"`
	InputPath string `json:"input_path"`
	OutPath   string `json:"output_path"`
	Transpose int    `json:"transpose"`
}

type SeparateAudioCommand struct {
	InputPath string `json:"input_path"`
}

type SeparateAudioResponse struct {
	VocalPath string `json:"output_vocal_path"`
	InstPath  string `json:"output_instrument_path"`
}
