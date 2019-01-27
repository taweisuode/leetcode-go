package LeetCode

func Code83() {
	l1 := InitSingleList([]int{1, 1, 2, 2})
	PrintSingleList(deleteDuplicates1(l1))
}

/**
给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。

示例 1:

输入: 1->1->2
输出: 1->2
示例 2:

输入: 1->1->2->3->3
输出: 1->2->3
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	pre := &ListNode{}
	pre.Next = head
	p, q := pre, head

	for q != nil {
		if q.Next != nil && q.Val == q.Next.Val {
			p.Next = q.Next
		} else {
			p = q
		}
		q = q.Next
	}
	return pre.Next

}
