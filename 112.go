package LeetCode

import (
	"fmt"
)

func Code112() {
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
