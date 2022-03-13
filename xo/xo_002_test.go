package xo

import (
	"fmt"
	"testing"
)

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
		// test cases
		{
			name: "base1",
			args: args{
				a: "100",
				b: "110010",
			},
			want: "110110",
		},
		{
			name: "case1",
			args: args{
				a: "11",
				b: "10",
			},
			want: "101",
		},
		{
			name: "case2",
			args: args{
				a: "1010",
				b: "1011",
			},
			want: "10101",
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
func Test_addBinarya(t *testing.T) {
	var a = "1"
	i := int(a[0])
	fmt.Print(i)
	fmt.Print(int('0'))
}
