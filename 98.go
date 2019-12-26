package LeetCode

import "fmt"

func Code98() {
	head := InitTree()
	head = &TreeNode{5, nil, nil}
	head.Left = &TreeNode{1, nil, nil}
	head.Right = &TreeNode{4, nil, nil}
	head.Right.Left = &TreeNode{3, nil, nil}
	head.Right.Right = &TreeNode{6, nil, nil}
	fmt.Println(isValidBST(head))
}

/**
给定一个二叉树，判断其是否是一个有效的二叉搜索树。

假设一个二叉搜索树具有如下特征：

节点的左子树只包含小于当前节点的数。
节点的右子树只包含大于当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。
示例 1:

输入:
    2
   / \
  1   3
输出: true
示例 2:

输入:
    5
   / \
  1   4
     / \
    3   6
输出: false
解释: 输入为: [5,1,4,null,null,3,6]。
     根节点的值为 5 ，但是其右子节点值为 4 。

*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST1(root *TreeNode) bool {

	var list []int
	res := inOrder(root, &list)
	for i := 1; i < len(res); i++ {
		if res[i] <= res[i-1] {
			return false
		}
	}
	fmt.Println(res)
	return true
}

func inOrder(root *TreeNode, res *[]int) []int {
	if root != nil {
		inOrder(root.Left, res)
		*res = append(*res, root.Val)
		inOrder(root.Right, res)
	}
	return *res
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	res := inOrderBst(root)
	for i := 1; i < len(res); i++ {
		if res[i] <= res[i-1] {
			return false
		}
	}
	fmt.Println(res)
	return true
}

func inOrderBst(root *TreeNode) []int {
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
