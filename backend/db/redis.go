package db

import (
	"log"
	"strconv"
	"time"

	"github.com/go-redis/redis"

	"github.com/Jim-Lin/bee-bee-alert/backend/model"
	"github.com/Jim-Lin/bee-bee-alert/backend/util"
)

var client *redis.Client

func IncrCounter(prod model.Prod) {
	key := prod.Url
	result, err := getClient().Incr(key).Result()
	util.CheckError(err)

	// https://stackoverflow.com/a/49251938/4436392
	if result == 30 {
		pipe := getClient().Pipeline()
		pipe.Expire(key, 30*time.Second)
		_, err := pipe.Exec()
		util.CheckError(err)
	}

	if result > 10 {
		log.Println(key)
		_, err := getClient().Del(key).Result()
		util.CheckError(err)

		go (&util.MailTemplate{
			Subject: "[Notice] Hot product!",
			Msg:     prod.Name + "\r\n$" + strconv.Itoa(prod.Price) + "\r\n" + prod.Url + "\r\n",
		}).GetMail().Notify()
	}
}

func getClient() *redis.Client {
	if client == nil {
		client = redis.NewClient(&redis.Options{
			Addr:         util.GetConfig().RedisUrl,
			DialTimeout:  10 * time.Second,
			ReadTimeout:  30 * time.Second,
			WriteTimeout: 30 * time.Second,
			PoolSize:     10,
			PoolTimeout:  30 * time.Second,
		})
		client.FlushDB()
	}

	return client
}
