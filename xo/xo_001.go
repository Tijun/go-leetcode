package xo

//https://leetcode-cn.com/problems/xoh6Oh/

import "math"

func divide(a int, b int) int {
	if a == 0 {
		return 0
	}
	if b == 1 {
		return a
	}
	/**
	case 1 最大值
	防止溢出
	MaxInt32  = 1<<31 - 1
	MinInt32  = -1 << 31
	*/
	if a == math.MinInt32 && b == -1 {
		return math.MaxInt32
	}

	if b == -1 {
		return -a
	}

	// 转换为负数处理
	// 1. 判断符号
	sign := 1
	if a < 0 && b > 0 || a > 0 && b < 0 {
		sign = -1
	}
	// 将正数转换为负数
	if a > 0 {
		a = -a
	}
	if b > 0 {
		b = -b
	}
	if a > b {
		return 0
	}

	return dividedCore(a, b) * sign
}
func dividedCore(a int, b int) (result int) {
	// 一轮次一轮次的瓜分
	for a <= b {
		var count = 1
		var v = b
		// v 大于最小值才不会溢出
		for v >= math.MinInt32 && a < (v<<1) {
			count <<= 1
			v <<= 1
		}
		result += count
		a -= v
	}
	return result
}
