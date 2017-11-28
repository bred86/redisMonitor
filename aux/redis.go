package aux

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/go-redis/redis"
)

func getMemory(client *redis.Client) []string {
	memoryInfo := client.Info("memory").Val()
	return strings.Split(memoryInfo, "\n")
}

// ConnRedis - (*redis.Client) Connect to redis
func ConnRedis(addr string, port string, passwd string, db int) *redis.Client {
	if addr == "" {
		addr = "localhost"
	}

	if port == "" {
		port = "6379"
	}

	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", addr, port),
		Password: passwd,
		DB:       db,
	})
}

// GetUsedMemory - (int) Get used memory
func GetUsedMemory(client *redis.Client) int {
	lines := getMemory(client)
	usedMemoryStr := strings.Replace(strings.Split(lines[1], ":")[1], "\r", "", -1)
	usedMemoryInt, err := strconv.Atoi(usedMemoryStr)
	if err != nil {
		usedMemoryInt = 99999999
	}
	return usedMemoryInt
}

// GetTotalMemory - (int) Get used memory
func GetTotalMemory(client *redis.Client) int {
	lines := getMemory(client)
	totalMemoryStr := strings.Replace(strings.Split(lines[16], ":")[1], "\r", "", -1)
	totalMemoryInt, err := strconv.Atoi(totalMemoryStr)
	if err != nil {
		totalMemoryInt = 99999999
	}
	return totalMemoryInt
}

// GetKeyList - (string) Get a list of redis' keys and their length
func GetKeyList(client *redis.Client) string {
	var cursor uint64
	var keys []string
	var buffer bytes.Buffer
	var errorScan error

	for {
		keys, cursor, errorScan = client.Scan(cursor, "", 10).Result()
		if errorScan != nil {
			panic(errorScan)
		}

		for _, value := range keys {
			if !strings.Contains(buffer.String(), fmt.Sprintf("\"%s\":", value)) {
				buffer.WriteString(fmt.Sprintf("\"%s\":%d,", value, client.LLen(value).Val()))
			}
		}

		if cursor == 0 {
			break
		}
	}

	return buffer.String()
}

// PushToRedis - ()
func PushToRedis(client *redis.Client, key string, msg string) {
	result := client.RPush(key, msg)
	if result.Err() != nil {
		fmt.Println(result.Err())
	}
}
