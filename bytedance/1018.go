package bytedance

import (
	"fmt"
)

func Code1018() {
	arr := []int{3, 2, 1, 5, 6, 4}
	key := 2
	fmt.Println(findKthLargest(arr, key))

}

/**
在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

示例 1:

输入: [3,2,1,5,6,4] 和 k = 2
输出: 5
示例 2:

输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
输出: 4
说明:

你可以假设 k 总是有效的，且 1 ≤ k ≤ 数组的长度。


```
*/

func findKthLargest(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	sortArr := quick_rsort(nums)

	return sortArr[k-1]

}

func quick_rsort(nums []int) []int {
	/*if len(nums) <= 0 {
		return nums
	}
	head, tail := 0, len(nums)-1
	key := nums[0]
	for i := 1; i <= tail; {
		if nums[i] < key {
			nums[i], nums[tail] = nums[tail], nums[i]
			tail--
		} else {
			nums[i], nums[head] = nums[head], nums[i]
			head++
			i++
		}
	}
	quick_rsort(nums[:head])
	quick_rsort(nums[head+1:])
	return nums*/

	if len(nums) <= 0 {
		return nums
	}

	head := 0
	tail := len(nums) - 1
	key := nums[0]
	for i := 1; i <= tail; {
		if nums[i] < key {
			nums[i], nums[tail] = nums[tail], nums[i]
			tail--
		} else {
			nums[i], nums[head] = nums[head], nums[i]
			i++
			head++
		}
	}
	quick_rsort(nums[0:head])
	quick_rsort(nums[head+1:])
	return nums

}
