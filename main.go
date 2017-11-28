package main

import (
	"bytes"
	"fmt"
	"time"

	"github.com/bred86/redisMonitor/aux"
	"github.com/go-redis/redis"
)

func main() {
	var buffer bytes.Buffer
	var clientFrom *redis.Client
	var clientTo *redis.Client
	var config aux.Config
	var hostname string
	var interval int
	var listKeys string
	var localIP string
	var totalMemory int
	var usedMemory int

	config = aux.ReadConfigFile("config/config.json")
	if config.FromRedis == nil {
		panic("Needs Redis params to read from")
	}

	if (config.Interval != nil) && (*config.Interval != 0) {
		interval = *config.Interval
	} else {
		interval = 10
	}

	clientFrom = aux.ConnRedis(
		config.FromRedis.Addr,
		config.FromRedis.Port,
		config.FromRedis.Passwd,
		config.FromRedis.Db,
	)

	if config.ToRedis != nil {
		clientTo = aux.ConnRedis(
			config.ToRedis.Addr,
			config.ToRedis.Port,
			config.ToRedis.Passwd,
			config.ToRedis.Db,
		)
	}

	for {
		now := time.Now()
		_, _, sec := now.Clock()

		if (sec % interval) == 0 {
			usedMemory = aux.GetUsedMemory(clientFrom)
			totalMemory = aux.GetTotalMemory(clientFrom)
			listKeys = aux.GetKeyList(clientFrom)
			hostname = aux.GetHostname()
			localIP = aux.GetLocalIP()

			buffer.Reset()
			buffer.WriteString("{")
			buffer.WriteString(fmt.Sprintf("\"date\":\"%s\",", now.Format(time.RFC3339)))
			buffer.WriteString("\"application\":\"redis_monitor\",")
			buffer.WriteString(fmt.Sprintf("\"slave_ip\":\"%s\",", localIP))
			buffer.WriteString(fmt.Sprintf("\"hostname\":\"%s\",", hostname))
			buffer.WriteString(fmt.Sprintf("%s", listKeys))

			if config.Team != nil {
				buffer.WriteString(fmt.Sprintf("\"team\":\"%s\",", *config.Team))
			}

			if config.Type != nil {
				buffer.WriteString(fmt.Sprintf("\"cluster_type\":\"%s\",", *config.Type))
			}

			buffer.WriteString(fmt.Sprintf("\"usedMemory\":%d,", usedMemory))
			buffer.WriteString(fmt.Sprintf("\"totalMemory\":%d", totalMemory))
			buffer.WriteString("}")

			if clientTo != nil {
				aux.PushToRedis(clientTo, config.ToRedis.Key, buffer.String())
			}

			if (config.Output != nil) && (*config.Output == true) {
				fmt.Println(buffer.String())
			}
		}
		time.Sleep(1 * time.Second)
	}
}
