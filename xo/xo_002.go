package xo

import (
	"strconv"
	"strings"
)

// https://leetcode-cn.com/problems/JFETK5/
func addBinary(a string, b string) (result string) {
	var aIndex, bIndex = len(a) - 1, len(b) - 1
	// 用来保存进位
	var carry = 0
	for aIndex >= 0 || bIndex >= 0 {
		// val a,b 分别是同等位置的两个数，如果某一方没有了，那么就为 0，对结果不会有影响
		var valA, valB = 0, 0
		if aIndex >= 0 {
			valA = int(a[aIndex] - '0')
			aIndex--
		}
		if bIndex >= 0 {
			valB = int(b[bIndex] - '0')
			bIndex--
		}
		sum := carry + valA + valB
		// 需要处理进位
		if sum >= 2 {
			carry = 1
			result = strconv.Itoa(sum-2) + result
		} else {
			result = strconv.Itoa(sum) + result
			carry = 0
		}
	}
	if carry > 0 {
		result = strconv.Itoa(carry) + result
	}
	if strings.EqualFold(result, "") {
		return "0"
	}
	return result
}
