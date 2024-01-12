package node

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type KVStore interface {
	Set(key string, value interface{}) error
}

type RedisKVStore struct {
	cli *redis.Client
}

func NewRedisKVStore(redisHost, redisPasswd string) KVStore {
	cli := redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: redisPasswd,
	})

	return &RedisKVStore{
		cli: cli,
	}
}

func (rkv *RedisKVStore) Set(key string, value interface{}) error {
	ctx := context.Background()
	err := rkv.cli.Set(ctx, key, value, 0).Err()

	return err
}
