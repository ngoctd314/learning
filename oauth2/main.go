package main

import (
	"log"
	"time"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
func process(from, to int64) []int {
	// githubLogin()
	// mockPKCE()
	var secondInADay int64 = 86400
	nextFrom2Day := unixAtOOh00m00s(from) + secondInADay*2
	rs := []int{}
	if to < nextFrom2Day {
		rs = append(rs, 1)
	}
	if to > nextFrom2Day {
		rs = append(rs, 2)
	}
	if to > nextFrom2Day {
		rs = append(rs, 3)
	}
	if to > nextFrom2Day {
		rs = append(rs, 4)
	}

	return rs
}

func unixAtOOh00m00s(ts int64) int64 {
	date := time.Unix(ts, 0).UTC()
	return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC).Unix()
}
