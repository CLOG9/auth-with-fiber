package config

import (
	"github.com/gofiber/storage/redis/v3"
)

func Rds() *redis.Storage {

	// Initialize custom config
	store := redis.New(redis.Config{
		Host:      "localhost",
		Port:      6379,
		Username:  "",
		Password:  "",
		Database:  0,
		Reset:     false,
		TLSConfig: nil,
	})
	return store
}
