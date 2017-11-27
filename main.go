package main

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/go-redis/redis"
)

func main() {
	var cursor uint64

	var memoryInfo string
	var usedMemory string
	var jsonString string

	var lines []string
	var keys []string

	var errorScan error

	var client *redis.Client

	var buffer bytes.Buffer

	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	memoryInfo = client.Info("memory").Val()
	lines = strings.Split(memoryInfo, "\n")
	usedMemory = strings.Replace(strings.Split(lines[1], ":")[1], "\n", "", -1)

	for {
		keys, cursor, errorScan = client.Scan(cursor, "", 10).Result()
		if errorScan != nil {
			panic(errorScan)
		}

		for _, value := range keys {
			if !strings.Contains(buffer.String(), fmt.Sprintf("\"%s\":", value)) {
				buffer.WriteString(fmt.Sprintf("\"%s\": %d, ", value, client.LLen(value).Val()))
			}
		}

		if cursor == 0 {
			break
		}
	}

	jsonString = buffer.String()

	fmt.Printf("{ %s\"usedMemory\": %s }\n", jsonString, usedMemory)

}
