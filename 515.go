package LeetCode

import (
	"fmt"
)

func Code515() {
	head := InitTree()
	head = &TreeNode{1, nil, nil}
	head.Left = &TreeNode{3, nil, nil}
	head.Right = &TreeNode{2, nil, nil}
	head.Left.Left = &TreeNode{5, nil, nil}
	head.Left.Right = &TreeNode{3, nil, nil}
	head.Right.Right = &TreeNode{9, nil, nil}
	fmt.Println(largestValues(head))
}

/**
您需要在二叉树的每一行中找到最大的值。

示例：

输入:

          1
         / \
        3   2
       / \   \
      5   3   9

输出: [1, 3, 9]
在真实的面试中遇到过这道题？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-largest-value-in-each-tree-row
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func largestValues(root *TreeNode) []int {
	var res []int
	queue := &Queue{}
	inQueue(root, queue)
	for !queueEmpty(queue) {
		size := len(queue.item)
		var level []int
		for x := size; x > 0; x-- {
			popNode := deQueue(queue).(*TreeNode)
			level = append(level, popNode.Val)
			if popNode.Left != nil {
				inQueue(popNode.Left, queue)
			}
			if popNode.Right != nil {
				inQueue(popNode.Right, queue)
			}
		}
		fmt.Println(level)
		if len(level) > 0 {
			a := getMax(level)
			res = append(res, a)
		}
	}
	return res
}
func getMax(level []int) int {
	first := level[0]
	for i := 1; i < len(level); i++ {
		if level[i] > first {
			first = level[i]
		}
	}
	return first
}
