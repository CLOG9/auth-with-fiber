package config

import (
	"log"
	"os"
	"strconv"

	"github.com/gofiber/storage/redis/v3"
)

func Rds() *redis.Storage {

	// Initialize custom config
	intValue, err := strconv.Atoi(os.Getenv("REDIS_PORT"))
	if err != nil {
		log.Fatal(err)
	}
	store := redis.New(redis.Config{
		Host:      os.Getenv("REDIS_HOST"),
		Port:      intValue,
		Username:  "",
		Password:  "",
		Database:  0,
		Reset:     false,
		TLSConfig: nil,
	})
	return store
}
