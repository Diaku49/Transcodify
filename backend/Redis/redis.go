package Redis

import (
	"context"
	"os"
	"strings"
	"time"

	"github.com/Diaku49/FoodOrderSystem/backend/internals/model"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type RedisClient struct {
	R *redis.Client
}

func NewRedisClient() *RedisClient {
	redisAdd := os.Getenv("REDIS_ADDR")
	redisPass := os.Getenv("REDIS_PASS")
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAdd,
		Password: redisPass,
		DB:       0,
	})

	rdbc := RedisClient{R: rdb}
	return &rdbc
}

func (rdbc *RedisClient) SaveMetadata(videoID string, metadata *model.UploadedTempMetadata) error {
	key := "video:" + videoID
	matadataMap := map[string]any{
		"id":          metadata.Id,
		"videoName":   metadata.VideoName,
		"path":        metadata.Path,
		"resolutions": strings.Join(metadata.Resolutions, ","),
	}
	err := rdbc.R.HSet(ctx, key, matadataMap).Err()
	return err
}

func (rdbc *RedisClient) GetVideoProgress(videoID string) (string, error) {
	key := "video:" + videoID
	return rdbc.R.HGet(ctx, key, "progress").Result()
}

func (rdbc *RedisClient) UpdateVideoProgress(videoID string, progress string) error {
	key := "video:" + videoID
	return rdbc.R.HSet(ctx, key, "progress", progress).Err()
}

func (rdbc *RedisClient) SetVideoMetadataExpiry(videoID string, ttl time.Duration) error {
	key := "video:" + videoID
	return rdbc.R.Expire(ctx, key, ttl).Err()
}
