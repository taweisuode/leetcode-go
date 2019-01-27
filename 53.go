package LeetCode

import "fmt"

func Code53() {
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(arr))
}

/**
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

示例:

输入: [-2,1,-3,4,-1,2,1,-5,4],
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
进阶:

如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func maxSubArray(nums []int) int {
	dp := []int{0}
	dp[0] = nums[0]
	for key, _ := range nums {
		if key > 0 {
			dp = append(dp, max(nums[key], dp[key-1]+nums[key]))
		}
	}
	temp := dp[0]

	for k, _ := range dp {
		if k > 0 {
			temp = max(dp[k], temp)
		}
	}
	return temp
}
func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
