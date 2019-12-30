package LeetCode

import (
	"fmt"
	"sort"
)

func Code90() {
	nums := []int{1, 2, 2}
	num := subsetsWithDup(nums)
	fmt.Println(num)
}

/**
给定一个可能包含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。

说明：解集不能包含重复的子集。

示例:

输入: [1,2,2]
输出:
[
  [2],
  [1],
  [1,2,2],
  [2,2],
  [1,2],
  []
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/subsets-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func subsetsWithDup(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	dfs_90(nums, 0, []int{}, &res)
	return res
}

func dfs_90(nums []int, start int, combine []int, res *[][]int) {
	t := make([]int, len(combine))
	copy(t, combine)
	*res = append(*res, t)
	for i := start; i < len(nums); i++ {
		if i-1 >= start && nums[i-1] == nums[i] {
			continue
		}
		combine = append(combine, nums[i])
		dfs_90(nums[1:], i, combine, res)
		combine = combine[:len(combine)-1]
	}
}
