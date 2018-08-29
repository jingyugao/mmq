package mmq

import "github.com/garyburd/redigo/redis"

type StableQueue struct {
	P *BaseQueue
	W *BaseQueue
}

func NewStableQueue(name string) (sq *StableQueue, err error) {
	p, err := NewBaseQueue("P::" + name)
	if err != nil {
		return
	}
	w, err := NewBaseQueue("W::" + name)
	if err != nil {
		return
	}

	sq = &StableQueue{P: p, W: w}

	return
}

func (sq *StableQueue) Put(msg string) (err error) {
	rc := RedisConnPool.Get()
	defer rc.Close()

	err = sq.P.Put(msg)

	return
}

func (sq *StableQueue) Consume() (msg string, err error) {
	rc := RedisConnPool.Get()
	defer rc.Close()

	msg, err = redis.String(rc.Do("RPOPLPUSH", sq.P.Name, sq.W.Name))

	return
}
