package go_chat

import (
	"fmt"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-redis/redis"
)

//ConnectRedis returns a connection to posgress instance
func connectRedis() *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr:        "localhost:6379",
		Password:    "",
		DialTimeout: time.Second * 20,
		DB:          0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	return client
}

//RemoveRedisKey delets a key from the db
func RemoveRedisKey(key string) {
	redis := connectRedis()

	defer redis.Close()

	err := redis.Del(key).Err()

	if err != nil {
		spew.Dump(err)
	}
}

//ValidateRedisKey checks if a key exist and returns it's value
func ValidateRedisKey(key string) (value interface{}, ok bool) {
	redis := connectRedis()

	defer redis.Close()
	value, err := redis.Get(key).Result()

	if err != nil {
		spew.Dump(err)
		return nil, false
	}
	return value, true
}

//SetRedisKey set a redis key and value to the application redis instance
func SetRedisKey(key string, value interface{}, expiration time.Duration) (valid bool, result interface{}) {
	redis := connectRedis()

	result, err := redis.Set(key, value, expiration).Result()

	if err != nil {
		fmt.Println(err)
		return false, nil
	}
	return true, result
}
