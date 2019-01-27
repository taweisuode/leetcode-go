package LeetCode

func Code21() {
	l1 := InitSingleList([]int{1, 2, 4})
	l2 := InitSingleList([]int{1, 3, 4})
	PrintSingleList(mergeTwoLists(l1, l2))
}

/**
将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	p := &ListNode{0, nil}

	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val >= l2.Val {
		p = l2
		p.Next = mergeTwoLists(l1, l2.Next)
	} else {
		p = l1
		p.Next = mergeTwoLists(l1.Next, l2)
	}
	return p
}
