package config

import (
	"encoding/json"
	"os"
    "strconv"

	"modak_ratelimit/internal/app/entity"
	"modak_ratelimit/internal/app/utils/logger"
)

var jsonString = `{
    "server": {
        "host": "",
        "port": "",
        "goenv": ""
    },
    "kvs": {
        "addr": "localhost:6379",
        "password": "",
        "db": 0
    },
    "RateLimitRules": [
        {
            "flowid": "notifications",
            "settings": [
                {
                    "key": "status",
                    "maxrequests": 2,
                    "timeinterval": 1
                },
                {
                    "key": "news",
                    "maxrequests": 1,
                    "timeinterval": 1440
                },
                {
                    "key": "marketing",
                    "maxrequests": 3,
                    "timeinterval": 60
                }
            ]
        },
        {
            "flowid": "another_example",
            "settings": [
                {
                    "key": "tes1",
                    "maxrequests": 1,
                    "timeinterval": 1
                },
                {
                    "key": "test2",
                    "maxrequests": 10,
                    "timeinterval": 1
                },
                {
                    "key": "test3",
                    "maxrequests": 20,
                    "timeinterval": 1
                }
            ]
        }
    ]
}`

var App entity.Config

func LoadConfig() error {

    err := json.Unmarshal([]byte(jsonString), &App)
	if  err != nil {
		logger.Error("Error with configuration JSON: ", err)
		return err
	}

	App.Server.GoEnv = os.Getenv("GO_ENV")

	App.Server.Port = os.Getenv("PORT")
    if App.Server.Port == "" {
		App.Server.Port = "8080"
	}

	App.Kvs.Addr = os.Getenv("REDIS_ADD")
    App.Kvs.Password = os.Getenv("REDIS_PASSWORD")
    kvsDb := os.Getenv("REDIS_DB")
    if kvsDb == "" {
        App.Kvs.Db = 0
    }else{
        App.Kvs.Db, err =  strconv.Atoi(kvsDb)
        if err != nil {
            logger.Error("Error KVS DB: ", err)
            return err
        }   
    }
    

	return nil
}
