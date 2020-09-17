package database

import (
	"errors"
	"github.com/gomodule/redigo/redis"
	mantisError "github.com/sphireco/mantis/error"
)

type Redis struct {
	Client    *redis.Pool
	Conn      redis.Conn
	Addr      string
	Password  string
	DB        int
	Network   string // typically tcp
	MaxIdle   int    // Maximum number of idle connections in the pool
	MaxActive int    // max number of connections
	IsConnected bool // Did we achieve a connection?
}

// Connect
func (r *Redis) Connect() error {
	if len(r.Addr) <= 0 || len(r.Network) <= 0 {
		return errors.New("no valid configuration")
	}

	r.Client = &redis.Pool{
		MaxIdle:   r.MaxIdle,
		MaxActive: r.MaxActive,
		Dial: func() (redis.Conn, error) {
			return redis.Dial(r.Network, r.Addr)
		},
	}

	r.Conn = r.Client.Get()

	defer func() {
		if err := r.Conn.Close(); err != nil {
			mantisError.HandleError("error closing Redis connection", err)
		}
	}()

	if r.CheckIfConnected() == true {
		mantisError.HandleError("unable to connect to Redis", errors.New("ping failed"))
	}

	return nil
}

// CheckIfConnected Checks our connection status to our Redis DB
func (r *Redis) CheckIfConnected() bool {
	r.IsConnected = ping(r.Conn)
	return r.IsConnected
}

// ping returns true if we are connected
func ping(c redis.Conn) bool {
	pong, err := c.Do("PING")
	if err != nil {
		return false
	}

	if s, err := redis.String(pong, err); err != nil || s != "PONG" {
		return false
	}

	return true
}

// Get a key value pair from our Redis DB
func (r *Redis) Get(key string) (string, error) {
	value, err := redis.String(r.Conn.Do("GET", key))

	if err == redis.ErrNil || err != nil {
		return "", errors.New("key does not exist or otherwise not found")
	}

	return value, nil
}

// Set a key value pair in our Redis DB
func (r *Redis) Set(key string, value string) error {
	if _, err := r.Conn.Do("SET", key, value); err != nil {
		return err
	}
	return nil
}
