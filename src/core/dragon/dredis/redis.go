package dredis

import (
	"core/dragon"
	"github.com/go-redis/redis"
	"log"
	"strconv"
	"time"
)

var (
	Redis *redis.Client
)

// init redis
func InitRedis() {
	timeout, err := strconv.Atoi(dragon.Conf.Database.Redis.Timeout)
	if err != nil {
		log.Fatalln(err)
		return
	}
	Redis = redis.NewClient(&redis.Options{
		Addr:         dragon.Conf.Database.Redis.Host + ":" + dragon.Conf.Database.Redis.Port,
		Password:     dragon.Conf.Database.Redis.Auth, // password set
		DB:           dragon.Conf.Database.Redis.Db,   // use default DB
		ReadTimeout:  time.Duration(timeout) * time.Millisecond,
		WriteTimeout: time.Duration(timeout) * time.Millisecond,
		DialTimeout:  time.Duration(timeout) * time.Millisecond,
	})
}
