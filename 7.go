package LeetCode

import (
	"fmt"
)

func Code7() {
	fmt.Println(reverse1(-123))
}

/**
给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

示例 1:

输入: 123
输出: 321
 示例 2:

输入: -123
输出: -321
示例 3:

输入: 120
输出: 21
注意:

假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−231,  231 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverse1(x int) int {
	res := 0
	num := 0
	for x != 0 {
		a := x % 10
		res = 10*num + a
		num = res
		x = x / 10
		maxInt := 1<<31 - 1
		minInt := -1 << 31

		if res > maxInt || res < minInt {
			return 0
		}
	}
	return res
}
