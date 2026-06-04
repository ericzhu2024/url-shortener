package cache

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

var Client *redis.Client
var Ctx = context.Background()

func Init() {
	Client = redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR"),
	})

	if err := Client.Ping(Ctx).Err(); err != nil {
		log.Fatal("Fail to connect to Redis: ", err)
	}
	log.Println("Redis connected successfully")
}


func Set(key string, value string, expiration time.Duration) error {
	return Client.Set(Ctx, key, value, expiration).Err()
}

func Get(key string) (string, error) {
	return Client.Get(Ctx, key).Result()
}


func Delete(key string) error {
	return Client.Del(Ctx, key).Err()
}


func SetOriginal(original string, code string, expiration time.Duration) error {
	key := "original:" + original
	return Client.Set(Ctx, key, code, expiration).Err()
}

func GetOriginal(original string) (string, error) {
	key := "original:" + original
	return Client.Get(Ctx, key).Result()
}

func Refresh(key string, expiration time.Duration) error {
	return Client.Expire(Ctx, key, expiration).Err()
}


func RateLimit(ip string, limit int) (bool, error) {
	key := "rate:" + ip

	count, err := Client.Incr(Ctx, key).Result()
	if err != nil {
		return false, err
	}


	if count == 1 {
		Client.Expire(Ctx, key, time.Minute)
	}

	if count > int64(limit) {
		return false, nil 
	}

	return true, nil 
}