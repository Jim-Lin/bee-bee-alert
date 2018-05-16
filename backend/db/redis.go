package db

import (
	"github.com/go-redis/redis"
	"log"
	"strconv"
	"time"

	"github.com/Jim-Lin/bee-bee-alert/backend/model"
	"github.com/Jim-Lin/bee-bee-alert/backend/util"
)

var redisClient *redis.Client

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:         util.GetConfig().RedisUrl,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})
	redisClient.FlushDB()
}

func IncrCounter(prod model.Prod) {
	key := prod.Url
	result, err := redisClient.Incr(key).Result()
	util.CheckError(err)

	// https://stackoverflow.com/a/49251938/4436392
	if result == 30 {
		pipe := redisClient.Pipeline()
		pipe.Expire(key, 10 * time.Second)
		_, err := pipe.Exec()
		util.CheckError(err)
	}

	if result > 1 {
		log.Println(key)
		_, err := redisClient.Del(key).Result()
		util.CheckError(err)

		go (&util.MailTemplate{
			Subject: "[Notice] Hot product!",
			Msg: prod.Name + "\r\n$" + strconv.Itoa(prod.Price) + "\r\n" + prod.Url + "\r\n",
		}).
		GetMail().
		Notify()
	}
}
