package xo

/**
最基础的 case 是
target = nums[i]
targetsN = targets[N-1] +1 or targets[N-2] + 2

*/

// https://leetcode-cn.com/problems/2VG8Kg/
func minSubArrayLen(target int, nums []int) int {
	// 定义 targets结果集
	targetsMap := make(map[int]int, len(nums))
	for _, v := range nums {
		targetsMap[v] = 1
	}
	if targetsMap[target] != 0 {
		return targetsMap[target]
	}
	for i := 1; i <= target; i++ {
		if targetsMap[i] != 0 {
			//return
		}
	}
	// 初始化
	return 0
}
