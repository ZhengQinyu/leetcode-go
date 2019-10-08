package _201910

import (
	"reflect"
	"testing"
)

func Test_plusOne(t *testing.T) {
	type args struct {
		digits []int
	}
	tests := []struct {
		name   string
		digits []int
		want   []int
	}{
		{"1", []int{1, 2, 3}, []int{1, 2, 4}},
		{"2", []int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
		{"3", []int{1, 1, 9}, []int{1, 2, 0}},
		{"4", []int{1, 9, 9}, []int{2, 0, 0}},
		{"4", []int{9, 9, 9}, []int{1, 0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := plusOne(tt.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("plusOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
