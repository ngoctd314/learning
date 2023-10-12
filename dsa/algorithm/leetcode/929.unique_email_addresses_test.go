package leetcode

import "testing"

func Test_numUniqueEmails(t *testing.T) {
	type args struct {
		emails []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				emails: []string{"alice.z@leetcode.com", "alicez@leetcode.com"},
			},
			want: 1,
		},
		{
			name: "Test 2",
			args: args{
				emails: []string{"test.email+alex@leetcode.com", "test.email.leet+alex@code.com"},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numUniqueEmails(tt.args.emails); got != tt.want {
				t.Errorf("numUniqueEmails() = %v, want %v", got, tt.want)
			}
		})
	}
}
