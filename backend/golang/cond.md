```go
// Cond implements a condition variable, a rendezvous point for goroutines waiting for or annoucing the occurence
// of an event.
//
// Each Cond has an associated Locker L (often a *Mutex or *RWMutex)
// which must be held when changing the condition and whell calling the Wait method.
//
// A Cond must not be copied after first use.
//
// In the terminology of the Go memory model, Cond arranges that a call to Broadcast or Signal "synchronizes before" any Wait call that it unblocks.
type Cond struct {
    noCopy noCopy
    // L is held while observing or changing the condition
    L Locker
    notify notifyList
    checker copyChecker
}

// NewCond returns a new Cond with Locker l.
func NewCond(l Locker) *Cond {
    return &Cond{L: l}
}

// Wait atomically unlocks c.L and suspends execution
// of the calling goroutine. After later resuming execution,
// Wait locks c.L before returning. Unlike in other systems,
// Wait cannot return unless awoken by Broadcast or Signal.
func (c *Cond) Wait() {
    c.checker.check()
    t := runtime_notifyListAdd(&c.notify)
    c.L.Unlock()
    runtime_notifyListWait(&c.notify, t)
    c.L.Lock()
}
```
