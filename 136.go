package LeetCode

import "fmt"

func Code136() {
	nums := []int{2, 2, 1}
	fmt.Println(singleNumberV2(nums))
}

/**
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

说明：

你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？

示例 1:

输入: [2,2,1]
输出: 1
示例 2:

输入: [4,1,2,1,2]
输出: 4
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func singleNumber(nums []int) int {
	hashMap := make(map[int]int, 0)
	for _, item := range nums {
		hashMap[item]++
	}
	for key, val := range hashMap {
		if val == 1 {
			return key
		}
	}
	return 0
}

//异或操作  异或操作  a ^a = 0   a ^ 0 = a  相同为0 并且满足交换律
func singleNumberV2(nums []int) int {
	res := 0
	for _, item := range nums {
		res ^= item
	}
	return res
}
