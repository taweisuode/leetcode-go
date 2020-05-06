package LeetCode

import "fmt"

func Code102() {
	headTree := InitTree()
	headTree = &TreeNode{3, nil, nil}
	headTree.Left = &TreeNode{9, nil, nil}
	headTree.Right = &TreeNode{20, nil, nil}
	headTree.Right.Left = &TreeNode{15, nil, nil}
	headTree.Right.Right = &TreeNode{7, nil, nil}
	fmt.Println(levelOrder(headTree))
	//fmt.Println("leetcode 102 ")
	//nums := []int{2, 3, 4, 6, 8}
	//target := 8
	//fmt.Println(twoSum(nums, target))
}

func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	queue := &Queue{}
	inQueue(root, queue)
	for !queueEmpty(queue) {
		len := len(queue.item)

		var res1 []int
		for i := 0; i < len; i++ {
			popNode := deQueue(queue).(*TreeNode)
			res1 = append(res1, popNode.Val)
			if popNode.Left != nil {
				inQueue(popNode.Left, queue)
			}
			if popNode.Right != nil {
				inQueue(popNode.Right, queue)
			}
		}
		res = append(res, res1)
	}
	return res
}
