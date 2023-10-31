package generator

import (
	"context"
	"io"
)

type ImageGenerator interface {
	GenerateImage(ctx context.Context, text string) (io.Reader, error)
}
