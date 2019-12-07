package LeetCode

import (
	"fmt"
)

func Code637() {
	head := InitTree()
	head = &TreeNode{3, nil, nil}
	head.Left = &TreeNode{9, nil, nil}
	head.Right = &TreeNode{20, nil, nil}
	head.Right.Left = &TreeNode{15, nil, nil}
	head.Right.Right = &TreeNode{7, nil, nil}
	fmt.Println(averageOfLevels(head))
}

/**
给定一个非空二叉树, 返回一个由每层节点平均值组成的数组.

示例 1:

输入:
    3
   / \
  9  20
    /  \
   15   7
输出: [3, 14.5, 11]
解释:
第0层的平均值是 3,  第1层是 14.5, 第2层是 11. 因此返回 [3, 14.5, 11].
注意：

节点值的范围在32位有符号整数范围内。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/average-of-levels-in-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}
	queue := &Queue{}
	inQueue(root, queue)
	var level [][]int
	for !queueEmpty(queue) {
		size := len(queue.item)
		var eachLevel []int
		for i := size; i > 0; i-- {
			popNode := deQueue(queue).(*TreeNode)
			eachLevel = append(eachLevel, popNode.Val)
			if popNode.Left != nil {
				inQueue(popNode.Left, queue)

			}
			if popNode.Right != nil {
				inQueue(popNode.Right, queue)
			}
		}
		level = append(level, eachLevel)
	}
	var newList []float64
	for i := range level {
		eachCount := 0
		for j := range level[i] {
			eachCount += level[i][j]
		}
		fmt.Println(eachCount)
		eachFloat := float64(eachCount) / float64(len(level[i]))
		newList = append(newList, eachFloat)
	}
	return newList
}
