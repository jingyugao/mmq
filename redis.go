package mmq

import (
	"fmt"
	"log"
	"time"

	"github.com/garyburd/redigo/redis"
)

// Host :redis host
type Host struct {
	Addr string
	Port int64
	Db   int
}

var h *Host

// RedisConnPool :redis pool
var RedisConnPool *redis.Pool

func initRedisConnPool() {

	redisServer := fmt.Sprintf("%s:%d", h.Addr, h.Port)
	RedisConnPool = &redis.Pool{
		MaxIdle:     1,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisServer)
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: pingRedis,
	}
}

func pingRedis(c redis.Conn, t time.Time) error {
	_, err := c.Do("ping")
	if err != nil {
		log.Println("[ERROR] ping redis fail", err)
	}
	return err
}

func init() {
	h = &Host{
		Addr: "127.0.0.1",
		Port: 6379,
		Db:   0,
	}
	initRedisConnPool()
}
