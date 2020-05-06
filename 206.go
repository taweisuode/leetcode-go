package LeetCode

import (
	"bytes"
	"fmt"
)

func Code206() {
	l1 := InitSingleList([]int{1, 2, 3, 4, 5})
	fmt.Println(reverseList(l1))
}

/**
反转一个单链表。

示例:

输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL

*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	path := []byte("AAAA/BBBB")
	sepIndex := bytes.IndexByte(path, '/')
	dir1 := path[:sepIndex]
	dir2 := path[sepIndex+1:]
	fmt.Println(string(dir1), dir2)
	dir1 = append(dir1, "suffex"...)
	path = bytes.Join([][]byte{dir1, dir2}, []byte{'/'})
	fmt.Println(111, string(dir1), string(dir2))
	fmt.Println(222, string(path))
	fmt.Println(dir1)
	a := []int{1, 2}
	b := []int{3, 4}
	check := a
	a = b
	fmt.Println(a, b, check)
	if head == nil {
		return nil
	}
	var pre *ListNode
	p := head
	for p != nil {
		pre, p, p.Next = p, p.Next, pre

	}
	return pre
}
