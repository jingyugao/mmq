package demo

import (
	"fmt"

	"github.com/jingyugao/mmq"
)

func main() {
	q, err := mmq.NewQueue("qTest")

	if err != nil {
		panic(err)
	}
	for i := 0; i < 100; i++ {
		msg, err := q.Consume()
		if err != nil {
			panic(err)
		} else {
			fmt.Printf("consume msg :%s", msg)
		}
	}
}
