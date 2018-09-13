package mmq

import (
	"fmt"
	"strings"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/simplejia/clog"
)

type BaseQueue struct {
	Name string
}

func NewBaseQueue(name string) (bq *BaseQueue, err error) {

	rc := RedisConnPool.Get()
	defer rc.Close()
	name = GetQueueKey(name)
	ret, err := redis.Int(rc.Do("SADD", GetQueueSetKey(), name))
	if err != nil {
		clog.Error(" Create BaseQueue err: %v, %v", err, ret)
		return
	}
	if ret == 0 {
		clog.Info(" BaseQueue exists : %s", name)
	}

	bq = &BaseQueue{Name: name}

	return
}

func (bq *BaseQueue) Put(msg string) (err error) {
	rc := RedisConnPool.Get()
	defer rc.Close()

	ret, err := redis.Int(rc.Do("LPUSH", bq.Name, msg))
	if err != nil || ret != 1 {
		clog.Error(" Put msg err: %v, %v ,%v", err, ret, msg)
	}

	return
}

func (bq *BaseQueue) Consume() (msg string, err error) {
	rc := RedisConnPool.Get()
	defer rc.Close()

	msg, err = redis.String(rc.Do("RPOP", bq.Name))
	if err != nil {
		clog.Error(" Put msg err: %v, %v ,%v", err, msg, bq.Name)
	}

	return
}

func (bq *BaseQueue) BConsume(tout time.Duration) (msg string, err error) {
	rc := RedisConnPool.Get()
	defer rc.Close()

	rep, err := redis.Strings(rc.Do("BRPOP", bq.Name, tout))
	if err != nil {
		clog.Error(" Put msg err: %v, %v ,%v", err, msg, bq.Name)

		if strings.LastIndexAny(err.Error(), "nil returned") != -1 {
			return "", fmt.Errorf("timeout")
		} else {
			return "", err
		}
	}
	msg = rep[1]

	return
}
