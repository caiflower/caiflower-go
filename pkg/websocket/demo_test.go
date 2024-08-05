package websocket

import (
	"caiflower.com/caiflower-go/pkg/websocket/client"
	"caiflower.com/caiflower-go/pkg/websocket/server"
	"github.com/gorilla/websocket"
	"log"
	"testing"
	"time"
)

func TestStartTestDemo(t *testing.T) {
	// 启动服务端
	go server.Start()

	time.Sleep(3 * time.Second)

	// 启动客户端
	wbc, err := client.NewClient("ws://127.0.0.1:8080/ws", 3, 3)
	if err != nil {
		panic(err)
	}
	go wbc.Start()

	for i := 0; i < 10; i++ {
		err = wbc.SendMessage(websocket.TextMessage, []byte("test1"))
		if err != nil {
			panic(err)
		}
	}

	time.Sleep(30 * time.Second)

	log.Println("Close wbc")
	wbc.Close()

	time.Sleep(1000 * time.Second)
}
