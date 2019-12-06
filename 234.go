package LeetCode

import "fmt"

func Code234() {

	list := InitSingleList([]int{0, 0})
	fmt.Println(isPalindrome(list))
}

/**
请判断一个链表是否为回文链表。

示例 1:

输入: 1->2
输出: false
示例 2:

输入: 1->2->2->1
输出: true
进阶：
你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/palindrome-linked-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	pre := &ListNode{}
	pre.Next = head
	slow, fast := pre, pre

	fmt.Println(slow, fast)
	var arr []int
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		arr = append(arr, slow.Val)
		fast = fast.Next.Next
	}

	slow = slow.Next
	for slow != nil {
		arr = append(arr, slow.Val)
		slow = slow.Next
	}

	return isPalindromeList(arr)
}

func isPalindromeList(arr []int) bool {
	for index := range arr {
		if arr[index] != arr[len(arr)-index-1] {
			return false
		}
	}
	return true
}
