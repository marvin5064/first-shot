package main

import (
	"log"

	"github.com/go-redis/redis"
)

func tryXRANGE(redisdb *redis.Client, streamsName, start, end string) {
	result, err := redisdb.XRange(streamsName, start, end).Result()
	if err != nil {
		log.Println("redisdb.XRange", streamsName, start, end)
		return
	}

	log.Println("redisdb.XRange", streamsName, start, end, ":")
	for _, msg := range result {
		log.Println("\t", msg)
	}
}
