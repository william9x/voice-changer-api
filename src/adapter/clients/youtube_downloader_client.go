package clients

import (
	"github.com/kkdai/youtube/v2"
	"net/http"
)

// NewYouTubeDownloaderClient ...
func NewYouTubeDownloaderClient(httpClient *http.Client) *youtube.Client {
	return &youtube.Client{
		HTTPClient: httpClient,
	}
}
