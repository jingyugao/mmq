package mmq

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

type StableQueue struct {
	P              *BaseQueue
	W              *BaseQueue
	resendInterval time.Duration
	exitChan       *chan int
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

	err = json.Unmarshal(msgRaw, &msg)
	return
}

func (sq *StableQueue) ACK(msg Msg) (err error) {
	rc := RedisConnPool.Get()
	defer rc.Close()

	msgRaw, err := json.Marshal(msg)
	if err != nil {
		return
	}
	//ret, err := redis.Int(rc.Do("eval \"redis.call('lrem',KEYS[1],1,ARGV[1])\"", 1, "l1", "v1"))
	ret, err := redis.Int(rc.Do("lrem", sq.W.Name, msgRaw))

	if err != nil {
		return
	}
	if ret == 0 {
		err = fmt.Errorf("has ACKed")
	}
	return
}

func reSendLoop(sq *StableQueue) {
	rc := RedisConnPool.Get()
	defer rc.Close()

	//重新投递
	for {
		select {
		case <-time.Tick(time.Second):
			rc.Do("EVAL", "local json = redis.call('GET', KEYS[1]) local obj = cjson.decode(json) return obj['hotelId']")

		}
	}
}

func init() {

}
