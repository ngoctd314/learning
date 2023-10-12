package leetcode

import "strings"

type set map[string]struct{}

func newSet() set {
	return make(set)
}

var emptyStruct = struct{}{}

func (s set) add(keys ...string) {
	for _, key := range keys {
		s[strings.TrimSpace(key)] = emptyStruct
	}
}
