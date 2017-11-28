package aux

import (
	"encoding/json"
	"os"
)

// Config - Json config file
type Config struct {
	FromRedis     *redisType
	ToRedis       *redisType
	Elasticsearch *elasticsearchType
	Team          *string
	Type          *string
	Interval      *int
	Output        *bool
}

type redisType struct {
	Addr   string
	Port   string
	Passwd string
	Db     int
	Key    string
}

type elasticsearchType struct {
	Addr  string
	Port  string
	Index string
}

// ReadConfigFile - (Config)
func ReadConfigFile(filePath string) Config {
	var config Config

	configFile, err := os.Open(filePath)
	defer configFile.Close()
	if err != nil {
		panic(err.Error())
	}

	jsonConf := json.NewDecoder(configFile)
	jsonConf.Decode(&config)

	return config
}
