package LeetCode

import (
	"fmt"
	"strings"
)

func Code331() {
	//preOrder := "9,3,4,#,#,1,#,#,2,#,6,#,#"
	preOrder1 := "#,#,3,5,#"
	fmt.Println(isValidSerialization(preOrder1))
}

/**
序列化二叉树的一种方法是使用前序遍历。当我们遇到一个非空节点时，我们可以记录下这个节点的值。如果它是一个空节点，我们可以使用一个标记值记录，例如 #。

     _9_
    /   \
   3     2
  / \   / \
 4   1  #  6
/ \ / \   / \
# # # #   # #
例如，上面的二叉树可以被序列化为字符串 "9,3,4,#,#,1,#,#,2,#,6,#,#"，其中 # 代表一个空节点。

给定一串以逗号分隔的序列，验证它是否是正确的二叉树的前序序列化。编写一个在不重构树的条件下的可行算法。

每个以逗号分隔的字符或为一个整数或为一个表示 null 指针的 '#' 。

你可以认为输入格式总是有效的，例如它永远不会包含两个连续的逗号，比如 "1,,3" 。

示例 1:

输入: "9,3,4,#,#,1,#,#,2,#,6,#,#"
输出: true
示例 2:

输入: "1,#"
输出: false
示例 3:

输入: "9,#,#,1"
输出: false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/verify-preorder-serialization-of-a-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func isValidSerialization(preorder string) bool {
	leaves := 0
	node := 0
	list := strings.Split(preorder, ",")
	for i := range list {
		if list[i] == "#" {
			leaves++
		} else {
			node++
		}
		if leaves > node+1 {
			return false
		}
		if leaves == node+1 && i < len(list)-1 {
			return false
		}
	}
	if leaves == node+1 {
		return true
	}
	return false
}
