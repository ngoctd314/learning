package main

import (
	"reflect"
	"testing"
)

func Test_binaryTree_insert(t *testing.T) {
	type fields struct {
		root *binaryNode
	}
	type args struct {
		data int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *binaryTree
	}{
		{
			name:   "Test 1",
			fields: fields{},
			args: args{
				data: 1,
			},
			want: &binaryTree{
				root: &binaryNode{
					data: 1,
				},
			},
		},
		{
			name: "Test insert left",
			fields: fields{
				root: &binaryNode{
					data: 100,
				},
			},
			args: args{
				data: 50,
			},
			want: &binaryTree{
				root: &binaryNode{
					left: &binaryNode{
						data: 50,
					},
					data: 100,
				},
			},
		},
		{
			name: "Test insert right",
			fields: fields{
				root: &binaryNode{
					data: 100,
				},
			},
			args: args{
				data: 150,
			},
			want: &binaryTree{
				root: &binaryNode{
					right: &binaryNode{
						data: 150,
					},
					data: 100,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &binaryTree{
				root: tt.fields.root,
			}
			if got := tr.insert(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("binaryTree.insert() = %v, want %v", got, tt.want)
			}
		})
	}
}
