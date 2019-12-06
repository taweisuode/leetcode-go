/**
 * @Time : 2019-11-28 10:32
 * @Author : zhuangjingpeng
 * @File : 003
 * @Desc : file function description
 */
package jianzhi

import "fmt"

func Code015() {
	list := InitSingleList([]int{1, 2, 3, 4, 5})
	//PrintSingleList(list)
	list = reverseList(list)
	PrintSingleList(list)
}

/*
反转一个单链表。

示例:

输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL
进阶:
你可以迭代或递归地反转链表。你能否用两种方法解决这道题？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-linked-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := reverseList(head.Next)
	//fmt.Println(lastNode)
	fmt.Println(head, head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
