package main

import (
	"log"

	"github.com/go-redis/redis"
)

func tryXADD(redisdb *redis.Client) {
	keyName := "redis_streams_test_xadd"
	err := redisdb.Del(keyName).Err()
	if err != nil {
		log.Println("redisdb.Del", err)
		return
	}
	tryXRANGE(redisdb, keyName, "-", "+")
	input1 := &redis.XAddArgs{
		Stream:       keyName,
		MaxLen:       3,
		MaxLenApprox: 3,
		Values: map[string]interface{}{
			"test1": "test1",
			"test2": false,
			"test3": 3,
		},
	}
	log.Println("first input", input1)
	err = redisdb.XAdd(input1).Err()
	if err != nil {
		log.Println("redisdb.XAdd", err)
		return
	}
	tryXRANGE(redisdb, keyName, "-", "+")

	input2 := &redis.XAddArgs{
		Stream:       keyName,
		MaxLen:       3,
		MaxLenApprox: 3,
		Values: map[string]interface{}{
			"test1": "test1",
			"test2": false,
			"test4": 91.1,
		},
	}
	log.Println("second input", input2)
	err = redisdb.XAdd(input2).Err()
	if err != nil {
		log.Println("redisdb.XAdd", err)
		return
	}
	tryXRANGE(redisdb, keyName, "-", "+")

	input3 := &redis.XAddArgs{
		Stream:       keyName,
		MaxLen:       3,
		MaxLenApprox: 3,
		Values: map[string]interface{}{
			"test1": "test1",
			"test3": 3,
			"test4": 91.1,
		},
	}
	log.Println("thid input", input3)
	err = redisdb.XAdd(input3).Err()
	if err != nil {
		log.Println("redisdb.XAdd", err)
		return
	}
	tryXRANGE(redisdb, keyName, "-", "+")

	input4 := &redis.XAddArgs{
		Stream:       keyName,
		MaxLen:       3,
		MaxLenApprox: 3,
		Values: map[string]interface{}{
			"test2": false,
			"test3": 3,
			"test4": 91.1,
		},
	}
	log.Println("forth input", input4)
	err = redisdb.XAdd(input4).Err()
	if err != nil {
		log.Println("redisdb.XAdd", err)
		return
	}
	tryXRANGE(redisdb, keyName, "-", "+")
}
