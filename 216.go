package LeetCode

import (
	"fmt"
)

func Code216() {
	fmt.Println(combinationSum3(3, 15))
}

/**
找出所有相加之和为 n 的 k 个数的组合。组合中只允许含有 1 - 9 的正整数，并且每种组合中不存在重复的数字。

说明：

所有数字都是正整数。
解集不能包含重复的组合。
示例 1:

输入: k = 3, n = 7
输出: [[1,2,4]]
示例 2:

输入: k = 3, n = 9
输出: [[1,2,6], [1,3,5], [2,3,4]]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/combination-sum-iii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func combinationSum3(k int, n int) [][]int {
	var res [][]int

	dfs_216(k, n, 1, []int{}, &res)
	return res
}

func dfs_216(k int, n int, start int, combine []int, res *[][]int) {
	if n == 0 && len(combine) == k {
		t := make([]int, len(combine))
		copy(t, combine)
		*res = append(*res, t)
	}
	for i := start; i < 10; i++ {
		combine = append(combine, i)
		dfs_216(k, n-i, i+1, combine, res)
		combine = combine[:len(combine)-1]
	}
}
