package LeetCode

func Code2() {
	l1 := InitSingleList([]int{5})
	l2 := InitSingleList([]int{5})
	PrintSingleList(addTwoNumbers(l1, l2))
}

/**
给定两个非空链表来表示两个非负整数。位数按照逆序方式存储，它们的每个节点只存储单个数字。将两数相加返回一个新的链表。

你可以假设除了数字 0 之外，这两个数字都不会以零开头。

示例：
```
输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
```
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	p := l1
	q := l2
	head := &ListNode{0, nil}
	curr := head
	carry := 0
	x := 0
	y := 0
	for {
		if p != nil || q != nil {
			if p != nil {
				x = p.Val
			} else {
				x = 0
			}
			if q != nil {
				y = q.Val
			} else {
				y = 0
			}
			sum := carry + x + y
			carry = sum / 10
			curr.Next = &ListNode{sum % 10, nil}
			curr = curr.Next
			if p != nil {
				p = p.Next
			}
			if q != nil {
				q = q.Next
			}
		} else {
			break
		}
	}
	if carry > 0 {
		curr.Next = &ListNode{carry % 10, nil}
	}
	return head.Next
}
