package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

func main() {
	test2()
}

func test1() {
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

func test2() {
	c := make(chan int, 20)

	for i := 0; i < 1000; i++ {
		select {
		case c <- 1:
		default:
			fmt.Println("-------")
		}

	}

	time.Sleep(10 * time.Second)
}
