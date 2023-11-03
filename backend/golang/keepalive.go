package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"time"
)

func client() {
	httpclient := &http.Client{
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				log.Println("re-dial")
				return net.Dial(network, addr)
			},
			DisableKeepAlives: false,
			MaxConnsPerHost:   100,
			MaxIdleConns:      1000,
			IdleConnTimeout:   time.Second * 10,
		},
	}
	for i := 0; i < 10; i++ {
		go httpclient.Get("http://localhost:8080")

	}
	time.Sleep(time.Second * 5)
}

func server() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// disable keep-alive connection
		w.Header().Set("Connection", "close")
		w.Write([]byte("pong"))
	})
	http.ListenAndServe(":8080", nil)
}
