package LeetCode

import (
	"fmt"
	"math"
	"sort"
)

func Code5292() {
	nums := []int{3, 2, 1, 2, 3, 4, 3, 4, 5, 9, 10, 11}
	fmt.Println(isPossibleDivide(nums, 3))
}

/**
给你一个整数数组 nums 和一个正整数 k，请你判断是否可以把这个数组划分成一些由 k 个连续数字组成的集合。
如果可以，请返回 True；否则，返回 False。



示例 1：

输入：nums = [1,2,3,3,4,4,5,6], k = 4
输出：true
解释：数组可以分成 [1,2,3,4] 和 [3,4,5,6]。
示例 2：

输入：nums = [3,2,1,2,3,4,3,4,5,9,10,11], k = 3
输出：true
解释：数组可以分成 [1,2,3] , [2,3,4] , [3,4,5] 和 [9,10,11]。
示例 3：

输入：nums = [3,3,2,2,1,1], k = 3
输出：true
示例 4：

输入：nums = [1,2,3,4], k = 3
输出：false
解释：数组不能分成几个大小为 3 的子数组。


提示：

1 <= nums.length <= 10^5
1 <= nums[i] <= 10^9
1 <= k <= nums.length
*/
func isPossibleDivide(nums []int, k int) bool {
	if len(nums) == 0 || len(nums)%k != 0 {
		return false
	}
	sort.Ints(nums)

	numMap := make(map[int]int, len(nums))
	for i := range nums {
		if _, ok := numMap[nums[i]]; ok {
			numMap[nums[i]]++
		} else {
			numMap[nums[i]] = 1
		}
	}

	for i := range numMap {
		fmt.Println(i)
		first := getFirstNode(numMap)
		count := 0
		for count < k {
			if _, ok := numMap[first+count]; ok {
				if numMap[first+count] == 1 {
					delete(numMap, first+count)
				} else {
					numMap[first+count]--
				}
			} else {
				return false
			}
			count++
		}
	}
	return true
}

func getFirstNode(numMap map[int]int) int {
	if len(numMap) == 0 {
		return math.MinInt32
	}
	min := math.MaxInt32
	for i := range numMap {
		if i < min {
			min = i
		}
	}
	return min
}
