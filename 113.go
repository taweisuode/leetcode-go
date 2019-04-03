package LeetCode

import (
	"fmt"
)

func Code113() {
	head := InitTree()
	head = &TreeNode{1, nil, nil}
	/*	head.Left = &TreeNode{4, nil, nil}
		head.Right = &TreeNode{8, nil, nil}
		head.Left.Left = &TreeNode{11, nil, nil}
		head.Right.Left = &TreeNode{13, nil, nil}
		head.Right.Right = &TreeNode{4, nil, nil}
		head.Left.Left.Left = &TreeNode{7, nil, nil}
		head.Left.Left.Right = &TreeNode{2, nil, nil}
		head.Right.Right.Left = &TreeNode{5, nil, nil}
		head.Right.Right.Right = &TreeNode{1, nil, nil}*/
	fmt.Println(pathSum(head, 1))
}

/**
给定一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径。

说明: 叶子节点是指没有子节点的节点。

示例:
给定如下二叉树，以及目标和 sum = 22，

              5
             / \
            4   8
           /   / \
          11  13  4
         /  \    / \
        7    2  5   1
返回:

[
   [5,4,11,2],
   [5,8,4,5]
]
*/
func pathSum(root *TreeNode, sum int) [][]int {
	stack := &Stack{}
	resNums = [][]int{}
	return buildArr(root, sum, stack)
}

type BFS struct {
	node  *TreeNode
	depth int
}

var resNums [][]int

func buildArr(root *TreeNode, sum int, stack *Stack) [][]int {
	if root == nil {
		return [][]int{}
	}
	Stackpush(root.Val, stack)
	if root.Left == nil && root.Right == nil {
		if sum == root.Val {
			var res []int
			for _, item := range stack.item {
				if data, ok := item.(int); ok {
					res = append(res, data)
				}
			}
			resNums = append(resNums, res)
		}
	}

	buildArr(root.Left, sum-root.Val, stack)
	buildArr(root.Right, sum-root.Val, stack)
	Stackpop(stack)
	return resNums
}
