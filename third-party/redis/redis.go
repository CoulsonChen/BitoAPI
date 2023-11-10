package redis

import (
	"github.com/redis/go-redis/v9"
)

var instance *redis.Client

func Initial() {
	instance = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "your_password", // no password set
		DB:       0,               // use default DB
	})
}

func GetInstance() *redis.Client {
	return instance
}
