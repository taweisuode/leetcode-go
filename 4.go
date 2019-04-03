package LeetCode

import (
	"fmt"
	"sort"
)

func Code4() {
	num1 := []int{}
	num2 := []int{1}
	fmt.Println(findMedianSortedArrays(num1, num2))
}

/**
给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。

请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。

你可以假设 nums1 和 nums2 不会同时为空。

示例 1:

nums1 = [1, 3]
nums2 = [2]

则中位数是 2.0
示例 2:

nums1 = [1, 2]
nums2 = [3, 4]

则中位数是 (2 + 3)/2 = 2.5
```
*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	for _, item := range nums2 {
		nums1 = append(nums1, item)
	}
	sort.Ints(nums1)
	//获取num1 的中位数
	midNum1 := float64(0)
	if len(nums1) > 0 {
		if len(nums1)%2 == 0 {
			midNum1 = float64(nums1[len(nums1)/2-1]+nums1[len(nums1)/2]) / 2
		} else {
			midNum1 = float64(nums1[len(nums1)/2])
		}
	}
	return midNum1
}
