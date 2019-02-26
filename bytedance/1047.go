package bytedance

import (
	"fmt"
	"sort"
)

func Code1047() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(height))
}

/**

给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。



上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 感谢 Marcos 贡献此图。

示例:

输入: [0,1,0,2,1,0,1,3,2,1,2,1]
输入: [0,1,1,2,2,2,2,3,2,2,2,1]
输出: 6
```
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func trap(height []int) int {

	//左边找第一个进入升序的数字直到最高值
	//从最高值降序 直到最小值
	if len(height) == 0 {
		return 0
	}
	maxNum := 0
	maxIndex := 0
	for key, item := range height {
		if item > maxNum {
			maxNum = item
			maxIndex = key
		}
	}

	sortNum := make(map[int]int, 0)
	rsortNum := make(map[int]int, 0)
	first, tail := height[0], height[len(height)-1]
	for i := 0; i < maxIndex; i++ {
		if height[i] >= first {
			sortNum[i] = height[i]
			first = height[i]
		} else {
			sortNum[i] = first
		}
	}
	sortNum[maxIndex] = maxNum
	for j := len(height) - 1; j > maxIndex; j-- {
		if height[j] >= tail {
			rsortNum[j] = height[j]
			tail = height[j]
		} else {
			rsortNum[j] = tail
		}
	}
	sortArr := make([]int, 0)
	rsortArr := make([]int, 0)
	for _, item := range sortNum {
		sortArr = append(sortArr, item)
	}
	for _, item := range rsortNum {
		rsortArr = append(rsortArr, item)
	}
	sort.Ints(sortArr)
	sort.Sort(sort.Reverse(sort.IntSlice(rsortArr)))
	totalSum := 0
	for key, item := range height {
		if key <= maxIndex {
			each := sortArr[key] - item
			totalSum += each
		} else {
			each := rsortArr[key-maxIndex-1] - item
			totalSum += each
		}
	}
	return totalSum
}
