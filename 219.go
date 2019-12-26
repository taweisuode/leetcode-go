package LeetCode

import (
	"fmt"
)

func Code219() {
	s := []int{2147483647, -2147483648, 2147483647, -2147483648}
	t := 2
	num := containsNearbyDuplicate(s, t)
	fmt.Println(num)
}

/**
给定一个整数数组和一个整数 k，判断数组中是否存在两个不同的索引 i 和 j，使得 nums [i] = nums [j]，并且 i 和 j 的差的绝对值最大为 k。

示例 1:

输入: nums = [1,2,3,1], k = 3
输出: true
示例 2:

输入: nums = [1,0,1,1], k = 1
输出: true
示例 3:

输入: nums = [1,2,3,1,2,3], k = 2
输出: false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/contains-duplicate-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func containsNearbyDuplicate(nums []int, k int) bool {
	if len(nums) == 0 {
		return false
	}
	hashSet := make(map[int]int, 0)

	for key, val := range nums {
		if item, ok := hashSet[val]; ok {
			if key-item < item || item == 0 {
				hashSet[val] = key - item
			}
		} else {
			hashSet[val] = 0
		}
	}
	count := 0
	for _, item := range hashSet {
		if item > k {
			return false
		}
		if item == 0 {
			count++
		}
	}
	if count == len(hashSet) {
		return false
	}
	return true
}
