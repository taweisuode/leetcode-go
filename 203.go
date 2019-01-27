package LeetCode

import "fmt"

func Code203() {
	l1 := InitSingleList([]int{1, 2, 6, 3, 4, 5, 6})
	fmt.Println(removeElements(l1, 6))
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
func removeElements(head *ListNode, val int) *ListNode {
	for head != nil && head.Val == val {
		head = head.Next
	}
	if head == nil {
		return head
	}
	pre := &ListNode{}
	pre.Next = head
	p, q := head, head
	for q != nil {
		if q.Val == val {
			p.Next = q.Next
		} else {
			p = q
		}
		q = q.Next
	}
	return pre.Next
}
