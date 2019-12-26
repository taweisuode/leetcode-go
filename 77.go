package LeetCode

import (
	"fmt"
)

func Code77() {
	fmt.Println(combine(4, 2))
}

/**
给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。

示例:

输入: n = 4, k = 2
输出:
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/combinations
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func combine(n int, k int) [][]int {
	var res [][]int
	dfs_77(n, k, 1, []int{}, &res)
	return res
}
func dfs_77(n int, k int, startN int, combine []int, res *[][]int) {
	if len(combine) == k {
		*res = append(*res, combine)
	}
	for i := startN; i < n+1; i++ {
		dfs_77(n, k, i+1, append([]int{i}, combine...), res)
	}
}
