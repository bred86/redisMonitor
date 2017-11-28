package main

import (
	"bytes"
	"fmt"
	"time"

	"github.com/bred86/redisMonitor/aux"
	"github.com/go-redis/redis"
)

func main() {
	var addr string
	var buffer bytes.Buffer
	var client *redis.Client
	var config aux.Config
	var db int
	var hostname string
	var listKeys string
	var localIP string
	var passwd string
	var port string
	var totalMemory int
	var usedMemory int

	client = aux.ConnRedis(addr, port, passwd, db)
	config = aux.ReadConfigFile("config/config.json")
	if config.FromRedis == nil {
		panic("Needs Redis params to read from")
	}

	for {
		now := time.Now()
		_, _, sec := now.Clock()

		if (sec % 10) == 0 {
			usedMemory = aux.GetUsedMemory(client)
			totalMemory = aux.GetTotalMemory(client)
			listKeys = aux.GetKeyList(client)
			hostname = aux.GetHostname()
			localIP = aux.GetLocalIP()

			buffer.WriteString(fmt.Sprintf("{\"date\":\"%s\",", now.Format(time.RFC3339)))
			buffer.WriteString("\"application\":\"redis_monitor\",")
			buffer.WriteString(fmt.Sprintf("\"slave_ip\":\"%s\",", localIP))
			buffer.WriteString(fmt.Sprintf("\"hostname\":\"%s\",", hostname))
			buffer.WriteString(fmt.Sprintf("%s", listKeys))
			buffer.WriteString(fmt.Sprintf("\"usedMemory\":%d,", usedMemory))
			buffer.WriteString(fmt.Sprintf("\"totalMemory\":%d}", totalMemory))

			fmt.Println(buffer.String())
		}
		time.Sleep(1 * time.Second)
	}
}
