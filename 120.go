package LeetCode

import (
	"fmt"
)

func Code120() {
	triangle := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}
	fmt.Println(minimumTotal(triangle))
}

/**
给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。

例如，给定三角形：

[
     [2],
    [3,4],
   [6,5,7],
  [4,1,8,3]
]
自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。

说明：

如果你可以只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题，那么你的算法会很加分。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/triangle
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func minimumTotal(triangle [][]int) int {

	dp := make([][]int, len(triangle))
	for i := range triangle {
		dp[i] = make([]int, len(triangle[i]))
	}

	for i := range triangle {
		for j := range triangle[i] {
			if i == 0 && j == 0 {
				dp[i][j] = triangle[0][0]
			} else if j == 0 && i != 0 {
				dp[i][j] = dp[i-1][0] + triangle[i][j]
			} else if i != 0 && j == i {
				dp[i][j] = dp[i-1][j-1] + triangle[i][j]
			} else {
				dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]
			}
		}
	}
	minRex := dp[len(dp)-1][0]
	for i := 0; i < len(dp); i++ {
		if dp[len(dp)-1][i] < minRex {
			minRex = dp[len(dp)-1][i]
		}
	}
	return minRex
}
