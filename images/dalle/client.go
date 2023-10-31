package dalle

import (
	"bytes"
	"context"
	"encoding/base64"
	"github.com/rs/zerolog/log"
	"github.com/sashabaranov/go-openai"
	"io"
)

type Client struct {
	dalleClient *openai.Client
}

func NewClient(clientConfig openai.ClientConfig) *Client {
	client := openai.NewClientWithConfig(clientConfig)

	return &Client{
		dalleClient: client,
	}
}

func (c Client) GenerateImage(ctx context.Context, text string) (io.Reader, error) {

	req := openai.ImageRequest{
		Prompt:         text,
		Size:           openai.CreateImageSize256x256,
		ResponseFormat: openai.CreateImageResponseFormatB64JSON,
		N:              1,
	}

	res, err := c.dalleClient.CreateImage(ctx, req)
	if err != nil {
		log.Err(err).Msg("Image creation error")
		return nil, err
	}

	imgBytes, err := base64.StdEncoding.DecodeString(res.Data[0].B64JSON)
	if err != nil {
		log.Err(err).Msg("Base64 decode error")
		return nil, err
	}

	return bytes.NewReader(imgBytes), nil
}
