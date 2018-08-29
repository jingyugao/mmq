package mmq_test

import (
	"fmt"
	"testing"

	"github.com/jingyugao/mmq"
)

func PutTest(t *testing.T) {
	q, err := mmq.NewQueue("qTest")
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

func ConsumeTest(t *testing.T) {
	q, err := mmq.NewQueue("qTest")

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
