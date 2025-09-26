package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"testing"

	"github.com/tmaxmax/go-sse"
)

const (
	EventTypeOfChatModelAnswer = "chat.answer"
	EventTypeOfChatError       = "chat.error"
	EventTypeOfChatFinish      = "chat.finish"
)

func TestSEE(t *testing.T) {
	var sub string
	flag.StringVar(&sub, "sub", "all", "The topics to subscribe to. Valid values are: all, numbers, metrics")
	flag.Parse()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	wg := sync.WaitGroup{}
	msg := ""
	fn := func(userId string) {
		defer wg.Done()
		r, _ := http.NewRequestWithContext(ctx, http.MethodGet, getRequestURL(sub, userId), http.NoBody)
		r.Header.Set("X-User-Id", userId)
		r.Header.Set("Last-Event-Id", "1")
		conn := sse.NewConnection(r)

		conn.SubscribeToAll(func(ev sse.Event) {
			switch ev.Type {
			case EventTypeOfChatModelAnswer:
				msg += ev.Data
			case EventTypeOfChatError:
				cancel()
			case EventTypeOfChatFinish:
				cancel()
			}
		})

		if err := conn.Connect(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		fmt.Println(msg)
	}

	for i := 0; i < 1; i++ {
		wg.Add(1)
		go fn(fmt.Sprintf("user-%d", i))
	}

	wg.Wait()
}

func getRequestURL(sub string, sessionId string) string {
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
	q.Add("sessionId", sessionId)

	return "http://127.0.0.1:8080/v1/chat?input=%E4%BB%8B%E7%BB%8D%E4%B8%80%E4%B8%8B%E5%8C%97%E4%BA%AC&chatProtocol=ollama"
}
