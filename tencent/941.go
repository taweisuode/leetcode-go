package tencent

import (
	"fmt"
	"math"
)

func Code941() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majorityElement(nums))
}

/**

给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。

你可以假设数组是非空的，并且给定的数组总是存在多数元素。

示例 1:

输入: [3,2,3]
输出: 3
示例 2:

输入: [2,2,1,1,1,2,2]
输出: 2

*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func majorityElement(nums []int) int {
	resMap := make(map[int]int)
	for i := range nums {
		if _, ok := resMap[nums[i]]; ok {
			resMap[nums[i]]++
		} else {
			resMap[nums[i]] = 1
		}
	}
	max := math.MinInt32
	res := 0
	for i, item := range resMap {
		if item > max {
			max = item
			res = i
		}
	}
	return res
}
