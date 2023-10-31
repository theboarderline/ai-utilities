package deepai

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
	"os"
)

type Client struct {
	token string
	http  *resty.Client
}

type ImageRequest struct {
	Prompt string `json:"text"`
}

type ImageResponse struct {
	ID        string `json:"id"`
	OutputURL string `json:"output_url"`
}

func NewClient() *Client {
	apiKey := os.Getenv("DEEPAI_API_KEY")
	return &Client{
		token: apiKey,
		http:  resty.New(),
	}
}

func (c Client) Do(ctx context.Context, text string) (*ImageResponse, error) {

	formData := map[string]string{TextBodyKey: text}

	response, err := c.http.R().
		SetContext(ctx).
		SetHeader(APIKeyHeader, c.token).
		SetFormData(formData).
		Post(URL)
	if err != nil {
		return nil, err
	}

	if response.IsError() {
		err = errors.New(response.String())
		log.Error().Err(err).Msg("could not generate image from deepai")
		return nil, err
	}

	var imageResponse ImageResponse
	if err = json.Unmarshal(response.Body(), &imageResponse); err != nil {
		log.Err(err).Msg("could not unmarshal response from groupme api")
		return nil, err
	}

	return &imageResponse, nil
}
