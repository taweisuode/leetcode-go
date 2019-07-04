package LeetCode

import (
	"fmt"
)

func Code300() {
	num := lengthOfLIS([]int{1, 3, 6, 7, 9, 4, 10, 5, 6})
	fmt.Println(num)
}

/**
给定一个无序的整数数组，找到其中最长上升子序列的长度。

示例:

输入: [10,9,2,5,3,7,101,18]
输出: 4
解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。
说明:

可能会有多种最长上升子序列的组合，你只需要输出对应的长度即可。
你算法的时间复杂度应该为 O(n2) 。
进阶: 你能将算法的时间复杂度降低到 O(n log n) 吗?

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-increasing-subsequence
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func lengthOfLIS(nums []int) int {

	//dp[i] 表示i长度数组的最长上升子序列的长度}

	/**
	dp[0] = 1
	9 < 10  dp[1] = 1 min = 9

	2 < 9  dp[2] = 1 min = 2

	5 > 2  dp[3] = dp[2] + 1 = 2

	3 > 2
		for j := 0; j < i; j++ {
			if num[i] > num[j] {
				dp[4] = dp[2] + 1 = 2
				dp[4] = max(dp[4],dp[2] + 1)
			}
	}
	*/
	dp := make([]int, len(nums))
	if len(nums) == 0 {
		return 0
	}

	for i := 0; i < len(nums); i++ {
		dp[i] = 1

		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	maxVal := dp[0]
	for _, item := range dp {
		maxVal = max(maxVal, item)
	}

	return maxVal
}
