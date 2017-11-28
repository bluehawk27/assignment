package store

import (
	"context"
	"fmt"
	"time"

	"github.com/bluehawk27/assignment/config"
	"github.com/go-redis/redis"
)

// RedisClient : Redis "Object"
type RedisClient struct {
	Client *redis.Client
	Expiry int64
}

func setClient(address string, password string, db int) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password, // no password set if set to ""
		DB:       db,       // uses default DB if set to 0
	})

	return client
}

//  NewClient : Redis Client Instance
func NewClient() *RedisClient {
	config.Init()
	rc := config.GetRedisConfig()
	address := config.GetRedisConnectionString()
	password := rc.Password
	db := rc.DB
	client := setClient(address, password, db)

	rcClient := &RedisClient{
		Client: client,
		Expiry: rc.Expiry,
	}

	return rcClient
}

// Ping : Redis Instance Check
func (s *RedisClient) Ping(ctx context.Context) (string, error) {

	pong, err := s.Client.Ping().Result()
	if err != nil {
		return "", err
	}

	return pong, nil
}

// Set : Set the Key - Value in Redis
func (s *RedisClient) Set(ctx context.Context, key string, value interface{}) error {
	exp := time.Duration(s.Expiry) * time.Second
	err := s.Client.Set(key, value, exp).Err()
	if err != nil {
		fmt.Println("error is:", err)
		return err
	}
	return nil
}

// Get : Get the Value by Key from the Redis Instance
func (s *RedisClient) Get(ctx context.Context, key string) (*Message, error) {
	var msg *Message

	val, err := s.Client.Get(key).Result()
	if err == redis.Nil {
		msg = &Message{
			Key:   key,
			Value: "Key Does Not Exist",
		}

		return msg, nil
	} else if err != nil {
		msg = &Message{
			Key:   key,
			Value: err,
		}
		return msg, err
	}
	msg = &Message{
		Key:   key,
		Value: val,
	}

	return msg, nil
}
