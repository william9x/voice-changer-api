package adapter

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Braly-Ltd/voice-changer-api-adapter/properties"
	"github.com/Braly-Ltd/voice-changer-api-core/entities"
	"net/http"
)

// RVCAdapter ...
type RVCAdapter struct {
	client *http.Client
	props  *properties.RVCProperties
}

// NewRVCAdapter ...
func NewRVCAdapter(client *http.Client, props *properties.RVCProperties) *RVCAdapter {
	return &RVCAdapter{client: client, props: props}
}

func (r *RVCAdapter) CreateInference(ctx context.Context, cmd entities.InferenceCommand) error {
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(cmd); err != nil {
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

func (r *RVCAdapter) SeperateAudio(ctx context.Context, cmd entities.SeparateAudioCommand) (entities.SeparateAudioResponse, error) {
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(cmd); err != nil {
		return entities.SeparateAudioResponse{}, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", r.props.UVRPathURL, buf)
	if err != nil {
		return entities.SeparateAudioResponse{}, err
	}

	resp, err := r.client.Do(httpReq)
	if err != nil {
		return entities.SeparateAudioResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 201 && resp.StatusCode != 200 {
		return entities.SeparateAudioResponse{}, err
	}

	var respData entities.SeparateAudioResponse
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return entities.SeparateAudioResponse{}, err
	}

	return respData, nil
}
