package bytedance

import (
	"fmt"
)

func Code1019() {
	arr := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(longestConsecutive(arr))

}

/**

给定一个未排序的整数数组，找出最长连续序列的长度。

要求算法的时间复杂度为 O(n)。

示例:

输入: [100, 4, 200, 1, 3, 2]
输出: 4
解释: 最长连续序列是 [1, 2, 3, 4]。它的长度为 4。
```
*/

/*func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	newArr := quick_sort(nums)

	fmt.Println(newArr)

	filter := make(map[int]int, 0)
	filter[newArr[0]] = 1
	for i := 1; i < len(newArr); {
		if newArr[i] == newArr[i-1]+1 {
			filter[newArr[i]] = filter[newArr[i-1]] + 1
		}
		if newArr[i] == newArr[i-1] {
			filter[newArr[i]] = filter[newArr[i-1]]
		}
		if newArr[i] != newArr[i-1]+1 && newArr[i] != newArr[i-1] {
			filter[newArr[i]] = 1
		}
		i++
	}
	fmt.Println(filter)
	maxLen := 1
	for _, item := range filter {
		if item > maxLen {
			maxLen = item
		}
	}
	return maxLen
}*/

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	hashMap := make(map[int]bool, 0)
	for _, item := range nums {
		hashMap[item] = true
	}
	maxLen := 1
	for _, num := range nums {
		count := 1
		i := num
		for {
			i--
			if _, ok := hashMap[i]; ok {
				delete(hashMap, i)
				count++
			} else {
				break
			}
		}
		j := num
		for {
			j++
			if _, ok := hashMap[j]; ok {
				delete(hashMap, j)
				count++
			} else {
				break
			}

		}
		maxLen = max(maxLen, count)

	}
	return maxLen
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func quick_sort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	first, tail := 0, len(nums)-1
	key := nums[0]
	for i := 1; i <= tail; {
		if nums[i] > key {
			nums[i], nums[tail] = nums[tail], nums[i]
			tail--
		} else {
			nums[i], nums[first] = nums[first], nums[i]
			first++
			i++
		}
	}
	quick_sort(nums[:first])
	quick_sort(nums[first+1:])
	return nums
}
