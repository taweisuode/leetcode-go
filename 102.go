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
		size := len(queue.item)

		var res1 []int
		for i := size; i > 0; i-- {
			cur := deQueue(queue)
			curTree := cur.(*TreeNode)
			res1 = append(res1, curTree.Val)
			if curTree.Left != nil {
				inQueue(curTree.Left, queue)

			}

			if curTree.Right != nil {
				inQueue(curTree.Right, queue)
			}

		}
		if len(res1) > 0 {
			res = append(res, res1)
		}
	}

	fmt.Println(res)
	return res
}
