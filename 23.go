package LeetCode

func Code23() {
	l1 := InitSingleList([]int{-9, 5, 7})
	l2 := InitSingleList([]int{2, 4, 6})
	l3 := InitSingleList([]int{4, 5, 6})
	arr := []*ListNode{l1, l2, l3}
	PrintSingleList(mergeKLists(arr))
}

/**
合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。

示例:

输入:
[
  1->4->5,
  1->3->4,
  2->6
]
输出: 1->1->2->3->4->4->5->6
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {

	for i := 0; i < len(lists); i++ {
		if lists[i] == nil {
			lists = append(lists[:i], lists[i+1:]...)
		}
	}
	if len(lists) == 1 {
		return lists[0]
	}

	head := &ListNode{0, nil}
	finalList := head.Next
	var i, j = 0, 0
	for {
		if i < len(lists)-1 {
			j = i + 1
			finalList = mergeTwoLists(lists[i], lists[j])
			lists[j] = finalList
		} else {
			break
		}
		i++
	}
	return finalList
}
