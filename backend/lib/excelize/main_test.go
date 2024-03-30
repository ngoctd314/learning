package main

import (
	"reflect"
	"testing"
)

func Test_table_getCell(t *testing.T) {
	type fields struct {
		topLeft     [2]rune
		bottomRight [2]rune
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "Test 1",
			fields: fields{
				topLeft:     [2]rune{'A', 10},
				bottomRight: [2]rune{'F', 22},
			},
			want: []string{"A10", "B10", "C10", "D10", "E10", "F10"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &table{
				topLeft:     tt.fields.topLeft,
				bottomRight: tt.fields.bottomRight,
			}
			if got := tr.nextCell(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("table.getCell() = %v, want %v", got, tt.want)
			}
		})
	}
}
