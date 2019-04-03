package LeetCode

import (
	"fmt"
	"strings"
)

func Code71() {
	path := "/home//foo/"
	fmt.Println(simplifyPath(path))
}

/**
以 Unix 风格给出一个文件的绝对路径，你需要简化它。或者换句话说，将其转换为规范路径。

在 Unix 风格的文件系统中，一个点（.）表示当前目录本身；此外，两个点 （..） 表示将目录切换到上一级（指向父目录）；两者都可以是复杂相对路径的组成部分。更多信息请参阅：Linux / Unix中的绝对路径 vs 相对路径

请注意，返回的规范路径必须始终以斜杠 / 开头，并且两个目录名之间必须只有一个斜杠 /。最后一个目录名（如果存在）不能以 / 结尾。此外，规范路径必须是表示绝对路径的最短字符串。



示例 1：

输入："/home/"
输出："/home"
解释：注意，最后一个目录名后面没有斜杠。
*/

func simplifyPath(path string) string {
	if path[0] != '/' || path == "" {
		return ""
	}
	stack := &Stack{}
	strArr := strings.Split(path, "/")
	for i := 1; i < len(strArr); i++ {
		if strArr[i] == ".." {
			if !StackEmpty(stack) {
				Stackpop(stack)
			}
		} else {
			if strArr[i] == "." || strArr[i] == "" {
				continue
			} else {
				Stackpush(strArr[i], stack)
			}
		}
	}
	resArr := make([]string, 0)
	if StackEmpty(stack) {
		return "/"
	}
	for !StackEmpty(stack) {
		stackNode := Stackpop(stack).(string)
		resArr = append(resArr, stackNode)
	}
	resStr := "/"
	for i := len(resArr) - 1; i >= 0; i-- {
		resStr += resArr[i] + "/"
	}
	return resStr[0 : len(resStr)-1]
}
