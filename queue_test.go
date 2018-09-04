package mmq_test

import (
	"fmt"
	"testing"

	"github.com/jingyugao/mmq"
)

func TestPut(t *testing.T) {

	q, err := mmq.NewBaseQueue("qTest")
	if err != nil {
		t.Error(err)
	}
	for i := 0; i < 100; i++ {
		err := q.Put(fmt.Sprintf("msg-%d", i))
		if err != nil {
			t.Error(err)
		} else {
			t.Log(" put msg ")
		}
	}
}

func TestConsume(t *testing.T) {
	q, err := mmq.NewBaseQueue("qTest")

	if err != nil {
		t.Error(err)
	}
	for i := 0; i < 100; i++ {
		msg, err := q.Consume()
		if err != nil {
			t.Error(err)
		} else {
			t.Logf("consume msg :%s", msg)
		}
	}
}

func TestBConsume(t *testing.T) {
	q, err := mmq.NewBaseQueue("qTest")

	if err != nil {
		t.Error(err)
	}
	for i := 0; i < 100; i++ {
		msg, err := q.BConsume(10)
		if err != nil {
			t.Error(err)
		} else {
			t.Logf("consume msg :%s", msg)
		}
	}
}
