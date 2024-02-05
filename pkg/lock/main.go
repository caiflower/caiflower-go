package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

}

func testBlock() {
	mutex := sync.Mutex{}
	ch := make(chan int)

	go func() {
		mutex.Lock()
		defer mutex.Unlock()
		time.Sleep(10 * time.Second)
		ch <- 1
	}()

	time.Sleep(1 * time.Second)
	lock := mutex.TryLock()
	fmt.Printf("%v\n", lock)

	if lock {
		mutex.Unlock()
	}

	<-ch
	close(ch)
}
