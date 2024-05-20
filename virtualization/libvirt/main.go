package main

import (
	"log"
	"time"
)

func main() {
	now := time.Now()
	for i := 0; i < 10_000_000_000; i++ {
	}
	log.Printf("since: %d ms", time.Since(now).Milliseconds())
}
