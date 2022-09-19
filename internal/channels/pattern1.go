package channels

import (
	"fmt"
	"sync"
	"time"
)

func producer(ch chan<- string, wg *sync.WaitGroup) {
	time.Sleep(time.Second)
	ch <- "done"
	wg.Done()
}

func Pattern1() {
	wg := new(sync.WaitGroup)
	ch := make(chan string)
	for i := 0; i < 10; i++ {
		go producer(ch, wg)
		wg.Add(1)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for x := range ch {
		fmt.Println(x)
	}
}
