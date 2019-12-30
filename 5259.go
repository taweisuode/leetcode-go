package LeetCode

import (
	"fmt"
)

func Code5259() {
	fmt.Println(subtractProductAndSum(4421))
}

/**
删除链表中等于给定值 val 的所有节点。

示例:

输入: 1->2->6->3->4->5->6, val = 6
输出: 1->2->3->4->5
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func subtractProductAndSum(n int) int {
	a := 1
	b := 0
	for n > 0 {
		a *= n % 10
		b += n % 10
		n /= 10
	}
	return a - b
}
