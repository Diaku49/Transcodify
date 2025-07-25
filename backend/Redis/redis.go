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
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	redisPass := os.Getenv("REDIS_PASS")
	if redisPass == "" {
		redisPass = ""
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisHost + ":" + redisPort,
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
