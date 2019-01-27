package LeetCode

import "fmt"

func Code107() {
	head := InitTree()
	head = &TreeNode{3, nil, nil}
	head.Left = &TreeNode{9, nil, nil}
	head.Right = &TreeNode{20, nil, nil}
	head.Right.Left = &TreeNode{15, nil, nil}
	head.Right.Right = &TreeNode{7, nil, nil}
	fmt.Println(isSymmetric(head))
}

/**
给定一个二叉树，返回其节点值自底向上的层次遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）

例如：
给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其自底向上的层次遍历为：

[
  [15,7],
  [9,20],
  [3]
]
*/
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	stack := []*TreeNode{root}
	rets := [][]int{}
	for len(stack) != 0 {
		tmpStack := []*TreeNode{}
		ret := []int{}
		for i := range stack {
			ret = append(ret, stack[i].Val)
			if stack[i].Left != nil {
				tmpStack = append(tmpStack, stack[i].Left)
			}
			if stack[i].Right != nil {
				tmpStack = append(tmpStack, stack[i].Right)
			}
		}
		rets = append(rets, ret)
		stack = tmpStack
	}
	return rets
}
