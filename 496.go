package LeetCode

import (
	"fmt"
)

func Code496() {
	num1 := []int{2, 4}
	num2 := []int{1, 3, 2, 4}
	num := nextGreaterElement(num1, num2)
	fmt.Println(num)
}

/**
给定两个没有重复元素的数组 nums1 和 nums2 ，其中nums1 是 nums2 的子集。找到 nums1 中每个元素在 nums2 中的下一个比其大的值。

nums1 中数字 x 的下一个更大元素是指 x 在 nums2 中对应位置的右边的第一个比 x 大的元素。如果不存在，对应位置输出-1。

示例 1:

输入: nums1 = [4,1,2], nums2 = [1,3,4,2].
输出: [-1,3,-1]
解释:
    对于num1中的数字4，你无法在第二个数组中找到下一个更大的数字，因此输出 -1。
    对于num1中的数字1，第二个数组中数字1右边的下一个较大数字是 3。
    对于num1中的数字2，第二个数组中没有下一个更大的数字，因此输出 -1。
示例 2:

输入: nums1 = [2,4], nums2 = [1,2,3,4].
输出: [3,-1]
解释:
    对于num1中的数字2，第二个数组中的下一个较大数字是3。
    对于num1中的数字4，第二个数组中没有下一个更大的数字，因此输出 -1。
注意:

nums1和nums2中所有元素是唯一的。
nums1和nums2 的数组大小都不超过1000。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/next-greater-element-i
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func nextGreaterElement1(nums1 []int, nums2 []int) []int {
	var res []int
	for j := 0; j < len(nums1); j++ {
		for i := 0; i < len(nums2); i++ {
			if nums2[i] == nums1[j] {
				fmt.Println(nums1[j])
				res = append(res, getMaxInt(nums2[i], i, nums2))
			}
		}
	}
	return res
}
func getMaxInt(num int, index int, numList []int) int {
	for i := index + 1; i < len(numList); i++ {
		if numList[i] > num {
			return numList[i]
		}
	}
	return -1
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	stack := &Stack{}

	stackMap := make(map[int]int)
	for i := range nums2 {
		if !StackEmpty(stack) {
			top := StackTop(stack).(int)
			if top < nums2[i] {
				Stackpop(stack)
				stackMap[top] = nums2[i]
			}
		}
		Stackpush(nums2[i], stack)
	}

	stackPrint(stack)
	fmt.Println(stackMap)
	return nil
}
