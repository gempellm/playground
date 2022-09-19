package channels

import (
	"fmt"
	"time"
)

func worker(ch <-chan int) {
	for {
		v, ok := <-ch
		if !ok {
			return
		}
		fmt.Printf("%v\n", v)
		time.Sleep(time.Second)
	}
}

func Pattern3() {
	ch := make(chan int, 5)
	go worker(ch)
	// go worker(ch)
	// go worker(ch)
	for i := 0; i < 100; i++ {
		ch <- i
	}
}
