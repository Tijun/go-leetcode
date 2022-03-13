package leetcode

// https://leetcode-cn.com/problems/binary-search/

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (l+r)/2
		if nums[mid] == target {
			return mid
		}
		// 说明结果在左边
		if target < nums[mid] {
			r = mid - 1
			continue
		}
		if target > nums[mid] {
			l = mid + 1
		}
	}
	return -1
}
