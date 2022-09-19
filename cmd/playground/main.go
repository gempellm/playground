package main

import (
	"fmt"
)

func printChan(ch chan int) {
	fmt.Printf("%#+v\n", ch)
	fmt.Printf("cap(ch) = %v\n", cap(ch))
	fmt.Printf("len(ch) = %v\n", len(ch))
}

func main() {

	// c := rwcache.NewRWCache()
	// c.Storage["test"] = "test"
	// fmt.Printf("%#+v\n", c)
	// c.Set("hello", "world")
	// fmt.Println(c.Get("hello"))
	// c.Delete("hello")
	// fmt.Println(c.Get("hello"))

	// wg := sync.WaitGroup{}

	// for i := 0; i < 100; i++ {
	// 	wg.Add(1)
	// 	go func(k int) {
	// 		err := c.Set(fmt.Sprint(k), fmt.Sprint(k))
	// 		if err != nil {
	// 			fmt.Printf("error occured at i = %d, error = %e", k, err)
	// 		}

	// 		// err = c.Delete(fmt.Sprint(k - 1))
	// 		// if err != nil {
	// 		// 	fmt.Printf("error occured at i = %d, error = %e", k, err)
	// 		// }
	// 		wg.Done()
	// 	}(i)
	// 	//go c.Delete(fmt.Sprint(i))
	// }

	// wg.Wait()
	// fmt.Printf("%#+v\n", c)
	// fmt.Printf("len(storage) = %d\n", len(c.Storage))

	// for i := 0; i < 100; i++ {
	// 	wg.Add(1)
	// 	go func(k int) {
	// 		value, err := c.Get(fmt.Sprint(k))
	// 		if err != nil {
	// 			fmt.Printf("error occured at i = %d, error = %e", k, err)
	// 		}

	// 		fmt.Printf("value = %s\n", value)

	// 		// err = c.Delete(fmt.Sprint(k - 1))
	// 		// if err != nil {
	// 		// 	fmt.Printf("error occured at i = %d, error = %e", k, err)
	// 		// }
	// 		wg.Done()
	// 	}(i)
	// }

	// wg.Wait()

	// for i := 0; i < 100; i++ {
	// 	wg.Add(1)
	// 	go func(k int) {
	// 		err := c.Delete(fmt.Sprint(k))
	// 		if err != nil {
	// 			fmt.Printf("error occured at i = %d, error = %e", k, err)
	// 		}
	// 		wg.Done()
	// 	}(i)
	// }

	// wg.Wait()
	// fmt.Println("Final Storage state = ", c.Storage)

	// cs := map[string]int{"касса": 0}
	// mu := sync.Mutex{}

	// for i := 0; i < 1000; i++ {
	// 	go func(k int) {
	// 		mu.Lock()
	// 		cs["касса"] += 1
	// 		mu.Unlock()
	// 	}(i)
	// }
	// time.Sleep(time.Second)
	// fmt.Println(cs)

	// channels.Pattern1()
	// channels.Pattern2()
	// channels.Pattern3()

	// wg := new(sync.WaitGroup)
	// fmt.Printf("%#+v\n", wg)

	// var ch chan int
	// printChan(ch)

	// ch = make(chan int, 5)
	// printChan(ch)

	// someVariable := 1
	// ch <- someVariable
	// printChan(ch)

	// a := <-ch
	// printChan(ch)
	// a++

	// for i := 0; i < 5; i++ {
	// 	go func(counter int) {
	// 		time.Sleep(time.Second)
	// 		fmt.Println(counter)
	// 	}(i)
	// }
	// time.Sleep(2 * time.Second)
	// fmt.Println("end")
}
