package _201910

import "testing"

func Test_mySqrt(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want int
	}{
		{"1", 4, 2},
		{"2", 8, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mySqrt(tt.x); got != tt.want {
				t.Errorf("mySqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}
