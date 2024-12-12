package redis

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisService struct {
	client *redis.Client
}

func NewRedisService(addr, password string, db int) *RedisService {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	return &RedisService{client: rdb}
}

func (r *RedisService) Set(key string, value string, expiration time.Duration) error {
	ctx := context.Background()
	err := r.client.Set(ctx, key, value, expiration).Err()
	if err != nil {
		log.Printf("Error setting key %s: %v", key, err)
	}
	return err
}

func (r *RedisService) Get(key string) (string, error) {
	ctx := context.Background()
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		log.Printf("Error getting key %s: %v", key, err)
		return "", err
	}
	return val, nil
}

func (r *RedisService) Publish(ctx context.Context, channel, message string) error {
	return r.client.Publish(ctx, channel, message).Err()
}

func (r *RedisService) Subscribe(ctx context.Context, channel string, handleMessage func(message string)) error {
	sub := r.client.Subscribe(ctx, channel)

	for msg := range sub.Channel() {
		handleMessage(msg.Payload)
	}

	return nil
}
