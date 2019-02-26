package bytedance

func Code1040() {
	l1 := InitSingleList([]int{4, 2, 1, 3})
	PrintSingleList(sortList(l1))
}

/**

在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序。

示例 1:

输入: 4->2->1->3
输出: 1->2->3->4
示例 2:

输入: -1->5->3->4->0
输出: -1->0->3->4->5
```
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	arr := []*ListNode{}
	p := head
	i := 0
	for p != nil {
		arr = append(arr, p)
		p = p.Next
		i++
	}

	//快排
	newArr := quickSort(arr)
	index := 0
	for index < len(newArr)-1 {
		newArr[index+1].Next = nil
		newArr[index].Next = newArr[index+1]
		index++
	}

	return newArr[0]
}
func quickSort(arr []*ListNode) []*ListNode {
	if len(arr) <= 1 {
		return arr
	}
	head, tail := 0, len(arr)-1
	key := arr[0].Val

	for i := 1; i <= tail; {
		if arr[i].Val >= key {
			arr[i], arr[tail] = arr[tail], arr[i]
			tail--
		} else {
			arr[i], arr[head] = arr[head], arr[i]
			head++
			i++
		}
	}

	quickSort(arr[:head])
	quickSort(arr[head+1:])

	return arr
}
