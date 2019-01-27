package LeetCode

func Code24() {
	l1 := InitSingleList([]int{1, 2, 3, 4})
	PrintSingleList(swapPairs(l1))
}

/**
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

示例:

给定 1->2->3->4, 你应该返回 2->1->4->3.
说明:

你的算法只能使用常数的额外空间。
你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	var p, q, temp = &ListNode{}, &ListNode{}, &ListNode{}
	if head == nil {
		return nil
	}
	for {
		if head != nil {
			p = head
			q = head.Next
			temp = p
			p = q
			q = temp
			head = head.Next
		}
	}
}
