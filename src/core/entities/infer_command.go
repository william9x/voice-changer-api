package entities

type InferenceCommand struct {
	ModelPath string `json:"model_path"`
	IndexPath string `json:"index_path"`
	InputPath string `json:"input_path"`
	OutPath   string `json:"output_path"`
	Transpose int    `json:"transpose"`
}
