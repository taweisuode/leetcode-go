package LeetCode

import (
	"fmt"
)

func Code119() {
	fmt.Println(getRow(4))
}

/**
给定一个非负索引 k，其中 k ≤ 33，返回杨辉三角的第 k 行。



在杨辉三角中，每个数是它左上方和右上方的数的和。

示例:

输入: 3
输出: [1,3,3,1]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/pascals-triangle-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func getRow(rowIndex int) []int {
	firstRow := []int{1}
	return initAngle2(rowIndex, firstRow)
}

func initAngle2(numRows int, nums []int) []int {
	if len(nums) > numRows {
		return nums
	}
	next := make([]int, len(nums)+1)

	for i := range next {
		if i == 0 || i == len(next)-1 {
			next[i] = 1
		} else {
			next[i] = nums[i-1] + nums[i]
		}
	}

	return initAngle2(numRows, next)
}
