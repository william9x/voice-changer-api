package services

import (
	"context"
	"fmt"
	"github.com/Braly-Ltd/voice-changer-api-core/constants"
	"github.com/Braly-Ltd/voice-changer-api-core/entities"
	"github.com/Braly-Ltd/voice-changer-api-core/ports"
	"github.com/kkdai/youtube/v2"
)

type DownloadService struct {
	objectStoragePort ports.ObjectStoragePort
	youtubeClient     *youtube.Client
}

func NewDownloadService(
	objectStoragePort ports.ObjectStoragePort,
	youtubeClient *youtube.Client,
) *DownloadService {
	return &DownloadService{
		objectStoragePort: objectStoragePort,
		youtubeClient:     youtubeClient,
	}
}

// Download currently only support YouTube
func (r *DownloadService) Download(ctx context.Context, srcProvider, srcURL string) (entities.File, error) {
	if srcProvider != string(constants.SourceProviderYouTube) {
		return entities.File{}, fmt.Errorf("unsupported source provider: %s", srcProvider)
	}

	video, err := r.youtubeClient.GetVideoContext(ctx, srcURL)
	if err != nil {
		return entities.File{}, fmt.Errorf("failed to get video: %w", err)
	}

	formats := video.Formats.WithAudioChannels().Type("mp4") // only get videos with audio
	if len(formats) == 0 {
		return entities.File{}, fmt.Errorf("no video with audio found")
	}
	formats.Sort()

	content, size, err := r.youtubeClient.GetStreamContext(ctx, video, &formats[0])
	if err != nil {
		return entities.File{}, fmt.Errorf("failed to get stream: %w", err)
	}

	return entities.File{
		Name:     video.Title,
		Ext:      ".mp4",
		Size:     size,
		Content:  content,
		MetaData: nil,
	}, nil
}
