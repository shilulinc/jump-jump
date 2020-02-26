package db

import (
	"github.com/astaxie/beego"
	"github.com/go-redis/redis"
)

var client *redis.Client

func GetRedisClient() *redis.Client {
	if client == nil {
		dbIdx, _ := beego.AppConfig.Int("redis_db")
		client = redis.NewClient(&redis.Options{
			Addr:        beego.AppConfig.String("redis_addr"),
			Password:    beego.AppConfig.String("redis_password"),
			DB:          dbIdx,
			IdleTimeout: -1,
		})
	}
	return client
}
