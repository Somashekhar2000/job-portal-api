package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"job-portal-api/internal/model"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

type RDBLayer struct {
	rdb *redis.Client
}

//go:generate mockgen -source=cache.go -destination=cache_mock.go -package=cache
type Caching interface {
	AddToTheCache(ctx context.Context, jID uint, jobData model.Job) error
	GetTheCacheData(ctx context.Context, jID uint) (string, error)
}

func NewRDBLayer(rdb *redis.Client) (Caching, error) {
	if rdb == nil {
		log.Info().Msg("Redis DB cannot be nil")
		return nil, errors.New("Redis DB cannot be nil")
	}
	return &RDBLayer{
		rdb: rdb,
	}, nil
}

func (r *RDBLayer) AddToTheCache(ctx context.Context, jID uint, jobData model.Job) error {
	jobID := strconv.FormatUint(uint64(jID), 10)
	val, err := json.Marshal(jobData)
	if err != nil {
		log.Error().Err(err).Msg("error in marshaling data")
		return fmt.Errorf("error in marshaling data : %w", err)
	}
	err = r.rdb.Set(ctx, jobID, val, 10*time.Second).Err()
	return err
}

func (r *RDBLayer) GetTheCacheData(ctx context.Context, jID uint) (string, error) {
	jobId := strconv.FormatUint(uint64(jID), 10)
	str, err := r.rdb.Get(ctx, jobId).Result()
	return str, err
}
