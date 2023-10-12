package leetcode

import "testing"

func Test_checkIfPangram(t *testing.T) {
	type args struct {
		sentence string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test 1",
			args: args{
				sentence: "thequickbrownfoxjumpsoverthelazydog",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkIfPangram(tt.args.sentence); got != tt.want {
				t.Errorf("checkIfPangram() = %v, want %v", got, tt.want)
			}
		})
	}
}
