package bytedance

import "fmt"

func Code1029() {
	arr := []int{-2, -1}
	fmt.Println(maxSubArray(arr))

}

/**

给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

示例:

输入: [-2,1,-3,4,-1,2,1,-5,4],
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
进阶:


```
*/

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	//dp[i]表示 以i为结尾的最大数组值的和
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = Max(dp[i-1]+nums[i], nums[i])
	}
	maxRes := dp[0]
	for _, item := range dp {
		if item > maxRes {
			maxRes = item
		}
	}

	return maxRes
}
