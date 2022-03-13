package xo

// https://leetcode-cn.com/problems/kLl5u1/
/**
方法1. 遍历 numbers， 取出第 i 个，然后 对第 i+1 到 numbers.length -1 进行二分查找
*/
func twoSum(numbers []int, target int) []int {
	// 边界条件判断
	if len(numbers) < 2 {
		return []int{-1, -1}
	}
	for idx, v := range numbers {
		// 理论上不会到达最后一位
		if position := binarySearch(&numbers, idx+1, len(numbers)-1, target-v); position != -1 {
			return []int{idx, position}
		}
	}
	return []int{-1, -1}
}

/**
target - numbers[idx] = remainValue
position == -1 则说明没有找到需要向下一位查找
*/
func binarySearch(numbers *[]int, l, r, remainValue int) (position int) {

	// 左右指针
	for l <= r {
		/**
		mid的计算为什么用: int mid = (high - low) / 2 + low; 而不是： int mid = (hight + low) / 2;因为 int 的最大值为 MaxInt
		数组长度很大则容易越界
		*/
		mid := l + (r-l)/2
		midVal := (*numbers)[mid]
		if remainValue == midVal {
			return mid
		}
		if l == r {
			return -1
		}
		// 说明目标在左边
		if remainValue < midVal {
			return binarySearch(numbers, l, mid, remainValue)
		}
		return binarySearch(numbers, mid+1, r, remainValue)
	}
	return -1
}

func binarySearch2(numbers *[]int, l, r, remainValue int) (position int) {
	for l <= r {
		mid := l + (r-l)/2
		midVal := (*numbers)[mid]
		if remainValue == midVal {
			return mid
		}
		// 说明目标在左边, mid 为啥要加减 1，因为 mid 这个位置不是答案
		if remainValue < midVal {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}

/**
方法二. 双指针法，
1. 两个指针 l 和 r 分别从最左 和最右开始，
2. 如果 numbers[l]+numbers[r] < target ,说明左边的值太小需要l++, 如果 大于，则说明右边的值太大，需要往前推一位 直到算出 target
*/

func twoSum2(numbers []int, target int) []int {
	// 遍历 numbers， 取出第 i 个，然后 对第 i+1 到 i-1
	l, r := 0, len(numbers)-1
	for l < r {
		sum := numbers[l] + numbers[r]
		if sum == target {
			break
		}
		if sum > target {
			r--
			continue
		}
		if sum < target {
			l++
		}
	}
	return []int{l, r}
}
