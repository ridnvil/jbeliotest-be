package service

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"jubeliotesting/internal/dto"
)

type PublisherService struct {
	Rdb     *redis.Client
	Channel string
}

func NewPublisherService(rdb *redis.Client, channel string) *PublisherService {
	return &PublisherService{
		Rdb:     rdb,
		Channel: channel,
	}
}

func (s *PublisherService) PublishMessage(ctx context.Context, payload dto.PublishDto) error {
	payloadParse, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	return s.Rdb.Publish(ctx, s.Channel, payloadParse).Err()
}
