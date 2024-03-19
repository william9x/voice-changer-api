package adapter

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Braly-Ltd/voice-changer-api-adapter/properties"
	"github.com/Braly-Ltd/voice-changer-api-core/entities"
	"github.com/golibs-starter/golib/log"
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

func (r *RVCAdapter) Infer(ctx context.Context, cmd entities.InferenceCommand) error {
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(cmd); err != nil {
		return fmt.Errorf("encoding request error: %v", err)
	}
	log.Debugc(ctx, "inference request: %s", buf.String())

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
