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
	AddOTP(ctx context.Context, otp string, emailID string) error
	GetOTP(ctx context.Context, otp string) (string, error)
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
	if err != nil {
		log.Err(err).Msg("error in adding job to redis")
		return err
	}
	return nil
}

func (r *RDBLayer) GetTheCacheData(ctx context.Context, jID uint) (string, error) {
	jobId := strconv.FormatUint(uint64(jID), 10)
	str, err := r.rdb.Get(ctx, jobId).Result()
	if err != nil {
		log.Err(err).Msg("error in getting job from redis")
		return "", err
	}
	return str, nil
}

func (r *RDBLayer) AddOTP(ctx context.Context, emailID string, otp string) error {
	err := r.rdb.Set(ctx, emailID, otp, 5*time.Minute).Err()
	fmt.Println("=============", err)
	if err != nil {
		log.Err(err).Msg("error in adding otp to redis")
		return err
	}
	return nil
}

func (r *RDBLayer) GetOTP(ctx context.Context, email string) (string, error) {
	str, err := r.rdb.Get(ctx, email).Result()
	if err != nil {
		log.Err(err).Msg("error in getting otp from redis")
		return "", err
	}
	return str, nil
}
