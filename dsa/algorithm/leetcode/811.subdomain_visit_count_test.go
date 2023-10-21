package leetcode

import (
	"reflect"
	"testing"
)

func Test_subdomainVisits(t *testing.T) {
	type args struct {
		cpdomains []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test 1",
			args: args{
				cpdomains: []string{"9001 discuss.leetcode.com"},
			},
			want: []string{"9001 leetcode.com", "9001 discuss.leetcode.com", "9001 com"},
		},
		{
			name: "Test 2",
			args: args{
				cpdomains: []string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"},
			},
			want: []string{"901 mail.com", "50 yahoo.com", "900 google.mail.com", "5 wiki.org", "5 org", "1 intel.mail.com", "951 com"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subdomainVisits(tt.args.cpdomains); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subdomainVisits() = %v, want %v", got, tt.want)
			}
		})
	}
}
