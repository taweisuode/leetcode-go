package LeetCode

import (
	"fmt"
)

func Code127() {
	beginWord := "hog"
	endWord := "cog"
	wordList := []string{"cog"}
	num := ladderLength(beginWord, endWord, wordList)
	fmt.Println(num)
}

/**
给定两个单词（beginWord 和 endWord）和一个字典，找到从 beginWord 到 endWord 的最短转换序列的长度。转换需遵循如下规则：

每次转换只能改变一个字母。
转换过程中的中间单词必须是字典中的单词。
说明:

如果不存在这样的转换序列，返回 0。
所有单词具有相同的长度。
所有单词只由小写字母组成。
字典中不存在重复的单词。
你可以假设 beginWord 和 endWord 是非空的，且二者不相同。
示例 1:

输入:
beginWord = "hit",
endWord = "cog",
wordList = ["hot","dot","dog","lot","log","cog"]

输出: 5

解释: 一个最短转换序列是 "hit" -> "hot" -> "dot" -> "dog" -> "cog",
     返回它的长度 5。
示例 2:

输入:
beginWord = "hit"
endWord = "cog"
wordList = ["hot","dot","dog","lot","log"]

输出: 0

解释: endWord "cog" 不在字典中，所以无法进行转换。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/word-ladder
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordMap := make(map[string]int, 0)
	for _, item := range wordList {
		wordMap[item] = 0
	}
	res := bfs(beginWord, endWord, wordMap)
	return res
}

func bfs(search string, endWord string, wordList map[string]int) int {
	queue := &Queue1{}
	inQueue1(search, queue)
	step := 0
	for !queueEmpty1(queue) {
		step++

		size := len(queue.item)

		for i := 0; i < size; i++ {
			node := deQueue1(queue)
			if node == endWord {
				return step
			}
			getArr := findSimilar(node, wordList)
			for _, item := range getArr {
				inQueue1(item, queue)
			}
		}

	}
	return 0
}

//find string is one-place-different with the search string
func findSimilar(search string, wordList map[string]int) []string {
	var strArr []string
	for key := range wordList {
		if wordList[key] == 0 {
			if judgeIsSimilar(search, key) {
				strArr = append(strArr, key)
				wordList[key] = 1
			}
		}
	}
	return strArr
}

func judgeIsSimilar(str1 string, str2 string) bool {
	//englishWord := "abcdefghijklmnopqrstuvwxyz"
	differentCount := 0
	for index := range str1 {
		if str1[index] != str2[index] {
			differentCount++
		}
		if differentCount > 1 {
			return false
		}
	}
	return differentCount == 1
}

type Queue1 struct {
	item []string
}

func inQueue1(p string, queue *Queue1) {
	queue.item = append(queue.item, p)
}

func deQueue1(queue *Queue1) string {
	length := len(queue.item)
	deleteItem := queue.item[0]
	queue.item = queue.item[1:length]
	return deleteItem
}
func queueEmpty1(queue *Queue1) bool {
	return len(queue.item) == 0
}
