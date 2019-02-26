package bytedance

func Code1023() {
	l1 := InitSingleList([]int{4, 2, 1, 3})
	PrintSingleList(sortList(l1))

}

/**

给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。

为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。

说明：不允许修改给定的链表。



示例 1：

输入：head = [3,2,0,-4], pos = 1
输出：tail connects to node index 1
解释：链表中有一个环，其尾部连接到第二个节点。


示例 2：

输入：head = [1,2], pos = 0
输出：tail connects to node index 0
解释：链表中有一个环，其尾部连接到第一个节点。


示例 3：

输入：head = [1], pos = -1
输出：no cycle
解释：链表中没有环。



```
*/

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head

	step1 := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			//这里再设置2个指针  同时走 相遇的点就是入口
			return letGo(step1, fast)
		}
	}
	return nil
}
func letGo(head *ListNode, fast *ListNode) *ListNode {
	for fast != nil {
		if head == fast {
			return head
		}
		head = head.Next
		fast = fast.Next
	}
	return head
}
