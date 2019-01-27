package LeetCode

import "fmt"

func Code102() {
	headTree := InitTree()
	headTree = &TreeNode{3, nil, nil}
	headTree.Left = &TreeNode{9, nil, nil}
	headTree.Right = &TreeNode{20, nil, nil}
	headTree.Right.Left = &TreeNode{15, nil, nil}
	headTree.Right.Right = &TreeNode{7, nil, nil}
	levelOrder(headTree)
	//fmt.Println("leetcode 102 ")
	//nums := []int{2, 3, 4, 6, 8}
	//target := 8
	//fmt.Println(twoSum(nums, target))
}

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{nil}
	queue := &Queue{}
	if root == nil {
		return res
	}
	i := 0
	inQueue(root, queue)
	res[i] = append(res[i], root.Val)
	fmt.Println(res)
	for !queueEmpty(queue) {

		currentNode := deQueue(queue)
		i++
		if node, ok := currentNode.(*TreeNode); ok {
			fmt.Println(node.Val)
			if node.Left != nil {
				inQueue(node.Left, queue)
				fmt.Println("--", node.Val)
				res[i] = append(res[i], 1)
			}
			if node.Right != nil {
				inQueue(node.Right, queue)
				res[i] = append(res[i], node.Val)
			}
		}

	}
	return res
}
