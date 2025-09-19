package main

import (
	"context"
	"flag"
	"fmt"
	"math/big"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"testing"

	"github.com/tmaxmax/go-sse"
)

func TestSEE(t *testing.T) {
	var sub string
	flag.StringVar(&sub, "sub", "all", "The topics to subscribe to. Valid values are: all, numbers, metrics")
	flag.Parse()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	wg := sync.WaitGroup{}
	fn := func(userId string) {
		defer wg.Done()
		r, _ := http.NewRequestWithContext(ctx, http.MethodGet, getRequestURL(sub), http.NoBody)
		r.Header.Set("X-User-Id", userId)
		conn := sse.NewConnection(r)

		conn.SubscribeToAll(func(event sse.Event) {
			switch event.Type {
			case "cycles", "ops":
				fmt.Printf("Metric %s: %s\n", event.Type, event.Data)
			case "close":
				fmt.Println("Server closed!")
				cancel()
			default: // no event name
				var sum, num big.Int
				for _, n := range strings.Split(event.Data, "\n") {
					_, _ = num.SetString(n, 10)
					sum.Add(&sum, &num)
				}

				fmt.Printf("Sum of random numbers: %s\n", &sum)
			}
		})

		if err := conn.Connect(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go fn(fmt.Sprintf("user-%d", i))
	}

	wg.Wait()
}

func getRequestURL(sub string) string {
	q := url.Values{}
	switch sub {
	case "all":
		q.Add("topic", "numbers")
		q.Add("topic", "metrics")
	case "numbers", "metrics":
		q.Set("topic", sub)
	default:
		panic(fmt.Errorf("unexpected subscription topic %q", sub))
	}

	return "http://localhost:8080/v1/agents:scheduling?" + q.Encode()
}
