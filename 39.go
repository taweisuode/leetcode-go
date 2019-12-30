package LeetCode

import (
	"fmt"
	"sort"
)

func Code39() {
	candidates := []int{2, 3, 5}
	target := 8
	fmt.Println(combinationSum(candidates, target))
}

/**
给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的数字可以无限制重复被选取。

说明：

所有数字（包括 target）都是正整数。
解集不能包含重复的组合。
示例 1:

输入: candidates = [2,3,6,7], target = 7,
所求解集为:
[
  [7],
  [2,2,3]
]
示例 2:

输入: candidates = [2,3,5], target = 8,
所求解集为:
[
  [2,2,2,2],
  [2,3,3],
  [3,5]
]

```
*/
func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := make([][]int, 0)

	dfs_39(candidates, target, 0, []int{}, &res)
	return res
}

func dfs_39(candidates []int, target int, index int, combine []int, res *[][]int) {
	if target < 0 {
		return
	}
	if target == 0 {
		fmt.Println(combine)
		t := make([]int, len(combine))
		copy(t, combine)
		*res = append(*res, t)
		return
	}
	for i := index; i < len(candidates); i++ {
		combine = append(combine, candidates[i])
		dfs_39(candidates, target-candidates[i], i, combine, res)
		combine = combine[:len(combine)-1]
	}

}
