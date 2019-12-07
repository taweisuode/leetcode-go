package LeetCode

import (
	"fmt"
	"strconv"
)

func Code257() {
	head := InitTree()
	head = &TreeNode{1, nil, nil}
	head.Left = &TreeNode{2, nil, nil}
	head.Right = &TreeNode{3, nil, nil}
	head.Left.Right = &TreeNode{5, nil, nil}
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

func binaryTreePaths1(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	stack := &Stack{}
	var res []string
	var eachList string
	for root != nil || !StackEmpty(stack) {
		if root != nil {
			if eachList == "" {
				eachList = strconv.Itoa(root.Val)
			} else {
				eachList += "->" + strconv.Itoa(root.Val)
			}

			Stackpush(root, stack)
			root = root.Left
		} else {
			popNode := Stackpop(stack).(*TreeNode)
			if popNode.Left == nil && popNode.Right == nil {
				res = append(res, eachList)
			}
			root = popNode.Right
		}
	}
	fmt.Println(res)
	return nil
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	var res []string
	return dfs(root, "", &res)
}
func dfs(root *TreeNode, str string, res *[]string) []string {
	str += strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		*res = append(*res, str)
	}
	str += "->"
	if root.Left != nil {
		dfs(root.Left, str, res)
	}
	if root.Right != nil {
		dfs(root.Right, str, res)
	}
	return *res
}
