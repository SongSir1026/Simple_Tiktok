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

func Set(key string, value []byte) {
	conn := GetRedisConnection()
	defer conn.Close()

	conn.Do("set", key, value)
}

func SetByUser(command string, key string, value string) (interface{}, error) {
	conn := GetRedisConnection()
	defer conn.Close()
	do, err := conn.Do(command, key, value)
	return do, err
}

func Get(key string) string {
	conn := GetRedisConnection()
	defer conn.Close()
	str, err := redis.String(conn.Do("get", key))
	if err != nil {
		fmt.Println(err)
	}
	return str
}

func Delete(key string) {
	conn := GetRedisConnection()
	defer conn.Close()
	conn.Do("del", key)
}

func Expire(key string, timeout int) {
	conn := GetRedisConnection()
	defer conn.Close()
	conn.Do("expire", key, timeout*60)
}
