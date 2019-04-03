package LeetCode

import "fmt"

func Code98() {
	head := InitTree()
	head = &TreeNode{0, nil, nil}
	//head.Left = &TreeNode{1, nil, nil}
	//head.Right = &TreeNode{3, nil, nil}
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
var res []int

func isValidBST(root *TreeNode) bool {
	if root == nil || (root != nil && root.Left == nil && root.Right == nil) {
		return true
	}
	res = inOrder(root)
	for i := 1; i < len(res); i++ {
		if res[i] <= res[i-1] {
			return false
		}
	}
	fmt.Println()
	return true
}

func inOrder(root *TreeNode) []int {
	if root != nil {
		inOrder(root.Left)
		res = append(res, root.Val)
		inOrder(root.Right)
	}
	return res
}
