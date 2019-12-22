package LeetCode

import (
	"fmt"
	"math"
)

func Code110() {
	head := InitTree()
	head = &TreeNode{1, nil, nil}
	head.Left = &TreeNode{2, nil, nil}
	head.Right = &TreeNode{2, nil, nil}
	head.Left.Left = &TreeNode{3, nil, nil}
	head.Left.Right = &TreeNode{3, nil, nil}
	head.Left.Left.Left = &TreeNode{4, nil, nil}
	head.Left.Left.Right = &TreeNode{4, nil, nil}
	fmt.Println(isBalanced(head))
}

/**
给定一个二叉树，判断它是否是高度平衡的二叉树。

本题中，一棵高度平衡二叉树定义为：

一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过1。

示例 1:

给定二叉树 [3,9,20,null,null,15,7]

    3
   / \
  9  20
    /  \
   15   7
返回 true 。

示例 2:

给定二叉树 [1,2,2,3,3,null,null,4,4]

       1
      / \
     2   2
    / \
   3   3
  / \
 4   4
返回 false 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/balanced-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func isBalanced(root *TreeNode) bool {
	return dfs_110(root) != -1
}
func dfs_110(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := dfs_110(root.Left)
	if left == -1 {
		return -1
	}
	right := dfs_110(root.Right)
	if right == -1 {
		return -1
	}
	if math.Abs(float64(left-right)) < 2 {
		return max(left, right) + 1
	}
	return -1
}
