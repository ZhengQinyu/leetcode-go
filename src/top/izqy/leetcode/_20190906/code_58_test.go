package leetcode

import "testing"

func Test_lengthOfLastWord(t *testing.T) {
	tests := []struct {
		name string
		args string
		want int
	}{{
		name: "test_58",
		args: "Hello World",
		want: 5,
	}, {
		name: "test_58",
		args: " Hello11 ",
		want: 7,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLastWord(tt.args); got != tt.want {
				t.Errorf("lengthOfLastWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
