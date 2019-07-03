package LeetCode

import (
	"fmt"
)

func Code213() {
	rob := rob2([]int{1})
	fmt.Println(rob)
}

/**
你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都围成一圈，这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。

给定一个代表每个房屋存放金额的非负整数数组，计算你在不触动警报装置的情况下，能够偷窃到的最高金额。

示例 1:

输入: [2,3,2]
输出: 3
解释: 你不能先偷窃 1 号房屋（金额 = 2），然后偷窃 3 号房屋（金额 = 2）, 因为他们是相邻的。
示例 2:

输入: [1,2,3,1]
输出: 4
解释: 你可以先偷窃 1 号房屋（金额 = 1），然后偷窃 3 号房屋（金额 = 3）。
     偷窃到的最高金额 = 1 + 3 = 4 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/house-robber-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func rob2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	return max(dpRob(nums[0:len(nums)-1]),dpRob(nums[1:]))
}

func dpRob(nums []int) int {
	//dp表示不相邻的最大值集合
	dp := make([]int,len(nums))
	/**
	[2,3,2]
	{
		dp[0] =>2
		3 >2  dp[1] = 3
		i + dp[i-2] &&
	}
	*/
	maxDp := 0
	if len(nums) <= 0 {
		return 0
	}else {
		if len(nums) == 1 {
			return nums[0]
		}
		if len(nums) >= 2 {
			dp[0] = nums[0]
			if nums[1] < dp[0] {
				dp[1] = dp[0]
			}else {
				dp[1] = nums[1]
			}

			//0,n-1
			for i:=2;i < len(nums);i++ {
				dp[i] = max(dp[i-2] + nums[i],dp[i-1])
			}
			for j:=0;j < len(dp);j ++ {
				maxDp = max(maxDp,dp[j])
			}
		}
	}
	return maxDp
}