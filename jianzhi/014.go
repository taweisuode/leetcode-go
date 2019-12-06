/**
 * @Time : 2019-11-28 10:32
 * @Author : zhuangjingpeng
 * @File : 003
 * @Desc : file function description
 */
package jianzhi

import "fmt"

func Code019() {
	list := InitSingleList([]int{1, 2, 3, 4, 5, 6, 7, 8})
	PrintSingleList(list)
	list = removeNthFromEnd(list, 2)
	fmt.Println("after...")
	PrintSingleList(list)
}

/*
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

示例：

给定一个链表: 1->2->3->4->5, 和 n = 2.

当删除了倒数第二个节点后，链表变为 1->2->3->5.
说明：

给定的 n 保证是有效的。

进阶：

你能尝试使用一趟扫描实现吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	pre := &ListNode{}
	pre.Next = head

	p, q := pre, pre

	count := 0
	for p != nil && p.Next != nil {
		if count < n {
			p = p.Next
			count++
		} else {
			break
		}
	}
	fmt.Println(p)
	for p != nil && p.Next != nil {
		p = p.Next
		q = q.Next
	}
	fmt.Println(p, q)
	q.Next = q.Next.Next
	fmt.Println(p, q, q.Next)
	return pre.Next
}
