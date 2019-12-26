package LeetCode

import (
	"fmt"
)

func Code513() {
	head := InitTree()
	head = &TreeNode{1, nil, nil}
	head.Left = &TreeNode{2, nil, nil}
	head.Right = &TreeNode{3, nil, nil}
	/*head.Left.Left = &TreeNode{4, nil, nil}
	head.Right.Left = &TreeNode{5, nil, nil}
	head.Right.Right = &TreeNode{6, nil, nil}
	head.Right.Left.Left = &TreeNode{7, nil, nil}*/
	fmt.Println(findBottomLeftValue(head))
}

/**
给定一个二叉树，在树的最后一行找到最左边的值。

示例 1:

输入:

    2
   / \
  1   3

输出:
1


示例 2:

输入:

        1
       / \
      2   3
     /   / \
    4   5   6
       /
      7

输出:
7

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-bottom-left-tree-value
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func findBottomLeftValue(root *TreeNode) int {
	queue := &Queue{}
	inQueue(root, queue)
	popNode := &TreeNode{}
	for !queueEmpty(queue) {
		popNode = deQueue(queue).(*TreeNode)

		if popNode.Right != nil {
			inQueue(popNode.Right, queue)
		}
		if popNode.Left != nil {
			inQueue(popNode.Left, queue)
		}
	}
	return popNode.Val
}
