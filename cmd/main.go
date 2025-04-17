package main

import (
	"bytes"
	"fmt"
	"net/http"
	_ "net/http/pprof"

	"k8s.io/apimachinery/pkg/util/rand"
)

func main() {
	httpServer := &http.Server{Addr: fmt.Sprintf(":8080")}
	go func() {
		for i := 0; i < 100; i++ {
			repeat(generate(100), 100)
		}
	}()

	if err := httpServer.ListenAndServe(); err != nil {
		panic(err)
	}
}

const Letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func generate(n int) string {
	var buf bytes.Buffer
	for i := 0; i < n; i++ {
		buf.WriteByte(Letters[rand.Intn(len(Letters))])
	}
	return buf.String()
}

func repeat(s string, n int) string {
	var result string
	for i := 0; i < n; i++ {
		result += s
	}

	return result
}
