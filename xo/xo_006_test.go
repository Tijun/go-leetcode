package xo

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		numbers []int
		target  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// test cases.
		{
			name: "case1",
			args: args{
				numbers: []int{1, 2, 4, 6, 10},
				target:  8,
			},
			want: []int{1, 3},
		},
		{
			name: "case2",
			args: args{
				numbers: []int{2, 3, 4},
				target:  6,
			},
			want: []int{0, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.numbers, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
