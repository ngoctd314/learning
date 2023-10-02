package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"sync"
	"time"
)

type connectionWatcher struct {
	mu sync.Mutex
	m  map[net.Conn]struct{}
}

func (cw *connectionWatcher) OnStateChange(conn net.Conn, state http.ConnState) {
	switch state {
	case http.StateNew:
		log.Println("StateNew")
		cw.mu.Lock()
		defer cw.mu.Unlock()
		if cw.m == nil {
			cw.m = make(map[net.Conn]struct{})
		}
		cw.m[conn] = struct{}{}
	case http.StateHijacked:
		log.Println("StateHijacked")
		cw.mu.Lock()
		defer cw.mu.Unlock()
		delete(cw.m, conn)
	default:
		log.Println("state", state)
	}
}

func (cw *connectionWatcher) Connections() []net.Conn {
	var result []net.Conn
	cw.mu.Lock()
	for conn := range cw.m {
		result = append(result, conn)
	}
	cw.mu.Unlock()

	return result
}

func main() {
	cw := connectionWatcher{}
	s := &http.Server{
		Addr:        "127.0.0.1:8080",
		ConnState:   cw.OnStateChange,
		IdleTimeout: time.Second,
	}

	http.HandleFunc("/ping", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("pong"))
	})
	http.HandleFunc("/connections", func(w http.ResponseWriter, _ *http.Request) {
		conns := cw.Connections()
		data, _ := json.Marshal(conns)
		w.Header().Add("Content-Type", "application/json")
		w.Write(data)
	})

	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
