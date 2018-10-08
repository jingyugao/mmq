package mmq

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

type Dlock struct {
	name    string
	token   string
	timeout time.Duration
}

func NewDlock(name string, token string, timeout time.Duration) (l *Dlock) {

	return &Dlock{name: name, token: token, timeout: timeout}
}

func (l *Dlock) tryLock() (ok bool, err error) {
	c := RedisConnPool.Get()
	defer c.Close()
	status, err := redis.String(c.Do("SET", l.key(), l.token, "EX", l.timeout/time.Second, "NX"))

	if err != nil {
		if err == redis.ErrNil {
			err = nil
		}
		return false, err
	}

	return status == "OK", nil
}

func (l *Dlock) Unlock() (err error) {
	c := RedisConnPool.Get()
	defer c.Close()
	_, err = c.Do("del", l.key())
	return
}

func (l *Dlock) key() string {
	return fmt.Sprintf("redislock:%s", l.name)
}
