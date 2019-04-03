package LeetCode

import (
	"fmt"
	"math"
)

func Code16() {
	name := []int{1, 2, 4, 8, 16, 32, 64, 128}
	target := 82
	fmt.Println(threeSumClosest(name, target))
}

/**
给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。

例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.

与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).

```
*/
func threeSumClosest(nums []int, target int) int {
	newNums := quick_sort(nums)
	if len(nums) < 3 {
		return 0
	}

	resMap := make(map[float64]int, 0)
	for i := 0; i < len(nums); i++ {
		left := i + 1
		right := len(nums) - 1
		for left < right {
			fmt.Println(newNums[left], newNums[right], newNums[i])
			sum := newNums[left] + newNums[right] + newNums[i]
			resMap[math.Abs(float64(sum-target))] = sum
			if sum-target < 0 {
				left++
			} else if sum-target > 0 {
				right--
			} else {
				return sum
			}
		}

	}
	fmt.Println(resMap)
	minNum := math.MaxFloat64
	for key, _ := range resMap {
		minNum = minfloat(minNum, key)
	}
	return resMap[minNum]

}
func minfloat(a float64, b float64) float64 {
	if a > b {
		return b
	} else {
		return a
	}
}

func quick_sort(nums []int) []int {
	if len(nums) <= 0 {
		return []int{}
	}
	head, tail := 0, len(nums)-1
	key := nums[0]
	for i := 1; i <= tail; {
		if nums[i] > key {
			nums[i], nums[tail] = nums[tail], nums[i]
			tail--
		} else {
			nums[i], nums[head] = nums[head], nums[i]
			i++
			head++
		}
	}
	quick_sort(nums[:head])
	quick_sort(nums[head+1:])
	return nums
}
