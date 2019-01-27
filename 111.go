package LeetCode

import (
	"fmt"
)

func Code111() {
	head := InitTree()
	head = &TreeNode{3, nil, nil}
	head.Left = &TreeNode{9, nil, nil}
	head.Right = &TreeNode{20, nil, nil}
	head.Right.Left = &TreeNode{15, nil, nil}
	head.Right.Right = &TreeNode{7, nil, nil}
	fmt.Println(minDepth(head))
}

/**
给定一个二叉树，找出其最小深度。

最小深度是从根节点到最近叶子节点的最短路径上的节点数量。

说明: 叶子节点是指没有子节点的节点。

示例:

给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回它的最小深度  2.
*/
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil {
		return minDepth(root.Right) + 1
	}
	if root.Right == nil {
		return minDepth(root.Left) + 1
	}
	min := min(minDepth(root.Left), minDepth(root.Right)) + 1
	return int(min)
}
func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
