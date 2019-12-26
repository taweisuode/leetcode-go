package LeetCode

import (
	"fmt"
)

func Code199() {
	head := InitTree()
	head = &TreeNode{1, nil, nil}
	head.Left = &TreeNode{2, nil, nil}
	head.Right = &TreeNode{3, nil, nil}
	head.Left.Right = &TreeNode{5, nil, nil}
	head.Right.Right = &TreeNode{4, nil, nil}
	fmt.Println(rightSideView(head))
}

/**
给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。

示例:

输入: [1,2,3,null,5,null,4]
输出: [1, 3, 4]
解释:

   1            <---
 /   \
2     3         <---
 \     \
  5     4       <---

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-right-side-view
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	queue := &Queue{}
	inQueue(root, queue)
	for !queueEmpty(queue) {
		size := len(queue.item)
		var level []int
		for c := size; c > 0; c-- {
			popNode := deQueue(queue).(*TreeNode)
			level = append(level, popNode.Val)
			if popNode.Left != nil {
				inQueue(popNode.Left, queue)
			}
			if popNode.Right != nil {
				inQueue(popNode.Right, queue)
			}
		}
		res = append(res, level[len(level)-1])
	}
	return res
}
