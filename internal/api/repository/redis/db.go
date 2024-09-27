package redis

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/mavissig/GUC.DynamicPricing-API/internal/api/domain"
	"github.com/mavissig/GUC.DynamicPricing-API/internal/api/repository"
	"log"
)

type Redis struct {
	client *redis.Client
}

func New(cfg *repository.Config) *Redis {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.PGRedis.Address,
		Password: cfg.PGRedis.Password,
		DB:       0,
	})
	return &Redis{
		client: client,
	}
}

func (r *Redis) GetDataByKey(key string) (*domain.Data, error) {
	ctx := context.Background()
	b, err := r.client.Get(ctx, key).Bytes()
	if err == redis.Nil {
		return nil, errors.New("not found")
	} else if err != nil {
		log.Println("[API][REPOSITORY][REDIS][GET] ERROR ", err)
		return nil, err
	}

	// Log the raw data for debugging
	fmt.Printf("Raw data from Redis: %s\n", string(b))

	id, err := uuid.Parse(key)
	if err != nil {
		return nil, err
	}

	res := &domain.Data{
		ID:   id,
		Data: b,
	}

	return res, nil
}
