package redisinaction

import (
	"sync"
)

// RedisString is a string type it can store string, int, float type
type RedisString[T ~string | ~int32 | ~int64 | ~int | ~float32 | ~float64] struct {
	rwm   sync.RWMutex
	store map[string]T
}

func (s *RedisString[T]) Get(k string) T {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	return s.store[k]
}

func (s *RedisString[T]) Set(k string, v T) error {
	s.rwm.Lock()
	s.store[k] = v
	s.rwm.Unlock()

	return nil
}

func (s *RedisString[T]) Del(k string) error {
	s.rwm.Lock()
	delete(s.store, k)
	s.rwm.Unlock()

	return nil
}

type Node[T ~string] struct {
	val  T
	next *Node[T]
}

type RedisList[T ~string] struct {
	head *Node[T]
	tail *Node[T]
}

// LPush push the value onto the left end of the list
func (s RedisList[T]) LPush(val T) {
	newNode := &Node[T]{val: val}
	if s.head == nil {
		s.head = newNode
	} else if s.tail == nil {
		s.head, s.head.next, s.tail = newNode, s.head, s.head
	} else {
		s.head, newNode.next = newNode, s.head.next
	}
}

// RPush push the value onto the right end of the list
func (s RedisList[T]) RPush(val T) {
	newNode := &Node[T]{val: val}
	s.tail.next, s.tail = newNode, newNode
}

// LPop pops the value from the left of the list and returns it
func (s RedisList[T]) LPop() T {
	return s.head.val
}

func (s RedisList[T]) RPop() T {
	return s.tail.val
}

// LIndex fetches an item at a given position in list
func (s RedisList[T]) LIndex(index int) T {
	cur := s.head
	for i := 0; i < index; i++ {
		cur = cur.next
	}

	return cur.val
}

// LRange fetches a range of values from the list
func (s RedisList[T]) LRange(from, to int) T {
	return s.head.val
}

type RedisSet struct {
	store map[string]any
}

func (RedisSet) SADD() {}

func (RedisSet) SMEMBERS() {}

func (RedisSet) SISMEMBER() {}

func (RedisSet) SREM() {}
