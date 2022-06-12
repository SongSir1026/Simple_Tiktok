package utils

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func GetRedisConnection() redis.Conn {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return nil
	}

	fmt.Println("redis conn success")

	return c
}
