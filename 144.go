package LeetCode

import "fmt"

func Code144() {
	head := &TreeNode{1, nil, nil}
	head.Right = &TreeNode{2, nil, nil}
	head.Right.Left = &TreeNode{3, nil, nil}
	fmt.Println(preorderTraversal(head))
}

/**
给定一个二叉树，返回它的 前序 遍历。

 示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [1,2,3]
进阶: 递归算法很简单，你可以通过迭代算法完成吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-preorder-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var res []int
	stack := &Stack{}
	for !StackEmpty(stack) || root != nil {
		if root != nil {
			Stackpush(root, stack)
			res = append(res, root.Val)
			root = root.Left
		} else {
			popNode := Stackpop(stack).(*TreeNode)
			root = popNode.Right
		}
	}
	return res
}
