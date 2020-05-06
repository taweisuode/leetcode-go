package LeetCode

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Code49() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagramsV3(strs))
}

/**
给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。

示例:

输入: ["eat", "tea", "tan", "ate", "nat", "bat"],
输出:
[
  ["ate","eat","tea"],
  ["nat","tan"],
  ["bat"]
]
说明：

所有输入均为小写字母。
不考虑答案输出的顺序。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/group-anagrams
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
```
*/

type HashArr struct {
	Arr map[string]int
}

//  正常解法：每个字符串按照字符排序，然后 string -> index 做hash  然后拼成[][]string
func groupAnagrams(strs []string) [][]string {
	hashArr := &HashArr{
		Arr: make(map[string]int, 0),
	}

	res := make([][]string, 0)
	for i := 0; i < len(strs); i++ {
		getStr := getOrderStr(strs[i])
		if index, ok := hashArr.Arr[getStr]; ok {
			res[index] = append(res[index], strs[i])
		} else {
			hashArr.Arr[getStr] = len(res)
			res = append(res, []string{strs[i]})
		}
	}
	return res
}

func getOrderStr(str string) string {
	strs := strings.Split(str, "")
	sort.Strings(strs)
	return strings.Join(strs, "")
}

//每个字符串映射为1到26个数字的次数 用#号连接  比如 abb 为 #1#2#0#0#..  那么 bba 也是 #1#2#0#0#...

func groupAnagramsV2(strs []string) [][]string {
	res := make([][]string, 0)
	hashArr := make(map[string]int, 0)
	for _, item := range strs {
		s := getChange(item)
		if index, ok := hashArr[s]; ok {
			res[index] = append(res[index], item)
		} else {
			hashArr[s] = len(res)
			res = append(res, []string{item})
		}
	}
	return res
}

func getChange(str string) string {
	nums := make([]int, 26)
	res := "#"
	for _, item := range []byte(str) {
		n := item - 'a'
		nums[n]++

	}
	for _, v := range nums {
		res = res + strconv.Itoa(v) + "#"
	}
	return res
}

//解法三 利用数学质数 来将每个str  映射成n个 质数和
// 比如  abc 为 2*3*5  bac 为 3*2*5  这2这的hash值是一致的

func groupAnagramsV3(strs []string) [][]string {
	res := make([][]string, 0)
	strMap := make(map[int]int, 0)
	for _, item := range strs {
		key := getHashKey(item)
		if index, ok := strMap[key]; ok {
			res[index] = append(res[index], item)
		} else {
			strMap[key] = len(res)
			res = append(res, []string{item})
		}
	}
	return res
}

func getHashKey(str string) int {
	var numArr = map[byte]int{
		'a': 2, 'b': 3, 'c': 5, 'd': 7, 'e': 11, 'f': 13, 'g': 17, 'h': 19,
		'i': 23, 'j': 29, 'k': 31, 'l': 37, 'm': 41, 'n': 43, 'o': 47, 'p': 53, 'q': 59, 'r': 61,
		's': 67, 't': 71, 'u': 73, 'v': 79, 'w': 83, 'x': 89, 'y': 97, 'z': 101,
	}

	res := 1
	for _, item := range []byte(str) {
		res *= numArr[item]
	}

	return res
}
