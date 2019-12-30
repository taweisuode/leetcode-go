package LeetCode

import (
	"fmt"
)

func Code491() {
	num := []int{4, 6, 7, 7}
	fmt.Println(findSubsequences(num))
}

/**
给定一个整型数组, 你的任务是找到所有该数组的递增子序列，递增子序列的长度至少是2。

示例:

输入: [4, 6, 7, 7]
输出: [[4, 6], [4, 7], [4, 6, 7], [4, 6, 7, 7], [6, 7], [6, 7, 7], [7,7], [4,7,7]]
说明:

给定数组的长度不会超过15。
数组中的整数范围是 [-100,100]。
给定数组中可能包含重复数字，相等的数字应该被视为递增的一种情况。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/increasing-subsequences
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func findSubsequences(nums []int) [][]int {
	var res [][]int
	dfs_491(nums, 0, []int{}, &res)
	return res
}

func dfs_491(nums []int, start int, combine []int, res *[][]int) {
	if len(combine) >= 2 {
		t := make([]int, len(combine))
		copy(t, combine)
		*res = append(*res, t)

	}
	numMap := make(map[int]bool)
	for i := start; i < len(nums); i++ {
		if numMap[nums[i]] == true {
			continue
		}
		numMap[nums[i]] = true
		if len(combine) == 0 || combine[len(combine)-1] <= nums[i] {
			//fmt.Println(combine)
			combine = append(combine, nums[i])
			dfs_491(nums, i+1, combine, res)
			combine = combine[:len(combine)-1]
		}

	}
}
