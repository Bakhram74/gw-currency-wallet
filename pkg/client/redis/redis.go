package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/Bakhram74/gw-currency-wallet/internal/config"
	"github.com/Bakhram74/gw-currency-wallet/internal/service/entity"
	"github.com/redis/go-redis/v9"
)

var (
	ErrCacheNotFound = errors.New("redis: cache not found")
)
var redisClient *Redis

type Redis struct {
	client *redis.Client
	config config.Redis
}

func InitRedis(config config.Redis) error {
	fmt.Println(config.Host, " lalal", config.Port)
	addr := fmt.Sprintf("%s:%s", config.Host, config.Port)
	client := redis.NewClient(&redis.Options{
		Addr: addr,
		DB:   0,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return err
	}
	redisClient = &Redis{
		client: client,
		config: config,
	}

	return nil
}

func Close() {
	redisClient.client.Close()
}

func SetRate(ctx context.Context, id string, param entity.Cache) error {

	value, err := json.Marshal(param)
	if err != nil {
		return err
	}

	err = redisClient.client.Set(ctx, id, value, redisClient.config.ExpiredAt).Err()
	if err != nil {
		return err
	}
	return nil
}

func GetRate(ctx context.Context, id, fromCurrency, toCurrency string) (float32, error) {
	var cache entity.Cache

	value := redisClient.client.Get(ctx, id)

	err := json.Unmarshal([]byte(value.Val()), &cache)
	if err != nil {
		return 0, ErrCacheNotFound
	}

	if fromCurrency != cache.FromCurrency || toCurrency != cache.ToCurrency {
		return 0, ErrCacheNotFound
	}

	return cache.Rate, nil
}
