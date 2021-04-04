package common

import "testing"

func TestGcd(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{2, 4}, 2},
		{"2", args{3, 9}, 3},
		{"3", args{8, 9}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Gcd(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("Gcd() = %v, want %v", got, tt.want)
			}
		})
	}
}
