package mmq

import (
	"encoding/json"

	"github.com/garyburd/redigo/redis"
)

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

func (sq *StableQueue) Put(msg Msg) (err error) {
	rc := RedisConnPool.Get()
	defer rc.Close()
	msgRaw, err := json.Marshal(msg)
	if err != nil {
		return
	}

	err = sq.P.Put(string(msgRaw))
	return
}

func (sq *StableQueue) Consume() (msg Msg, err error) {
	rc := RedisConnPool.Get()
	defer rc.Close()

	msgRaw, err := redis.Bytes(rc.Do("RPOPLPUSH", sq.P.Name, sq.W.Name))

	err = json.Unmarshal(msgRaw, msg)
	return
}

func (sq *StableQueue) ACK(ID MsgID) (msg string, err error) {
	rc := RedisConnPool.Get()
	defer rc.Close()

	// msg, err = redis.String(rc.Do("RPOPLPUSH", sq.P.Name, sq.W.Name))
	// lua remove

	return
}
