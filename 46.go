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

	var res [][]int

	dfs_46(nums, []int{}, &res)
	return res
}

func dfs_46(nums []int, combine []int, res *[][]int) {
	if len(combine) == len(nums) {
		t := make([]int, len(combine))
		copy(t, combine)
		*res = append(*res, t)
		return
	}

	for i := 0; i < len(nums); i++ {
		if notIn(nums[i], combine) {
			combine = append(combine, nums[i])
			dfs_46(nums, combine, res)
			combine = combine[:len(combine)-1]
		}
	}
}

func notIn(num int, combine []int) bool {
	for i := range combine {
		if combine[i] == num {
			return false
		}
	}
	return true

}
