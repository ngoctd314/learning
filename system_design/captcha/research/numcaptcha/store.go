package numcaptcha

import (
	"container/list"
	"sync"
	"time"
)

// Store an object implementing Store interface can be registered with SetCustomStore
// function to handle storage and retrieval of captcha ids and solution for them
// replacing the default memory store
// It is the responsibility of an object to delete expired and used captchas
type Store interface {
	// Set sets the digits for the captcha id
	Set(id string, digits []byte)
	// Get returns stored digits for the captcha id. Clear indicates
	// whether the captcha must be deleted from the store
	Get(id string, clear bool) (digits []byte)
}

type idByTimeValue struct {
	time time.Time
	id   string
}

type memoryStore struct {
	sync.RWMutex
	digitsByID map[string][]byte
	idByTime   *list.List
	// Number of items stored since last collection
	numStored int
	// number of saved items that triggers collection
	collectNum int
	// expiration time of captchas
	expiration time.Duration
}

// NewMemoryStore returns a new standard memory store for captchas with the
// given collection threshold and expiration time (duration). The returned
// store must be registered with SetCustomStore to replace the default one
func NewMemoryStore(collectNum int, expiration time.Duration) *memoryStore {
	s := &memoryStore{
		digitsByID: map[string][]byte{},
		idByTime:   list.New(),
		collectNum: collectNum,
		expiration: expiration,
	}
	return s
}

func (s *memoryStore) Set(id string, digits []byte) {
	s.Lock()
	s.digitsByID[id] = digits
	s.idByTime.PushBack(idByTimeValue{
		time: time.Now(),
		id:   id,
	})
	s.numStored++
	if s.numStored <= s.collectNum {
		s.Unlock()
		return
	}
	s.Unlock()
	go s.collect()
}
func (s *memoryStore) Get(id string, clear bool) (digits []byte) {
	if !clear {
		s.RLock()
		defer s.RUnlock()
	} else {
		s.Lock()
		defer s.Unlock()
	}
	digits, ok := s.digitsByID[id]
	if !ok {
		return
	}
	if clear {
		delete(s.digitsByID, id)
	}

	return
}

func (s *memoryStore) collect() {
	now := time.Now()
	s.Lock()
	defer s.Unlock()
	s.numStored = 0
	for e := s.idByTime.Front(); e != nil; {
		ev, ok := e.Value.(idByTimeValue)
		if !ok {
			return
		}
		if ev.time.Add(s.expiration).Before(now) {
			delete(s.digitsByID, ev.id)
			next := e.Next()
			s.idByTime.Remove(e)
			e = next
		} else {
			return
		}
	}
}
