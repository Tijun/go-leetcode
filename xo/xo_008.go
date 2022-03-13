package xo

// https://leetcode-cn.com/problems/2VG8Kg/
/**
使用滑动窗口
1. 定义窗口大小，0~len(nums)
初始化左边界 left = 0
初始化返回值 ret = 最小值 or 最大值
for 右边界 in 可迭代对象:
	更新窗口内部信息
	while 根据题意进行调整：
		比较并更新ret(收缩场景时)
		扩张或收缩窗口大小
	比较并更新ret(扩张场景时)
返回 ret

作者：qingfengpython
链接：https://leetcode-cn.com/problems/2VG8Kg/solution/shua-chuan-jian-zhi-offer-day06-shu-zu-i-d5ne/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func minSubArrayLen(target int, nums []int) (res int) {
	if len(nums) == 0 {
		return 0
	}
	// window size 从 0 ~ len(num), result = windowSize + 1, 为零的时候其实是1
	windowVal := 0
	for windowSize := 1; windowSize <= len(nums); windowSize++ {
		l, r := 0, windowSize-1
		windowVal = initWindowVal(&nums, windowSize)
		if windowVal >= target {
			return windowSize
		}
		r++
		l++
		for r < len(nums) {
			// l-1 为滑动窗口中被抛弃掉的元素
			windowVal = windowVal - nums[l-1] + nums[r]
			if windowVal >= target {
				return windowSize
			}
			r++
			l++
		}
	}
	// 初始化
	return 0
}

func initWindowVal(nums *[]int, windowSize int) (res int) {
	for i := 0; i < windowSize; i++ {
		res += (*nums)[i]
	}
	return res
}
