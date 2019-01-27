package LeetCode

import "fmt"

func Code103() {
	headTree := InitTree()
	headTree = &TreeNode{3, nil, nil}
	headTree.Left = &TreeNode{9, nil, nil}
	headTree.Right = &TreeNode{20, nil, nil}
	headTree.Right.Left = &TreeNode{15, nil, nil}
	headTree.Right.Right = &TreeNode{7, nil, nil}
	fmt.Println(zigzagLevelOrder(headTree))
	//fmt.Println("leetcode 102 ")
	//nums := []int{2, 3, 4, 6, 8}
	//target := 8
	//fmt.Println(twoSum(nums, target))
}

/**
 给定一个二叉树，返回其节点值的锯齿形层次遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

例如：
给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回锯齿形层次遍历如下：

[
  [3],
  [20,9],
  [15,7]
]
*/
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	stack := []*TreeNode{root}
	rets := [][]int{}
	newRes := [][]int{}
	for len(stack) != 0 {
		tmpStack := []*TreeNode{}
		ret := []int{}
		for i := range stack {
			ret = append(ret, stack[i].Val)
			if stack[i].Left != nil {
				tmpStack = append(tmpStack, stack[i].Left)
			}
			if stack[i].Right != nil {
				tmpStack = append(tmpStack, stack[i].Right)
			}
		}
		rets = append(rets, ret)
		stack = tmpStack

	}
	for key, item := range rets {
		if key%2 == 0 {
			newRes = append(newRes, rets[key])
		} else {
			newItem := []int{}
			for eachkey, _ := range item {
				newItem = append(newItem, item[len(item)-1-eachkey])
			}
			newRes = append(newRes, newItem)

		}
	}
	return newRes
}
