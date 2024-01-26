package adapter

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Braly-Ltd/voice-changer-api-adapter/properties"
	"net/http"
)

// SoVitsVCAdapter ...
type SoVitsVCAdapter struct {
	client *http.Client
	props  *properties.SoVitsVcProperties
}

type CreateInferenceRequest struct {
	InputPath  string `json:"input_path,omitempty"`
	OutputPath string `json:"output_path,omitempty"`
	ModelPath  string `json:"model_path,omitempty"`
	ConfigPath string `json:"config_path,omitempty"`
	Transpose  int    `json:"transpose,omitempty"`
}

// NewSoVitsVCAdapter ...
func NewSoVitsVCAdapter(client *http.Client, props *properties.SoVitsVcProperties) *SoVitsVCAdapter {
	return &SoVitsVCAdapter{client: client, props: props}
}

func (r *SoVitsVCAdapter) CreateInference(
	ctx context.Context,
	inputPath, outputPath,
	modelPath, configPath string,
	transpose int,
) error {
	req := CreateInferenceRequest{
		InputPath:  inputPath,
		OutputPath: outputPath,
		ModelPath:  modelPath,
		ConfigPath: configPath,
		Transpose:  transpose,
	}
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(req); err != nil {
		return fmt.Errorf("encoding request error: %v", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", r.props.InferURL, buf)
	if err != nil {
		return fmt.Errorf("build http request error: %v", err)
	}

	resp, err := r.client.Do(httpReq)
	if err != nil {
		return fmt.Errorf("create inference error: %v", err)
	}

	if resp.StatusCode != 201 && resp.StatusCode != 200 {
		return fmt.Errorf("create inference failed")
	}

	return nil
}
