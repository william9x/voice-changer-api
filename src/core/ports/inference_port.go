package ports

import (
	"context"
)

type InferencePort interface {
	CreateInference(
		ctx context.Context,
		inputPath, outputPath,
		modelPath, configPath string,
		transpose int,
	) error
}
