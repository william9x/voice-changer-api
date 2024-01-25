package clients

import (
	"net/http"
	"time"
)

// NewHTTPClient ...
func NewHTTPClient() *http.Client {
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = 100
	t.MaxConnsPerHost = 100
	t.MaxIdleConnsPerHost = 100

	return &http.Client{
		Timeout:   60 * time.Second,
		Transport: t,
	}
}
