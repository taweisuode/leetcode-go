package LeetCode

import (
	"fmt"
)

func Code257() {
	head := InitTree()
	head = &TreeNode{1, nil, nil}
	head.Left = &TreeNode{2, nil, nil}
	head.Right = &TreeNode{3, nil, nil}
	head.Right.Right = &TreeNode{5, nil, nil}
	fmt.Println(binaryTreePaths(head))
}

/**
给定一个二叉树，返回所有从根节点到叶子节点的路径。

说明: 叶子节点是指没有子节点的节点。

示例:

输入:

   1
 /   \
2     3
 \
  5

输出: ["1->2->5", "1->3"]

解释: 所有根节点到叶子节点的路径为: 1->2->5, 1->3

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-paths
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

var NodeRes []string

func binaryTreePaths(root *TreeNode) []string {
	var path string
	dfs(root, path, NodeRes)
	return NodeRes
}

func dfs(root *TreeNode, path string, res []string) {
}
