package database

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	mantisError "github.com/sphireco/mantis/error"
	"time"
)

type Redis struct {
	client      *redis.Client
	context     context.Context
	Options     *redis.Options
	IsConnected bool
}

// Init
func (r *Redis) Init() error {
	r.context = context.Background()
	r.client = redis.NewClient(r.Options)

	if r.CheckIfConnected() == true {
		mantisError.HandleError("unable to connect to Redis", errors.New("ping failed"))
		return errors.New("unable to connect to Redis")
	}

	return nil
}

// CheckIfConnected Checks our connection status to our Redis DB
func (r *Redis) CheckIfConnected() bool {
	if pong, err := r.client.Ping(r.context).Result(); err != nil && pong == "PONG" {
		r.IsConnected = true
	}
	return r.IsConnected
}

// Get a key value pair from our Redis DB
func (r *Redis) Get(key string) (string, error) {
	value, err := r.client.Get(r.context, key).Result()

	if err == redis.Nil {
		return "", errors.New("key does not exist")
	}

	if err != nil {
		return "", errors.New("err not nil attempting to Get key from Redis")
	}

	return value, nil
}

// Set a key value pair in our Redis DB
func (r *Redis) Set(key string, value string, expiration time.Duration) error {
	return r.client.Set(r.context, key, value, expiration).Err()
}

func (r *Redis) GetRawConnectionAndContext() (*redis.Client, context.Context) {
	return r.client, r.context
}
