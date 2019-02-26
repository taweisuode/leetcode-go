package bytedance

import (
	"fmt"
	"strconv"
)

func Code1020() {
	arr := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(arr))

}

/**

给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。

例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]

```
*/

func threeSum(nums []int) [][]int {
	if len(nums) == 0 {
		return make([][]int, 0)
	}
	//fmt.Println(nums)
	newArr := get_sort(nums)
	res := make(map[string][]int, 0)
	for i := 0; i < len(newArr); i++ {
		left := i + 1
		right := len(newArr) - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				addNums := []int{nums[i], nums[left], nums[right]}
				hashKey := strconv.Itoa(nums[i]) + strconv.Itoa(nums[left]) + strconv.Itoa(nums[right])
				res[hashKey] = addNums
			}
			if sum <= 0 {
				left++
			} else {
				right--
			}
		}
	}
	resArr := make([][]int, 0)
	for _, item := range res {
		resArr = append(resArr, item)
	}
	return resArr
}

func get_sort(nums []int) []int {
	if len(nums) <= 0 {
		return nums
	}
	head, tail := 0, len(nums)-1
	key := nums[0]
	for i := 1; i <= tail; {
		if nums[i] > key {
			nums[i], nums[tail] = nums[tail], nums[i]
			tail--
		} else {
			nums[i], nums[head] = nums[head], nums[i]
			head++
			i++
		}
	}
	get_sort(nums[:head])
	get_sort(nums[head+1:])
	return nums
}
