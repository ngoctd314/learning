package leetcode

import (
	"reflect"
	"testing"
)

func Test_suggestedProducts(t *testing.T) {
	type args struct {
		products   []string
		searchWord string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "Test 1",
			args: args{
				products:   []string{"mobile", "mouse", "moneypot", "monitor", "mousepad"},
				searchWord: "mouse",
			},
			want: [][]string{
				{"mobile", "moneypot", "monitor"}, {"mobile", "moneypot", "monitor"}, {"mouse", "mousepad"}, {"mouse", "mousepad"}, {"mouse", "mousepad"},
			},
		},
		{
			name: "Test 2",
			args: args{
				products:   []string{"bags", "baggage", "banner", "box", "cloths"},
				searchWord: "bags",
			},
			want: [][]string{
				{"baggage", "bags", "banner"}, {"baggage", "bags", "banner"}, {"baggage", "bags"}, {"bags"},
			},
		},
		{
			name: "Test 3",
			args: args{
				products:   []string{"havana"},
				searchWord: "tatiana",
			},
			want: [][]string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := suggestedProducts(tt.args.products, tt.args.searchWord); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("suggestedProducts() = %v, want %v", got, tt.want)
			}
		})
	}
}
