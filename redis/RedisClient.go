package redis

import (
	constant "GoTest/constant"
	"fmt"
	rds "github.com/go-redis/redis"
)

var RC *rds.Client

func initClient() (err error) {
	RC = rds.NewClient(&rds.Options{
		Addr:     constant.RedisAddress,
		Password: constant.RedisPassword, // no password set
		DB:       constant.RedisDb,       // use default DB
	})

	_, err = RC.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func init() {
	fmt.Println("connection to redis")
	err := initClient()
	if err != nil {
		fmt.Println("redis connection error :", err)
		return
	}
	fmt.Println("redis connection is ready")
}
