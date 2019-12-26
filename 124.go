package LeetCode

import (
	"fmt"
)

func Code124() {
	head := InitTree()
	head = &TreeNode{-10, nil, nil}
	head.Left = &TreeNode{9, nil, nil}
	head.Right = &TreeNode{20, nil, nil}
	head.Right.Left = &TreeNode{15, nil, nil}
	head.Right.Right = &TreeNode{7, nil, nil}
	fmt.Println(maxPathSum(head))
}

/**
给定一个非空二叉树，返回其最大路径和。

本题中，路径被定义为一条从树中任意节点出发，达到任意节点的序列。该路径至少包含一个节点，且不一定经过根节点。

示例 1:

输入: [1,2,3]

       1
      / \
     2   3

输出: 6
示例 2:

输入: [-10,9,20,null,null,15,7]

   -10
   / \
  9  20
    /  \
   15   7

输出: 42

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-maximum-path-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func maxPathSum(root *TreeNode) int {
	max := int(^uint(0) >> 1)
	min := ^max
	fmt.Println(max, min)
	min = -1 << 31
	fmt.Println(min)
	dfs3(root, &min)
	return min
}

func dfs3(root *TreeNode, min *int) int {
	if root == nil {
		return 0
	}

	left := dfs3(root.Left, min)
	right := dfs3(root.Right, min)

	fmt.Println(left, right, 1111)
	if left < 0 {
		return 0
	}
	if right < 0 {
		return 0
	}

	if root.Val+left+right > *min {
		*min = root.Val + left + right
	}
	if left < right {
		return root.Val + right
	} else {
		return root.Val + left
	}
}
