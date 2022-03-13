package xo

import "testing"

func Test_minSubArrayLen(t *testing.T) {
	type args struct {
		target int
		nums   []int
	}
	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		// Add test cases.
		{
			name: "case1",
			args: args{
				target: 11,
				nums:   []int{1, 2, 3, 4, 5},
			},
			wantRes: 3,
		},
		{
			name: "case2",
			args: args{
				target: 7,
				nums:   []int{2, 3, 1, 2, 4, 3},
			},
			wantRes: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := minSubArrayLen(tt.args.target, tt.args.nums); gotRes != tt.wantRes {
				t.Errorf("minSubArrayLen() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
