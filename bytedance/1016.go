package bytedance

import "fmt"

func Code1016() {
	s1 := "abcdxabcde"
	s2 := "abcdeabcdx"
	fmt.Println(checkInclusion(s1, s2))

}

/**

给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的排列。

换句话说，第一个字符串的排列之一是第二个字符串的子串。

示例1:

输入: s1 = "ab" s2 = "eidbaooo"
输出: True
解释: s2 包含 s1 的排列之一 ("ba").


示例2:

输入: s1= "ab" s2 = "eidboaoo"
输出: False


注意：

输入的字符串只包含小写字母
两个字符串的长度都在 [1, 10,000] 之间
```
*/

func checkInclusion(s1 string, s2 string) bool {
	len1 := len(s1)
	len2 := len(s2)
	if len1 > len2 {
		return false
	}
	for i := 0; i <= len2-len1; i++ {
		if compare(s1, s2[i:i+len1]) {
			return true
		}
	}
	return false
}
func compare(s1 string, s2 string) bool {
	ca := make([]int, 26)
	cb := make([]int, 26)

	for _, item := range s1 {
		ca[item-'a']++
	}
	for _, item := range s2 {
		cb[item-'a']++
	}

	fmt.Println(ca)
	fmt.Println(cb)
	for i := 0; i < 26; i++ {
		if ca[i] != cb[i] {
			return false
		}
	}
	return true
}
