package client

import (
	"context"
	"errors"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
	"time"
)

type Client struct {
	lock     sync.Mutex
	conn     *websocket.Conn
	retryCnt int
	timeout  time.Duration
	running  bool
	ctx      context.Context
	cancel   context.CancelFunc
}

func NewClient(uri string, timeout, retryCnt int) (*Client, error) {
	dialer := websocket.Dialer{
		HandshakeTimeout: time.Duration(timeout) * time.Second,
	}

	header := http.Header{
		"User-Agent": []string{"test"},
	}
	dial, _, err := dialer.Dial(uri, header)
	if err != nil {
		return nil, err
	}
	return &Client{
		conn:     dial,
		timeout:  time.Duration(timeout) * time.Second,
		retryCnt: retryCnt,
	}, nil
}

func (c *Client) Start() {
	c.lock.Lock()
	defer c.lock.Unlock()
	if c.running {
		return
	}

	go func() {
		for {
			messageType, p, err := c.conn.ReadMessage()
			if err != nil {
				log.Println("read err: ", err)
				return
			}

			if messageType == websocket.TextMessage {
				msg := string(p)
				log.Println("Msg=", msg)
			}
		}
	}()

	ctx, cancelFunc := context.WithCancel(context.Background())
	c.ctx = ctx
	c.cancel = cancelFunc
	c.running = true
	tick := time.NewTicker(time.Second * 10)
	for {
		select {
		case <-tick.C:
			err := c.SendMessage(websocket.PingMessage, []byte("testPing"))
			if err != nil {
				log.Println("Send ping message error:", err)
			}
		case <-ctx.Done():
			tick.Stop()
			return
		}
	}
}

func (c *Client) SendMessage(messageType int, message []byte) error {
	timeout, cancelFunc := context.WithTimeout(context.Background(), c.timeout)
	defer cancelFunc()
	for {
		select {
		case <-timeout.Done():
			return errors.New("timeout")
		default:
			err := c.conn.WriteMessage(messageType, message)
			if err != nil {
				time.Sleep(time.Millisecond * 100)
			} else {
				return nil
			}
		}
	}
}

func (c *Client) Close() {
	c.cancel()
	c.running = false
	err := c.conn.Close()
	if err != nil {
		log.Println("close err: ", err)
	}
}
