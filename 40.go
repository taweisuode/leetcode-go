package LeetCode

import (
	"fmt"
	"sort"
)

func Code40() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	fmt.Println(combinationSum2(candidates, target))
}

/**
给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的每个数字在每个组合中只能使用一次。

说明：

所有数字（包括目标数）都是正整数。
解集不能包含重复的组合。
示例 1:

输入: candidates = [10,1,2,7,6,1,5], target = 8,
所求解集为:
[
  [1, 7],
  [1, 2, 5],
  [2, 6],
  [1, 1, 6]
]
示例 2:

输入: candidates = [2,5,2,1,2], target = 5,
所求解集为:
[
  [1,2,2],
  [5]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/combination-sum-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

```
*/
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var res [][]int
	dfs_40(candidates, 0, target, []int{}, &res)
	return res
}

func dfs_40(candidates []int, index int, target int, combine []int, res *[][]int) {
	if index > len(candidates) || target < 0 {
		return
	}
	if target == 0 {
		t := make([]int, len(combine))
		copy(t, combine)
		*res = append(*res, t)
		return
	}

	for i := index; i < len(candidates); i++ {
		if i > index && candidates[i-1] == candidates[i] {
			continue
		}

		combine = append(combine, candidates[i])
		dfs_40(candidates[1:], i, target-candidates[i], combine, res)
		combine = combine[:len(combine)-1]
	}
}

func combinationSum1(candidates []int, target int) [][]int {
	res := [][]int{}
	sort.Ints(candidates)
	helper(candidates, 0, 0, target, []int{}, &res)
	return res
}

func helper(candidates []int, start, sum, target int, out []int, res *[][]int) {
	if start > len(candidates) || sum > target {
		return
	}
	if sum == target {
		*res = append(*res, out)
		return
	}
	//if start>0&&candidates[start-1]==candidates[start]{return}
	for i := start; i < len(candidates); i++ {
		if i > start && candidates[i-1] == candidates[i] {
			continue
		}
		helper(candidates, i+1, sum+candidates[i], target, append([]int{candidates[i]}, out...), res)
	}

}
