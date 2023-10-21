# Examples

## Deadlock

1. Two keys

```go
// You are join a game
// This game is very simple: you must find 2 keys in 2 differents location
// After you find 2 keys. You goto lock and open it.
// Gift is in give to you after you open that lock.
func main() {
	key1 := key{}
	key2 := key{}
	// there are 10 players in a game
	for i := 0; i < 10; i++ {
		go findKeySequence(&key1, &key2, time.Second)
	}
	// block forever
	var ch chan struct{}
	<-ch
}

type key struct {
	m sync.Mutex
}

func findKeySequence(key1, key2 *key, d time.Duration) {
	// you find key1 first
	// you lock it to notify: you are currently hold key, no one can use it
	key1.m.Lock()
	// you return lock after use
	defer key1.m.Unlock()

	// you take time find key2 at another location
	time.Sleep(d)

	// similar to key1
	key2.m.Lock()
	defer key2.m.Unlock()
}
```

2. Interview deadlock

```go
// In an interview
// Interviewer: tell me about deadlock. If your question is true, you will pass this interview.
// Candidate: Allow me pass this interview. After that, i will tell you about deadlock.
// Interviewer: ...
func main() {
	passInterviewLock := sync.Mutex{}
	answerLock := sync.Mutex{}

	// interview process (interview expect this happen)
	go func() {
		passInterviewLock.Lock()
		defer passInterviewLock.Unlock()

		fmt.Println("Tell me about deadlock. If your question is true, you will pass this interview.")

		fmt.Println("Answering...")
		answerLock.Lock()
		defer answerLock.Unlock()
	}()

	// candidate process (candidate expect this happen)
	answerLock.Lock()
	defer answerLock.Unlock()

	time.Sleep(time.Second)

	fmt.Println("Allow me pass this interview. After that, i will tell you about deadlock.")
	passInterviewLock.Lock()
	defer passInterviewLock.Unlock()
}
```

## Starvation

Database delete
