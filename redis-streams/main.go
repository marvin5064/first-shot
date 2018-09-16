package main

import (
	"fmt"
	"log"

	"github.com/go-redis/redis"
)

const (
	RedisHost = "127.0.0.1"
	RedisPort = 6379
)

func main() {
	redisdb := setupRedisDB()
	tryXADD(redisdb)
}

func setupRedisDB() *redis.Client {
	redisdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%v:%v", RedisHost, RedisPort),
		DB:   0, // use default DB
	})

	pong, err := redisdb.Ping().Result()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("redis is connected ... ", pong)
	return redisdb
}
