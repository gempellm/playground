package channels

import (
	"fmt"
	"time"
)

func consumer(ch <-chan string, i int) {
	answer := <-ch
	fmt.Printf("This is channel %d: %v\n", i, answer)
}

func Pattern2() {
	ch := make(chan string)
	for i := 0; i < 10; i++ {
		go consumer(ch, i)
	}

	time.Sleep(time.Second * 2)
	for i := 0; i < 10; i++ {
		ch <- fmt.Sprintf("Hello %d", i)
	}
	// ch <- "Hello 1"
	// ch <- "Hello 2"
	// ch <- "Hello 3"
	// ch <- "Hello 4"
	// ch <- "Hello 5"
	// ch <- "Hello 6"
	time.Sleep(time.Second)
}
