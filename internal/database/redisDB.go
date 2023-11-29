package database

import (
	"fmt"
	"job-portal-api/config"
	"strconv"

	"github.com/redis/go-redis/v9"
)

func ConnectToRedis() *redis.Client {
	cfg := config.GetConfig()

	PASS, _ := strconv.Atoi(cfg.RedisConfig.Db)
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisConfig.Address,
		Password: cfg.RedisConfig.Password,
		DB:       PASS,
	})

	fmt.Println("===============", rdb)
	return rdb
}
