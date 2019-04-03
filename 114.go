package LeetCode

func Code114() {
	head := InitTree()
	head = &TreeNode{1, nil, nil}
	head.Left = &TreeNode{4, nil, nil}
	head.Right = &TreeNode{8, nil, nil}
	head.Left.Left = &TreeNode{11, nil, nil}
	head.Right.Left = &TreeNode{13, nil, nil}
	head.Right.Right = &TreeNode{4, nil, nil}
	head.Left.Left.Left = &TreeNode{7, nil, nil}
	head.Left.Left.Right = &TreeNode{2, nil, nil}
	head.Right.Right.Left = &TreeNode{5, nil, nil}
	head.Right.Right.Right = &TreeNode{1, nil, nil}
	flatten(head)
}

/**
给定一个二叉树，原地将它展开为链表。

例如，给定二叉树

    1
   / \
  2   5
 / \   \
3   4   6
将其展开为：

1
 \
  2
   \
    3
     \
      4
       \
        5
         \
          6
*/
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	if root.Left != nil {
		flatten(root.Left)
	}
	if root.Right != nil {
		flatten(root.Right)
	}
	temp := root.Right
	root.Left, root.Right = root.Right, root.Left
	for root.Right != nil {
		root = root.Right
	}
	root.Right = temp
}
