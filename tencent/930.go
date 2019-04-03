package tencent

/**

给定一个二叉树，找出其最大深度。

二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

说明: 叶子节点是指没有子节点的节点。

示例：
给定二叉树 [3,9,20,null,null,15,7]，

    3
   / \
  9  20
    /  \
   15   7
返回它的最大深度 3 。

*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftLength := maxDepth(root.Left) + 1
	rightLenth := maxDepth(root.Right) + 1
	return Max(leftLength, rightLenth)

}
func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
