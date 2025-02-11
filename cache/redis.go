package cache

import (
	"context"
	"fmt"

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

func (c *Cache) AddUserToNotify(ctx context.Context, userID int64, jsonData []byte) error {
	return c.client.HSet(ctx, "users_to_notify", userID, jsonData).Err()

}

func (c *Cache) GetUsersToNotifyFromYES(ctx context.Context, userID string) (string, error) {

	result, err := c.client.HGet(ctx, "users_notificated", userID).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("user not found in redis")
	}
	if err != nil {
		return "", err
	}

	return result, nil
}

func (c *Cache) DeleteUserToNotifyFromYes(ctx context.Context, userID string) error {
	return c.client.HDel(ctx, "users_notificated", fmt.Sprintf("%d", userID)).Err()
}

func (c *Cache) GetUsersToNotify(ctx context.Context) (map[string]string, error) {
	return c.client.HGetAll(ctx, "users_to_notify").Result()
}

func (c *Cache) ClearUsersToNotify(ctx context.Context, userID int64, jsonData string) error {
	err := c.client.HSet(ctx, "users_notificated", userID, jsonData).Err()
	if err != nil {
		return err
	}

	return c.client.HDel(ctx, "users_to_notify", fmt.Sprintf("%d", userID)).Err()
}

func (c *Cache) Close() error {
	return c.client.Close()
}
