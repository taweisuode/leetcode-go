package LeetCode

import "fmt"

func Code46() {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))
}

/**
给定一个没有重复数字的序列，返回其所有可能的全排列。

示例:

输入: [1,2,3]
输出:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]

```
*/
func permute(nums []int) [][]int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
		permute(nums[i+1:])
	}
	return [][]int{}
}
