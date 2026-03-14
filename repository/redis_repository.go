package repository

import (
	"context"
	"fmt"

	"tugas-redis-rahmat/domain" // Sesuaikan dengan nama module di go.mod

	"github.com/redis/go-redis/v9"
)

type redisTrackRepository struct {
	client *redis.Client
}

// NewRedisTrackRepository adalah fungsi pembuat (constructor)
func NewRedisTrackRepository(client *redis.Client) domain.TrackRepository {
	return &redisTrackRepository{
		client: client,
	}
}

// IncrementEvent mengeksekusi perintah INCR ke Redis
func (r *redisTrackRepository) IncrementEvent(ctx context.Context, eventName string) (int64, error) {
	key := fmt.Sprintf("track:event:%s", eventName)
	return r.client.Incr(ctx, key).Result()
}
