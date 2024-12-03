package database

import (
	"log"
	"time"

	"lumoshive-be-chap41/config"

	"github.com/go-redis/redis"
)

type Cacher struct {
	rdb      *redis.Client
	expiracy time.Duration
	prefix   string
}

func newRedisClient(url, password string, dbIndex int) *redis.Client {
	log.Printf("Initializing Redis client with URL: %s, DBIndex: %d", url, dbIndex)
	if password != "" {
		log.Println("Redis client is using password authentication.")
	} else {
		log.Println("Redis client is not using password authentication.")
	}

	client := redis.NewClient(&redis.Options{
		Addr:     url,
		Password: password,
		DB:       dbIndex,
	})

	// Test the connection with a PING command
	_, err := client.Ping().Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	} else {
		log.Println("Connected to Redis successfully.")
	}

	return client
}

func NewCacher(cfg config.Config, expiracy int) Cacher {
	log.Printf("Creating new Cacher instance with prefix: %s and expiracy: %d minutes", cfg.Redis.Prefix, expiracy)
	cache := Cacher{
		rdb:      newRedisClient(cfg.Redis.Url, cfg.Redis.Password, 0),
		expiracy: time.Duration(expiracy) * time.Minute,
		prefix:   cfg.Redis.Prefix,
	}
	return cache
}

func (c *Cacher) Push(name string, value []byte) error {
	log.Printf("Pushing to list: %s, value: %s", c.prefix+"_"+name, string(value))
	return c.rdb.RPush(c.prefix+"_"+name, string(value)).Err()
}

func (c *Cacher) Pop(name string) (string, error) {
	log.Printf("Popping from list: %s", c.prefix+"_"+name)
	return c.rdb.LPop(c.prefix + "_" + name).Result()
}

func (c *Cacher) GetLength(name string) int64 {
	length := c.rdb.LLen(c.prefix + "_" + name).Val()
	log.Printf("Length of list %s: %d", c.prefix+"_"+name, length)
	return length
}

func (c *Cacher) Set(name string, value string) error {
	log.Printf("Setting key: %s, value: %s, expiry: %s", c.prefix+"_"+name, value, c.expiracy)
	return c.rdb.Set(c.prefix+"_"+name, value, c.expiracy).Err()
}

func (c *Cacher) SaveToken(name string, value string) error {
	log.Printf("Saving token with key: %s, value: %s, expiry: %s", c.prefix+"_"+name, value, 24*time.Hour)
	return c.rdb.Set(c.prefix+"_"+name, value, 24*time.Hour).Err()
}

func (c *Cacher) Get(name string) (string, error) {
	log.Printf("Getting value for key: %s", c.prefix+"_"+name)
	return c.rdb.Get(c.prefix + "_" + name).Result()
}

func (c *Cacher) Delete(name string) error {
	log.Printf("Deleting key: %s", c.prefix+"_"+name)
	return c.rdb.Del(c.prefix + "_" + name).Err()
}

func (c *Cacher) DeleteByKey(key string) error {
	log.Printf("Deleting key: %s", key)
	return c.rdb.Del(key).Err()
}

func (c *Cacher) PrintKeys() {
	var cursor uint64
	log.Println("Printing all Redis keys:")
	for {
		var keys []string
		var err error
		keys, cursor, err = c.rdb.Scan(cursor, "", 0).Result()
		if err != nil {
			log.Fatalf("Error scanning keys: %v", err)
		}

		for _, key := range keys {
			log.Println("Key:", key)
		}

		if cursor == 0 { // no more keys
			break
		}
	}
}

func (c *Cacher) GetKeys() []string {
	var cursor uint64
	var result []string
	log.Println("Fetching all Redis keys:")
	for {
		var keys []string
		var err error
		keys, cursor, err = c.rdb.Scan(cursor, "", 0).Result()
		if err != nil {
			log.Fatalf("Error scanning keys: %v", err)
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
	log.Printf("Fetching keys with pattern: %s", pattern)
	for {
		var keys []string
		var err error
		keys, cursor, err = c.rdb.Scan(cursor, pattern, 0).Result()
		if err != nil {
			log.Fatalf("Error scanning keys: %v", err)
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
	log.Printf("Publishing message to channel: %s, message: %s", channelName, message)
	return c.rdb.Publish(channelName, message).Err()
}

func (c *Cacher) Subcribe(channelName string) (*redis.Message, error) {
	log.Printf("Subscribing to channel: %s", channelName)
	subscriber := c.rdb.Subscribe(channelName)
	message, err := subscriber.ReceiveMessage()
	if err != nil {
		log.Printf("Error subscribing to channel: %s, error: %v", channelName, err)
	}
	return message, err
}
