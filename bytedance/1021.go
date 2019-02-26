package bytedance

import (
	"fmt"
	"sort"
	"strconv"
)

func Code1021() {
	fmt.Println(getPermutation(4, 9))

}

/**

给出集合 [1,2,3,…,n]，其所有元素共有 n! 种排列。

按大小顺序列出所有排列情况，并一一标记，当 n = 3 时, 所有排列如下：

"123"
"132"
"213"
"231"
"312"
"321"

给定 n 和 k，返回第 k 个排列。

说明：

给定 n 的范围是 [1, 9]。
给定 k 的范围是[1,  n!]。
示例 1:

输入: n = 3, k = 3
输出: "213"
示例 2:

输入: n = 4, k = 9
输出: "2314"
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

func getPermutation(n int, k int) string {

	res := make([]int, 0)
	numsMap := make(map[int]bool, 0)
	for i := 1; i <= n; i++ {
		numsMap[i] = false
	}
	for i := 1; i <= n; i++ {
		canMod := k % jiechen(n-i)
		firstChar := k / jiechen(n-i)
		if canMod != 0 {
			firstChar++
		}
		numsArr := getUserfulArr(numsMap)
		if numsMap[numsArr[firstChar-1]] == false {
			res = append(res, numsArr[firstChar-1])
		}
		numsMap[numsArr[firstChar-1]] = true
		k = k - (firstChar-1)*jiechen(n-i)
	}

	resStr := ""
	for _, item := range res {
		resStr = resStr + strconv.Itoa(item)
	}

	return resStr
}
func jiechen(num int) int {
	i := num
	res := 1
	for i > 0 {
		res = res * i
		i--
	}
	return res
}
func getUserfulArr(numsMap map[int]bool) []int {
	res := make([]int, 0)
	for key, item := range numsMap {
		if item == false {
			res = append(res, key)
		}
	}
	sort.Ints(res)
	return res
}
