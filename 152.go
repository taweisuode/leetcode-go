package LeetCode

import (
	"fmt"
	"math"
)

func Code152() {
	nums := []int{2, 3, -2, 4}
	fmt.Println(maxProduct(nums))
}

/**
给定一个整数数组 nums ，找出一个序列中乘积最大的连续子序列（该序列至少包含一个数）。

示例 1:

输入: [2,3,-2,4]
输出: 6
解释: 子数组 [2,3] 有最大乘积 6。
示例 2:

输入: [-2,0,-1]
输出: 0
解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-product-subarray
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func maxProduct(nums []int) int {
	iimax := math.MinInt64
	imax, imin := 1, 1
	for i := range nums {
		if nums[i] < 0 {
			imax, imin = imin, imax
		}
		imax = max(imax*nums[i], nums[i])
		imin = min(imin*nums[i], nums[i])
		iimax = max(iimax, imax)
	}
	return iimax
}
