package LeetCode

import (
	"fmt"
)

func Code118() {
	fmt.Println(generate(5))
}

/**
给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。



在杨辉三角中，每个数是它左上方和右上方的数的和。

示例:

输入: 5
输出:
[
     [1],
    [1,1],
   [1,2,1],
  [1,3,3,1],
 [1,4,6,4,1]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/pascals-triangle
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func generate(numRows int) [][]int {
	firstRow := []int{1}
	var res [][]int
	initAngle(numRows, firstRow, &res)
	return res
}

func initAngle(numRows int, nums []int, res *[][]int) [][]int {
	if len(nums) > numRows {
		return nil
	}
	*res = append(*res, nums)
	next := make([]int, len(nums)+1)

	for i := range next {
		if i == 0 || i == len(next)-1 {
			next[i] = 1
		} else {
			next[i] = nums[i-1] + nums[i]
		}
	}

	initAngle(numRows, next, res)
	return *res
}
