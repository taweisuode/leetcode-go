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
	if root == nil {
		return res
	}
	inQueue(root, queue)
	for !queueEmpty(queue) {
		i := len(queue.item)
		var tempRes []int
		for i > 0 {
			currentNode := deQueue(queue)
			node := currentNode.(*TreeNode)
			tempRes = append(tempRes, node.Val)
			if node.Left != nil {
				inQueue(node.Left, queue)
			}
			if node.Right != nil {
				inQueue(node.Right, queue)

			}
			i--
		}
		res = append(res, tempRes)
	}
	return res
}
