package main

import (
	"fmt"
	"sync"
	"time"
)

func work() {
	time.Sleep(1 * time.Second)
	fmt.Println("thats all")
}

func main() {
	wg := new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			work()
		}(wg)
	}
	wg.Wait()
}
