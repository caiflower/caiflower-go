package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

func main() {
	errCh := make(chan error, 2)
	wait := sync.WaitGroup{}

	fn := func() {
		defer wait.Done()
		errCh <- errors.New("test")
	}

	wait.Add(1)
	go fn()
	time.Sleep(3 * time.Second)
	wait.Add(1)
	go fn()

	wait.Wait()
	close(errCh)

	for err := range errCh {
		fmt.Printf("err: %v \n", err.Error())
	}
}
