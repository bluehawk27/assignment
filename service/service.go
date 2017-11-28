package service

import (
	"context"
	"fmt"

	"github.com/bluehawk27/assignment/cache"

	log "github.com/Sirupsen/logrus"
	"github.com/bluehawk27/assignment/store"
)

// ServiceType : Service Layer Interface
type ServiceType interface {
	Add(ctx context.Context, key string, body []byte) error
	Get(ctx context.Context, key string) (*store.Message, error)
	Ping(ctx context.Context) (string, error)
}

// Service : Service layer Object
type Service struct {
	Store store.RedisClient
	Cache cache.Cache
}

// NewService : Instantiate New Service
func NewService() *Service {

	store := store.NewClient()
	cache := cache.NewCache()

	service := &Service{
		Store: *store,
		Cache: *cache,
	}

	return service
}

// Ping : Service layer Ping Handler
func (s *Service) Ping(ctx context.Context) (string, error) {
	resp, err := s.Store.Ping(ctx)
	if err != nil {
		return "", err
	}

	return resp, err
}

// Add : Service Layer Add Handler
func (s *Service) Add(ctx context.Context, key string, body []byte) error {

	err := s.Store.Set(ctx, key, string(body))
	if err != nil {
		log.Error("Error Adding value to Redis: ", err)
		return err
	}

	s.Cache.Set(key, string(body))

	return nil
}

// Get : Service layer Get Handler
func (s *Service) Get(ctx context.Context, key string) (*store.Message, error) {
	var (
		err error
		msg *store.Message
	)

	item, err := s.Cache.Get(key)
	if item == nil || err != nil {
		log.Info("CacheMiss:  Getting from Redis")
		msg, err = s.Store.Get(ctx, key)
		if err != nil {
			log.Error("Error Redis: ", err)
			return nil, err
		}
		s.Cache.Set(key, msg.Value)

		return msg, nil
	}
	fmt.Println("Item came from Cache:", msg)
	msg = &store.Message{
		Key:   key,
		Value: item,
	}

	return msg, nil
}
