package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	var keys []string
	var cursor uint64
	var err error

	keys, cursor, err = client.Scan(cursor, "", 10).Result()
	if err != nil {
		panic(err)
	}

	fmt.Println(keys)
	fmt.Println(cursor)
}
