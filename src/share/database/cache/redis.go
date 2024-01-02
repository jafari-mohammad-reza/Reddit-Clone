package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/reddit-clone/src/share/config"
	"github.com/redis/go-redis/v9"
	"time"
)

var redisClient *redis.Client

func InitRedis(cfg *config.Config) (context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*cfg.Redis.IdleCheckFrequency)
	fmt.Println(fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port))
	redisClient = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port),
		DB:           cfg.Redis.Db,
		DialTimeout:  cfg.Redis.DialTimeout * time.Second,
		ReadTimeout:  cfg.Redis.ReadTimeout * time.Second,
		WriteTimeout: cfg.Redis.WriteTimeout * time.Second,
		PoolSize:     cfg.Redis.PoolSize,
		PoolTimeout:  cfg.Redis.PoolTimeout,
	})

	_, err := redisClient.Ping(ctx).Result()

	if err != nil {

		return cancel, err
	}
	return cancel, nil
}

func CloseRedisClient() error {
	return redisClient.Close()
}
func GetRedisClient() *redis.Client {
	return redisClient
}

func Set[T any](ctx context.Context, c *redis.Client, key string, value T, duration time.Duration) error {
	v, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return c.Set(ctx, key, v, duration).Err()
}

func Get[T any](ctx context.Context, c *redis.Client, key string) (T, error) {
	var dest T = *new(T)
	v, err := c.Get(ctx, key).Result()
	if err != nil {
		return dest, err
	}
	err = json.Unmarshal([]byte(v), &dest)
	if err != nil {
		return dest, err
	}
	return dest, nil
}
