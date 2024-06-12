package main

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"
)

/*
TODO:
+ http.Server gracefulShutdown()
+ database gracefulShutdown()
+ background job gracefulShutdown()
+ long running request
*/

type appIDKey string

const (
	appID     appIDKey = "api-server-key"
	requestID appIDKey = "requestID"
)

func gracefulShutdown() {
	// ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	// defer stop()

	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		go func() {
			ticker := time.NewTicker(time.Millisecond * 500)
			for {
				ticker.Reset(time.Millisecond * 500)
				select {
				case <-ticker.C:
					if err := ctx.Err(); err != nil {
						slog.Error("ctx.Err()", "err", err)
					} else {
						slog.Info("still alive")
					}
				}
			}
		}()

		// for {
		time.Sleep(time.Second * 3)
		w.Write([]byte("pong"))
		// time.Sleep(time.Second)
		// 	flusher := w.(http.Flusher)
		// 	flusher.Flush()
		// }
	})

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
		// ReadTimeout:                  0,
		// ReadHeaderTimeout:            0,
		WriteTimeout: time.Second * 20,
		IdleTimeout:  time.Second * 5,
		// ConnState: func(_ net.Conn, state http.ConnState) {
		// 	slog.Info("ConnState is changed", "state", state.String())
		// },
		// ErrorLog: &log.Logger{},
		BaseContext: func(net.Listener) context.Context {
			fmt.Println("RUN ONCE")
			ctx := context.Background()
			ctx = context.WithValue(ctx, appID, "api-server-value")
			return ctx
		},
		ConnContext: func(ctx context.Context, c net.Conn) context.Context {
			fmt.Println("===========================")
			appIDValue := ctx.Value(appID)
			slog.Info("extract context", "appID", appIDValue)

			return context.WithValue(ctx, requestID, fmt.Sprint(time.Now().Unix()))
		},
	}

	go func() {
		slog.Info("ListenAndServe", "addr", server.Addr)
		if err := server.ListenAndServe(); err != nil {
			slog.Error("ListenAndServe()", "err", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	ctx := context.Background()
	server.Shutdown(ctx)

	select {
	case <-time.After(time.Second * 5):
		return
	}
}

// cancel long running requests
