package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/yalagtyarzh/unibot/internal/config"
	"github.com/yalagtyarzh/unibot/pkg/client/youtube"
	"github.com/yalagtyarzh/unibot/pkg/logging"
	"log"
)

type service struct {
	client youtube.Client
	logger *logging.Logger
}

func NewYoutubeService(client youtube.Client, logger *logging.Logger) YoutubeService {
	return &service{client: client, logger: logger}
}

type YoutubeService interface {
	FindTrackByName(ctx context.Context, trackName string) (string, error)
}

func (s *service) FindTrackByName(ctx context.Context, trackName string) (string, error) {
	response, err := s.client.SearchTrack(ctx, trackName)
	if err != nil {
		return "", err
	}

	var responseData map[string]interface{}
	if err = json.NewDecoder(response.Body).Decode(&responseData); err != nil {
		return "", err
	}

	log.Print(responseData)

	a := responseData["items"].([]interface{})
	b := a[0].(map[string]interface{})["id"].(map[string]interface{})["videoId"].(string)

	return fmt.Sprintf(config.YoutubeEndpoint, b), nil
}
