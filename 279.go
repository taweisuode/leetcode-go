package LeetCode

import (
	"fmt"
	"math"
)

func Code279() {
	num := numSquares(13)
	fmt.Println(num)
}

/**
给定正整数 n，找到若干个完全平方数（比如 1, 4, 9, 16, ...）使得它们的和等于 n。你需要让组成和的完全平方数的个数最少。

示例 1:

输入: n = 12
输出: 3
解释: 12 = 4 + 4 + 4.
示例 2:

输入: n = 13
输出: 2
解释: 13 = 4 + 9.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/perfect-squares
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func numSquares(n int) int {
	//dp[i] 表示i能至少有多个完全平方数组成

	/**
	 	dp[12] = min(dp[12],dp[12-1*1]+1)
		dp[12] = min(dp[12],dp[12-2*2]+1)
		dp[12] = min(dp[12],dp[12-3*3]+1)
	*/
	dp := make([]int, n+1)
	for init := 0; init < n+1; init++ {
		dp[init] = math.MaxInt32
	}
	dp[0] = 0
	for i := 0; i <= n; i++ {
		for j := 1; j*j <= i; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}
	fmt.Println(dp)
	return dp[n]
}
