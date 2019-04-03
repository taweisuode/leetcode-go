package LeetCode

import "fmt"

func Code34() {
	num := []int{1}
	target := 1
	fmt.Println(searchRange(num, target))
}

/**
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。

你的算法时间复杂度必须是 O(log n) 级别。

如果数组中不存在目标值，返回 [-1, -1]。

示例 1:

输入: nums = [5,7,7,8,8,10], target = 8
输出: [3,4]
示例 2:

输入: nums = [5,7,7,8,8,10], target = 6
输出: [-1,-1]

```
*/
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	left := left_bound(nums, target)
	right := right_bound(nums, target)
	//fmt.Println(left, right)
	return []int{left, right}
}
func left_bound(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	mid := (right + left) / 2
	for left <= right {
		mid = (right + left) / 2
		if nums[mid] == target {
			if mid == 0 || nums[mid-1] < target {
				return mid
			}
			right = mid - 1
		}
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	return -1
}

func right_bound(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	mid := (right + left) / 2

	for left <= right {
		mid = (right + left) / 2
		fmt.Println(mid)
		if nums[mid] == target {
			fmt.Println(mid, nums[mid])
			if mid == len(nums)-1 || nums[mid+1] > target {
				return mid
			}
			left = mid + 1
		}
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	return -1
}
