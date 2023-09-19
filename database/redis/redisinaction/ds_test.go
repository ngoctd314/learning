package redisinaction

import (
	"reflect"
	"testing"
)

func Test_LPush(t *testing.T) {
	type field[T ~string] struct {
		val  string
		ll   *RedisList[T]
		want *RedisList[T]
	}

	tests := []struct {
		name  string
		field field[string]
	}{
		{
			name: "Test 1",
			field: field[string]{
				val: "1",
				ll:  &RedisList[string]{},
				want: &RedisList[string]{
					head: &Node[string]{
						val: "1",
					},
					tail: nil,
				},
			},
		},
		{
			name: "Test 2",
			field: field[string]{
				val: "1",
				ll: &RedisList[string]{
					head: &Node[string]{val: "1"},
				},
				want: &RedisList[string]{
					head: &Node[string]{
						val:  "1",
						next: &Node[string]{val: "2"},
					},
					tail: &Node[string]{val: "2"},
				},
			},
		},
		{
			name: "Test 3",
			field: field[string]{
				val: "3",
				ll: &RedisList[string]{
					head: &Node[string]{
						val:  "1",
						next: &Node[string]{val: "2"},
					},
				},
				want: &RedisList[string]{
					head: &Node[string]{
						val: "1",
						next: &Node[string]{
							val:  "2",
							next: &Node[string]{val: "3"},
						},
					},
					tail: &Node[string]{val: "3"},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.field.ll.LPush("1")
			if reflect.DeepEqual(tt.field.ll, tt.field.want) {
				t.Errorf("ll.LPush() want %v, got %v", tt.field.want, tt.field.ll)
			}
		})
	}
}
