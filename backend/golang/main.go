package main

import (
	"fmt"
	"sync"
)

// In an interview
// Interviewer: tell me about deadlock. If your question is true, you will pass this interview.
// Candidate: Allow me pass this interview. After that, i will tell you about deadlock.
// Interviewer: ...
func main() {
	passInterviewLock := sync.Mutex{}
	answer := make(chan interface{})

	// interview process (interview expect this happen)
	go func() {
		passInterviewLock.Lock()
		defer passInterviewLock.Unlock()

		fmt.Println("Tell me about deadlock. If your question is true, you will pass this interview.")

		fmt.Println("Answering...")
		// waiting answer
		msg := <-answer
		fmt.Println(msg)
	}()

	// candidate process (candidate expect this happen)
	go func() {
		passInterviewLock.Lock()
		answer <- "Allow me pass this interview. After that, i will tell you about deadlock."
		defer passInterviewLock.Unlock()
	}()

	select {}
}
