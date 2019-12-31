package LeetCode

import "fmt"

func Code47() {
	nums := []int{1, 1, 2}
	fmt.Println(permuteUnique(nums))
}

/**
给定一个可包含重复数字的序列，返回所有不重复的全排列。

示例:

输入: [1,1,2]
输出:
[
  [1,1,2],
  [1,2,1],
  [2,1,1]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/permutations-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func permuteUnique(nums []int) [][]int {

	var res [][]int

	dfs_47(nums, 0, []int{}, &res)
	return res
}

func dfs_47(nums []int, start int, combine []int, res *[][]int) {
	if len(combine) == len(nums) {
		t := make([]int, len(combine))
		copy(t, combine)
		*res = append(*res, t)
		return
	}

	for i := start; i < len(nums); i++ {
		combine = append(combine, nums[i])
		fmt.Println(combine)
		dfs_47(nums, i, combine, res)
		combine = combine[:len(combine)-1]
	}
}

func notIn2(num int, combine []int) bool {
	for i := range combine {
		if combine[i] == num {
			return false
		}
	}
	return true

}
