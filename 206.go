package LeetCode

import "fmt"

func Code206() {
	l1 := InitSingleList([]int{1, 2, 3, 4, 5})
	fmt.Println(reverseList(l1))
}

/**
反转一个单链表。

示例:

输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL

*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	pre := head
	p, q := head, head
	for p.Next != nil {
		q = p.Next
		p.Next = q.Next
		q.Next = pre
		pre = q

	}
	return pre
}
