package deepai

import (
	"bytes"
	"context"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
)

func (c Client) GenerateImage(ctx context.Context, text string) (io.Reader, error) {

	res, err := c.Do(ctx, text)
	if err != nil {
		log.Err(err).Msg("Image creation error")
		return nil, err
	}

	log.Debug().Msgf("DeepAI image response: %s", res.OutputURL)

	imgDataResponse, err := http.Get(res.OutputURL)
	if err != nil {
		return nil, err
	}
	defer imgDataResponse.Body.Close()

	imgData, err := io.ReadAll(imgDataResponse.Body)
	if err != nil {
		log.Err(err).Msg("could not read image data")
		return nil, err
	}

	return bytes.NewReader(imgData), nil
}
