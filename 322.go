package LeetCode

import (
	"fmt"
)

func Code322() {
	coins := []int{1, 2, 4, 5, 8, 3, 5}
	fmt.Println(coinChange(coins, 11))
}

/**
给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。

示例 1:

输入: coins = [1, 2, 5], amount = 11
输出: 3
解释: 11 = 5 + 5 + 1
示例 2:

输入: coins = [2], amount = 3
输出: -1
说明:
你可以认为每种硬币的数量是无限的。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/coin-change
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func coinChange(coins []int, amount int) int {
	coins = quick_sort2(coins)

	var count int
	dfs_322(coins, amount, 0, &count)
	return count
}

func dfs_322(coins []int, amount int, sum int, count *int) {
	if amount == 0 {
		if *count < sum {
			*count = sum
		}
	}
	for i := 0; i < len(coins); i++ {
		if amount < 0 {
			continue
		}
		dfs_322(coins, amount-coins[i], sum+coins[i], count)
	}

}
func quick_sort2(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}

	head, tail := 0, len(nums)-1
	key := nums[0]

	for i := 1; i <= tail; {
		if nums[i] < key {
			nums[i], nums[tail] = nums[tail], nums[i]
			tail--
		} else {
			nums[i], nums[head] = nums[head], nums[i]
			head++
			i++
		}
	}
	quick_sort2(nums[:head])
	quick_sort2(nums[head+1:])
	return nums
}
