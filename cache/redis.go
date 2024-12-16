package cache

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type Cache struct {
	client *redis.Client
}

func NewCache(addr string, password string, db int) *Cache {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	return &Cache{client: client}
}

func (c *Cache) SetAPIStatus(ctx context.Context, status bool) error {
	return c.client.Set(ctx, "api_status", status, 0).Err()
}

func (c *Cache) GetAPIStatus(ctx context.Context) (bool, error) {
	val, err := c.client.Get(ctx, "api_status").Bool()
	if err == redis.Nil {
		return false, nil
	}
	return val, err
}

func (c *Cache) AddUserToNotify(ctx context.Context, userID int64) error {
	return c.client.SAdd(ctx, "users_to_notify", userID).Err()
}

func (c *Cache) GetUsersToNotify(ctx context.Context) ([]string, error) {
	return c.client.SMembers(ctx, "users_to_notify").Result()
}

func (c *Cache) ClearUsersToNotify(ctx context.Context) error {
	return c.client.Del(ctx, "users_to_notify").Err()
}

func (c *Cache) Close() error {
	return c.client.Close()
}
