package _202104

import "testing"

func Test_isRectangleOverlap(t *testing.T) {
	type args struct {
		rec1 []int
		rec2 []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{[]int{0, 0, 2, 2}, []int{1, 1, 3, 3}}, true},
		{"2", args{[]int{0, 0, 1, 1}, []int{1, 0, 2, 1}}, false},
		{"3", args{[]int{0, 0, 1, 1}, []int{2, 2, 3, 3}}, false},
		{"4", args{[]int{-1, 0, 1, 1}, []int{0, -1, 0, 1}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isRectangleOverlap(tt.args.rec1, tt.args.rec2); got != tt.want {
				t.Errorf("isRectangleOverlap() = %v, want %v", got, tt.want)
			}
		})
	}
}
