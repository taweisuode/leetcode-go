package LeetCode

func Code82() {
	l1 := InitSingleList([]int{1, 1})
	PrintSingleList(deleteDuplicates(l1))
}

/**
给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。

示例 1:

输入: 1->2->3->3->4->4->5
输出: 1->2->5
示例 2:

输入: 1->1->1->2->3
输出: 2->3
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	head1 := &ListNode{}
	head1.Next = head
	p, q := head1, head

	for q != nil {
		duplicate := false
		for q.Next != nil && q.Val == q.Next.Val {
			q = q.Next
			duplicate = true
		}

		if !duplicate {
			p = q
		} else {
			p.Next = q.Next
		}
		q = q.Next
	}
	return head1.Next

}
