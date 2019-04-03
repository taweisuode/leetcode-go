package bytedance

import (
	"fmt"
)

func Code85() {
	head := InitTree()
	head = &TreeNode{1, nil, nil}
	head.Right = &TreeNode{2, nil, nil}
	head.Right.Left = &TreeNode{3, nil, nil}
	fmt.Println(postorderTraversal(head))

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
进阶: 递归算法很简单，你可以通过迭代算法完成吗？
```
*/

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	stack := &Stack{}
	head := root
	res := make([]int, 0)

	for head != nil || len(stack.item) != 0 {
		if head != nil {
			Stackpush(head, stack)
			head = head.Left
		} else {
			a := Stackpop(stack)
			if item, ok := a.(*TreeNode); ok {
				res = append(res, item.Val)
				head = item.Right
			}
		}
	}
	return res
}
func inorder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	stack := &Stack{}
	head := root
	res := make([]int, 0)
	for root != nil || len(stack.item) > 0 {
		if root != nil {
			Stackpush(head, stack)
			head = head.Left
		} else {
			a := Stackpop(stack)
			if item, ok := a.(*TreeNode); ok {
				res = append(res, item.Val)
				head = item.Right
			}
		}
	}
	return res

}
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	stack := &Stack{}
	head := root
	res := make([]int, 0)

	for head != nil || len(stack.item) != 0 {
		if head != nil {
			res = append(res, head.Val)
			Stackpush(head, stack)
			head = head.Left
		} else {
			a := Stackpop(stack)
			if item, ok := a.(*TreeNode); ok {
				head = item.Right
			}
		}
	}
	return res
}

func postorderTraversal(root *TreeNode) []int {
	stack1 := &Stack{}
	stack2 := &Stack{}
	res := make([]int, 0)
	Stackpush(root, stack1)

	for len(stack1.item) > 0 {
		elem := Stackpop(stack1).(*TreeNode)
		Stackpush(elem, stack2)

		if elem.Left != nil {
			Stackpush(elem.Left, stack1)
		}

		if elem.Right != nil {
			Stackpush(elem.Right, stack1)
		}
	}

	for len(stack2.item) > 0 {
		elem := Stackpop(stack2).(*TreeNode)
		res = append(res, elem.Val)
	}
	return res
}

func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i < n; i++ {
		for j := 0; j < i; j++ {
			dp[i] += dp[j] * dp[i-j-1]
		}
	}
	return dp[n]
}
