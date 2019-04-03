package LeetCode

func Code105() {
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

/**
根据一棵树的前序遍历与中序遍历构造二叉树。

注意:
你可以假设树中没有重复的元素。

例如，给出

前序遍历 preorder = [3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]
返回如下的二叉树：

    3
   / \
  9  20
    /  \
   15   7
*/

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) < 1 {
		return nil
	}
	root := &TreeNode{}
	for i := 0; i < len(preorder); i++ {
		if preorder[0] == inorder[i] {
			root.Left = buildTree(preorder[1:i+1], inorder[0:i])
			root.Right = buildTree(preorder[i+2:], inorder[i+1:])
		}
	}
	return root
}
