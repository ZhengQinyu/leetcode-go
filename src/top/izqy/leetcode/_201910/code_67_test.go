package _201910

import "testing"

func Test_addBinary(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"1",
			args{
				a: "11",
				b: "1",
			},
			"100",
		},
		{
			"2",
			args{
				a: "1010",
				b: "1011",
			},
			"10101",
		},
		{
			"3",
			args{
				a: "101111",
				b: "10",
			},
			"110001",
		},
		{
			"4",
			args{
				a: "1111",
				b: "1111",
			},
			"11110",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addBinary(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("addBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
