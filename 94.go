package LeetCode

import "fmt"

func Code94() {
	head := &TreeNode{1, nil, nil}
	head.Right = &TreeNode{2, nil, nil}
	head.Right.Left = &TreeNode{3, nil, nil}
	fmt.Println(inorderTraversal(head))
}

/**
给定一个二叉树，返回它的中序 遍历。

示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [1,3,2]

*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func inorderTraversal1(root *TreeNode) []int {
	list := make([]int, 0)
	inorderTree(root, &list)
	return list
}
func inorderTree(root *TreeNode, list *[]int) []int {

	if root != nil {
		inorderTree(root.Left, list)
		*list = append(*list, root.Val)
		inorderTree(root.Right, list)
	}
	return *list
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	stack := &Stack{}

	var res []int
	for root != nil || !StackEmpty(stack) {
		if root != nil {
			Stackpush(root, stack)
			root = root.Left
		} else {
			popNode := Stackpop(stack).(*TreeNode)
			res = append(res, popNode.Val)
			root = popNode.Right
		}
	}
	return res
}
