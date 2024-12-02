package database

import (
	"lumoshive-be-chap41/config"
	"time"

	"fmt"

	"github.com/go-redis/redis"
)

type Cacher struct {
	rdb      *redis.Client
	expiracy time.Duration
	prefix   string
}

func newRedisClient(url, password string, dbIndex int) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     url,
		Password: password,
		DB:       dbIndex,
	})
}
func NewCacher(cfg config.Config, expiracy int) Cacher {
	cache := Cacher{
		rdb:      newRedisClient(cfg.Redis.Url, cfg.Redis.Password, 0),
		expiracy: time.Duration(expiracy) * time.Minute,
		prefix:   cfg.Redis.Prefix,
	}
	return cache
}

func (c *Cacher) Push(name string, value []byte) error {
	return c.rdb.RPush(c.prefix+"_"+name, string(value)).Err()
}

func (c *Cacher) Pop(name string) (string, error) {
	return c.rdb.LPop(c.prefix + "_" + name).Result()
}

func (c *Cacher) GetLength(name string) int64 {
	return c.rdb.LLen(c.prefix + "_" + name).Val()
}

func (c *Cacher) Set(name string, value string) error {
	return c.rdb.Set(c.prefix+"_"+name, value, c.expiracy).Err()
}

func (c *Cacher) SaveToken(name string, value string) error {
	return c.rdb.Set(c.prefix+"_"+name, value, 24*time.Hour).Err()
}

func (c *Cacher) Get(name string) (string, error) {
	return c.rdb.Get(c.prefix + "_" + name).Result()
}

func (c *Cacher) Delete(name string) error {
	return c.rdb.Del(c.prefix + "_" + name).Err()
}

func (c *Cacher) DeleteByKey(key string) error {
	return c.rdb.Del(key).Err()
}

func (c *Cacher) PrintKeys() {
	var cursor uint64
	for {
		var keys []string
		var err error
		keys, cursor, err = c.rdb.Scan(cursor, "", 0).Result()
		if err != nil {
			panic(err)
		}

		for _, key := range keys {
			fmt.Println("key", key)
		}

		if cursor == 0 { // no more keys
			break
		}
	}
}

func (c *Cacher) GetKeys() []string {
	var cursor uint64
	var result []string
	for {
		var keys []string
		var err error
		keys, cursor, err = c.rdb.Scan(cursor, "", 0).Result()
		if err != nil {
			panic(err)
		}

		result = append(result, keys...)

		if cursor == 0 { // no more keys
			break
		}
	}

	return result
}

func (c *Cacher) GetKeysByPattern(pattern string) []string {
	var cursor uint64
	var result []string
	for {
		var keys []string
		var err error
		keys, cursor, err = c.rdb.Scan(cursor, pattern, 0).Result()
		if err != nil {
			panic(err)
		}

		result = append(result, keys...)

		if cursor == 0 { // no more keys
			break
		}
	}

	return result
}

// Pub and Sub
func (c *Cacher) Publish(channelName string, message string) error {
	return c.rdb.Publish(channelName, message).Err()
}

func (c *Cacher) Subcribe(channelName string) (*redis.Message, error) {
	subscriber := c.rdb.Subscribe(channelName)
	message, err := subscriber.ReceiveMessage()
	return message, err
}
