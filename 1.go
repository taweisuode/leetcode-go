package LeetCode

import (
	"fmt"
)

func Code1() {
	fmt.Println("leetcode 1 ")
	nums := []int{2, 3, 4, 6, 8}
	target := 8
	fmt.Println(twoSum(nums, target))
}

/**
给定一个整数数组和一个目标值，找出数组中和为目标值的两个数。
你可以假设每个输入只对应一种答案，且同样的元素不能被重复利用。

```
	给定 nums = [2, 7, 11, 15], target = 9
	因为 nums[0] + nums[1] = 2 + 7 = 9
	所以返回 [0, 1]
```
*/
func twoSum(nums []int, target int) []int {
	fmt.Println(111)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == target-nums[j] {
				return []int{i, j}
			}
		}
	}
	return nil
}
