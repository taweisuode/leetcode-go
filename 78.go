package LeetCode

import (
	"fmt"
)

func Code78() {
	//nums := []int{1, 2, 3}

	str := "abcd"
	fmt.Println(subsets2(str))
}

/**
给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。

说明：解集不能包含重复的子集。

示例:

输入: nums = [1,2,3]
输出:
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/subsets
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func subsets2(str string) []string {
	var res []string
	dfs_78_2(str, 0, "", &res)
	return res
}
func dfs_78_2(str string, index int, next string, res *[]string) {
	if len(next) >= 2 {
		*res = append(*res, next)
	}
	for i := index; i < len(str); i++ {
		//这里要用 i+ 1 用来防止重复使用
		next = str[:i]
		dfs_78_2(str, i+1, next, res)
		str = str[:i]
	}
}
func subsets(nums []int) [][]int {

	var res [][]int
	dfs_78(nums, 0, []int{}, &res)
	return res
}

func dfs_78(nums []int, index int, next []int, res *[][]int) {
	temp := make([]int, len(next))
	copy(temp, next)
	*res = append(*res, temp)
	for i := index; i < len(nums); i++ {
		next = append(next, nums[i])

		//这里要用 i+ 1 用来防止重复使用
		dfs_78(nums, i+1, next, res)
		next = next[:len(next)-1]
	}
}

//位运算
func subsets1(nums []int) [][]int {
	length := uint(len(nums))
	res := [][]int{}
	for i := 0; i < (1 << length); i++ {
		sub := []int{}
		for j := 0; j < len(nums); j++ {
			if (uint(i)>>uint(j))&1 == 1 {
				sub = append(sub, nums[j])
			}
		}
		res = append(res, sub)
	}
	return res
}
