package redis

import (
	"github.com/go-redis/redis/v8"
	. "login/conf"
)

var RedisDB *redis.Client

func InitRedisDB(config Config) {
	RedisConf := config.GteRedisConfig()
	RedisDB = redis.NewClient(&redis.Options{
		Addr: RedisConf.Addr,
		Password: RedisConf.Password,
		DB: RedisConf.DB,
	})
}