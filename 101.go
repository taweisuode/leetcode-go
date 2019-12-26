package LeetCode

import "fmt"

func Code101() {
	head := InitTree()
	head = &TreeNode{1, nil, nil}
	head.Left = &TreeNode{2, nil, nil}
	head.Left.Left = &TreeNode{3, nil, nil}
	head.Left.Right = &TreeNode{4, nil, nil}
	head.Right = &TreeNode{2, nil, nil}
	head.Right.Left = &TreeNode{4, nil, nil}
	head.Right.Right = &TreeNode{3, nil, nil}
	fmt.Println(isSymmetric(head))
}

/**
给定一个二叉树，检查它是否是镜像对称的。

例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

    1
   / \
  2   2
 / \ / \
3  4 4  3
但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:

    1
   / \
  2   2
   \   \
   3    3
*/
func isSymmetric(root *TreeNode) bool {
	return isMirror(root, root)
}

func isMirror(node1 *TreeNode, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 == nil || node2 == nil {
		return false
	}
	return node1.Val == node2.Val && isMirror(node1.Left, node2.Right) && isMirror(node1.Left, node2.Right)
}
