/**
 * @Time : 2019-11-28 10:32
 * @Author : zhuangjingpeng
 * @File : 003
 * @Desc : file function description
 */
package jianzhi

import "fmt"

func Code083() {
	list := InitSingleList([]int{1, 1, 2, 2, 3, 3})
	list = deleteDuplicates(list)
	PrintSingleList(list)
}

/**
给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。

示例 1:

输入: 1->1->2
输出: 1->2
示例 2:

输入: 1->1->2->3->3
输出: 1->2->3

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func deleteDuplicates(head *ListNode) *ListNode {
	nodeMap := make(map[int]*ListNode)
	root := head

	for root != nil {
		if item, ok := nodeMap[root.Val]; ok {

			fmt.Println(item)
			item.Next = root.Next
		} else {
			nodeMap[root.Val] = root
		}
		root = root.Next
	}

	return head
}
