package main

import (
	"fmt"

	"github.com/bred86/redisMonitor/auxRedis"
	"github.com/bred86/redisMonitor/auxSystem"
	"github.com/go-redis/redis"
)

func main() {
	var db int
	var usedMemory int
	var totalMemory int

	var jsonString string
	var hostname string
	var addr string
	var port string
	var passwd string
	var localIP string

	var client *redis.Client

	client = auxRedis.ConnRedis(addr, port, passwd, db)
	usedMemory = auxRedis.GetUsedMemory(client)
	totalMemory = auxRedis.GetTotalMemory(client)
	jsonString = auxRedis.GetKeyList(client)
	hostname = auxSystem.GetHostname()
	localIP = auxSystem.GetLocalIP()

	fmt.Printf("{ \"application\": \"redis_monitor\", \"slave_ip\": \"%s\", \"hostname\": \"%s\", %s\"usedMemory\": %d, \"totalMemory\": %d }\n", localIP, hostname, jsonString, usedMemory, totalMemory)

}
