package LeetCode

import (
	"fmt"
)

func Code17() {
	fmt.Println(letterCombinations("234"))
}

/**
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。



示例:

输入："23"
输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
说明:
尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
```
*/

func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}
	res := make([]string, 0)
	result := dfs_17(digits, 0, nil, &res)
	return *result
}

func dfs_17(digits string, index int, next []byte, res *[]string) *[]string {
	if index == len(digits) {
		*res = append(*res, string(next))
		return res
	}
	keyMap := make(map[byte][]byte, 8)
	keyMap['2'] = []byte{'a', 'b', 'c'}
	keyMap['3'] = []byte{'d', 'e', 'f'}
	keyMap['4'] = []byte{'g', 'h', 'i'}
	keyMap['5'] = []byte{'j', 'k', 'l'}
	keyMap['6'] = []byte{'m', 'n', 'o'}
	keyMap['7'] = []byte{'p', 'q', 'r', 's'}
	keyMap['8'] = []byte{'t', 'u', 'v'}
	keyMap['9'] = []byte{'w', 'x', 'y', 'z'}

	for i := 0; i < len(keyMap[digits[index]]); i++ {
		cur := append(next, keyMap[digits[index]][i])
		dfs_17(digits, index+1, cur, res)
	}
	return res
}
