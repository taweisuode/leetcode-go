package LeetCode

import (
	"fmt"
)

func Code494() {
	num := []int{1, 1, 1, 1, 1}
	fmt.Println(findTargetSumWays(num, 3))
}

/**
给定一个非负整数数组，a1, a2, ..., an, 和一个目标数，S。现在你有两个符号 + 和 -。对于数组中的任意一个整数，你都可以从 + 或 -中选择一个符号添加在前面。

返回可以使最终数组和为目标数 S 的所有添加符号的方法数。

示例 1:

输入: nums: [1, 1, 1, 1, 1], S: 3
输出: 5
解释:

-1+1+1+1+1 = 3
+1-1+1+1+1 = 3
+1+1-1+1+1 = 3
+1+1+1-1+1 = 3
+1+1+1+1-1 = 3

一共有5种方法让最终目标和为3。
注意:

数组非空，且长度不会超过20。
初始的数组的和不会超过1000。
保证返回的最终结果能被32位整数存下。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/target-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func findTargetSumWays(nums []int, S int) int {
	var count int
	if len(nums) == 0 {
		return 0
	}
	dfs_494(nums, 0, S, 0, &count)
	return count
}

func dfs_494(nums []int, start int, s int, combine int, count *int) {
	if start == len(nums) {
		if combine == s {
			*count++
		}
	} else {

		//二叉树 那就是递归2次  多叉树如39 ，40 需要循环遍历递归
		sum := combine + nums[start]
		dfs_494(nums, start+1, s, sum, count)
		del := combine - nums[start]
		dfs_494(nums, start+1, s, del, count)
	}
}
