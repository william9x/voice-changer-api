package adapter

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/Braly-Ltd/voice-changer-api-adapter/properties"
	"github.com/Braly-Ltd/voice-changer-api-core/entities"
	"net/http"
)

// UVRAdapter ...
type UVRAdapter struct {
	client *http.Client
	props  *properties.UVRProperties
}

// NewUVRAdapter ...
func NewUVRAdapter(client *http.Client, props *properties.UVRProperties) *UVRAdapter {
	return &UVRAdapter{client: client, props: props}
}

func (r *UVRAdapter) Infer(ctx context.Context, cmd entities.SeparateAudioCommand) (entities.SeparateAudioResponse, error) {
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(cmd); err != nil {
		return entities.SeparateAudioResponse{}, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", r.props.InferURL, buf)
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
