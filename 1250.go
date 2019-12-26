package LeetCode

import (
	"fmt"
)

func Code1250() {
	arr := []int{6, 10, 15}
	res := isGoodArray(arr)
	fmt.Println(res)
}

/**
给你一个正整数数组 nums，你需要从中任选一些子集，然后将子集中每一个数乘以一个 任意整数，并求出他们的和。

假如该和结果为 1，那么原数组就是一个「好数组」，则返回 True；否则请返回 False。



示例 1：

输入：nums = [12,5,7,23]
输出：true
解释：挑选数字 5 和 7。
5*3 + 7*(-2) = 1
示例 2：

输入：nums = [29,6,10]
输出：true
解释：挑选数字 29, 6 和 10。
29*1 + 6*(-3) + 10*(-1) = 1
示例 3：

输入：nums = [3,6]
输出：false


提示：

1 <= nums.length <= 10^5
1 <= nums[i] <= 10^9

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/check-if-it-is-a-good-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func isGoodArray(nums []int) bool {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		res = gcd2(res, nums[i])
	}
	return res == 1
}

//辗转相除法
func gcd1(a1 int, a2 int) int {
	if a1%a2 == 0 {
		return a2
	} else {
		return gcd1(a2, a1%a2)
	}
}

//更相减损法
func gcd2(a1 int, a2 int) int {
	if a1 == a2 {
		return a1
	}
	if a1 > a2 {
		return gcd2(a1-a2, a2)
	} else {
		return gcd2(a1, a2-a1)
	}
}

//数学证明  ax + by = 1  只有在 a,b 互质的时候才有整数解
/*
	反证法：
		假设a,b 不互质，则 必有  a = p * gcd(a,b)   b = q * gcd(a,b)
		带入公式

		px + qy = 1/gcb(a,b)
		如果ab 不互质  则 右边已经是小数了 所以 ab 不互质 不成立

		所以ab 互质的时候  ax+ by = 1 有整数解

*/
