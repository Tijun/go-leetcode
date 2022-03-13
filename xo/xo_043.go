package xo

// https://leetcode-cn.com/problems/1nzheng-shu-zhong-1chu-xian-de-ci-shu-lcof/
func countDigitOne(n int) int {
	// 第一位初始为 1
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	var total = 1
	for i := 2; i <= n; i++ {
		total = total + i&(i-1)
	}
	return total
}
