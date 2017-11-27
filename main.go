package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func main() {
	var cursor uint64

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	fmt.Println(client.Info("Memory").Val()[3])

	keys := client.Scan(cursor, "", 10).Iterator()
	if keys.Err() != nil {
		panic(keys.Err())
	}

	// Print nome da fila e tamanho da fila
	for keys.Next() {
		fmt.Println(keys.Val(), client.LLen(keys.Val()))
	}
}
