package xo

// 将最右边的1变为0，111 & 110，n=n&(n-1)

// https://leetcode-cn.com/problems/er-jin-zhi-zhong-1de-ge-shu-lcof/
func hammingWeight(num uint32) int {
	var count = 0
	var j = num
	for j != 0 {
		count++
		j = j & (j - 1)
	}
	return count
}
