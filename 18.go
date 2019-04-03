package LeetCode

import (
	"fmt"
	"strconv"
)

func Code18() {
	name := []int{1, 0, -1, 0, -2, 2}
	target := 0
	fmt.Println(fourSum(name, target))
}

/**
给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。

注意：

答案中不可以包含重复的四元组。

示例：

给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。

满足要求的四元组集合为：
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]

```
*/
func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return [][]int{}
	}
	newArr := quick_sort(nums)
	res := make(map[string][]int, 0)
	resArr := make([][]int, 0)
	for i := 0; i < len(newArr); i++ {
		for j := i + 1; j < len(newArr); j++ {
			left := j + 1
			right := len(newArr) - 1
			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum == target {
					addNums := []int{nums[i], nums[j], nums[left], nums[right]}
					hashKey := strconv.Itoa(nums[i]) + strconv.Itoa(nums[j]) + strconv.Itoa(nums[left]) + strconv.Itoa(nums[right])
					res[hashKey] = addNums
				}
				if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}
	for _, item := range res {
		resArr = append(resArr, item)
	}

	return resArr
}
