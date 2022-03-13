package xo

import (
	"fmt"
	"sort"
)

// https://leetcode-cn.com/problems/1fGaJU/

/**
1. a + b + c = 0 -> a = (-b + -c)
将 three sum 转换为 two sum
*/
func threeSum(nums []int) (result [][]int) {
	if len(nums) < 3 {
		return result
	}
	// 排序，方便后边的 tosum 元素进行二分查找
	sort.Ints(nums)
	// 去重
	keys := make(map[string]bool)
	for i := 0; i < len(nums)-2; i++ {
		// 将 three sum 转换为 tow sum 问题
		target := 0 - nums[i]
		if arr := toSum(&nums, i+1, target); len(arr) != 0 {
			for _, res := range arr {
				key := fmt.Sprintf("%d%d%d", nums[i], nums[res[0]], nums[res[1]])
				if keys[key] {
					continue
				}
				keys[key] = true
				result = append(result, []int{nums[i], nums[res[0]], nums[res[1]]})
			}
		}
	}
	return result
}

/**
two sum 问题
*/
func toSum(nums *[]int, idx, target int) (res [][]int) {
	// b+c = target
	for i := idx; i < len(*nums); i++ {
		remain := target - (*nums)[i]
		if r := searchTarget(nums, i+1, remain); r != -1 {
			res = append(res, []int{i, r})
		}
	}
	return res
}

// 返回目标值位置
func searchTarget(nums *[]int, lowerPointer, target int) int {
	l, r := lowerPointer, len(*nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if target == (*nums)[mid] {
			return mid
		}
		if target < (*nums)[mid] {
			r = mid - 1
			continue
		}
		l = mid + 1
	}
	return -1
}
