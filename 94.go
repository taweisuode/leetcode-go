package LeetCode

import "fmt"

func Code94() {
	tree1 := InitTree()
	fmt.Println(inorderTraversal(tree1))
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
func inorderTraversal(root *TreeNode) []int {
	list := make([]int, 100)
	fmt.Println(inorderTree(root, list))
	return list
}
func inorderTree(root *TreeNode, list []int) []int {

	if root != nil {
		inorderTree(root.Left, list)
		list = append(list, root.Val)
		inorderTree(root.Right, list)
	}
	return list
}
