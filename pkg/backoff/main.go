package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/cenkalti/backoff/v4"
)

func testBackoff() {
	operation := func() error {
		return errors.New("test")
	}

	notify := func(err error, time time.Duration) {
		fmt.Println("notify")
		fmt.Printf("err %v time:%v\n", err, time)
	}

	timeout, _ := context.WithTimeout(context.TODO(), time.Second*10)

	backoff.RetryNotify(operation, backoff.WithContext(backoff.WithMaxRetries(backoff.NewExponentialBackOff(), 100), timeout), notify)
}

func main() {
	testBackoff()
}
