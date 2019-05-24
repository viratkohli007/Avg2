package db 
import (
	 "github.com/gomodule/redigo/redis"
)

func newPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000, // max number of connections
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":6379")
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}
var Pool = newPool()
var Conn = Pool.Get()

func Get(key string) (string, error) {
	value, err := redis.Bytes(Conn.Do("get", key))
	return string(value), err
}

func Set(key string, value string) error {
	_, err := Conn.Do("set", key, value)
	return err
}

func Del(key string) error {
	_, err := Conn.Do("del", key)
	return err
}

func HGet(hash string, key string) (string, error) {
	value, err := redis.Bytes(Conn.Do("hget", hash, key))
	return string(value), err
}

func HSet(hash string, key string, value string) error {
	_, err := Conn.Do("hset", hash, key, value)
	return err
}

func HDel(hash string, key string) error {
	_, err := Conn.Do("hdel", hash, key)
	return err
}

func HGetAll(hash string) (map[string]string, error) {
	values, err := redis.StringMap(Conn.Do("hgetall", hash))
	return values, err
}

