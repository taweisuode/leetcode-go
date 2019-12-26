package LeetCode

import (
	"fmt"
)

func Code212() {
	/*	board := [][]byte{
			{'o', 'a', 'a', 'n'},
			{'e', 't', 'a', 'e'},
			{'i', 'h', 'k', 'r'},
			{'i', 'f', 'l', 'v'},
		}
		words := []string{"oath", "pea", "eat", "rain"}*/
	board1 := [][]byte{
		{'a', 'a', 'a', 'a'},
		{'a', 'a', 'a', 'a'},
		{'a', 'a', 'a', 'a'},
	}
	words1 := []string{"aaaaaaaaaaaa", "aaaaaaaaaaaaa", "aaaaaaaaaaab"}
	fmt.Println(findWords(board1, words1))
}

/**
给定一个二维网格 board 和一个字典中的单词列表 words，找出所有同时在二维网格和字典中出现的单词。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母在一个单词中不允许被重复使用。

示例:

输入:
words = ["oath","pea","eat","rain"] and board =
[
  ['o','a','a','n'],
  ['e','t','a','e'],
  ['i','h','k','r'],
  ['i','f','l','v']
]

输出: ["eat","oath"]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/word-search-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func findWords(board [][]byte, words []string) []string {
	res := make(map[string]int, 0)
	for i := range words {
		boardMap := initMap(board)
		for x := 0; x < len(board); x++ {
			for y := 0; y < len(board[0]); y++ {
				if words[i][0] == board[x][y] {
					boardMap[x][y] = 1
					cur := dfs_212(board, x, y, words[i], words[i][1:], boardMap)
					if cur != "" {
						if _, ok := res[cur]; !ok {
							res[cur] = 1
						}
					} else {
						boardMap[x][y] = 0
					}
				}
			}
		}
	}

	var result []string
	for key := range res {
		result = append(result, key)
	}
	return result
}
func initMap(board [][]byte) [][]int {
	boardMap := make([][]int, len(board))
	for i := range board {
		boardMap[i] = make([]int, len(board[0]))
	}
	return boardMap
}
func dfs_212(board [][]byte, x int, y int, originWord string, word string, boardMap [][]int) string {
	if word == "" {
		return originWord
	}

	direction := [][]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}

	for key := range direction {
		nextX := x + direction[key][0]
		nextY := y + direction[key][1]

		for i := 0; i < len(board); i++ {
			for j := 0; j < len(board[0]); j++ {

				if nextX >= 0 && nextX < len(board) && nextY >= 0 && nextY < len(board[0]) {
					if board[nextX][nextY] == word[0] && boardMap[nextX][nextY] == 0 {
						boardMap[nextX][nextY] = 1
						if dfs_212(board, nextX, nextY, originWord, word[1:], boardMap) != "" {
							return originWord
						} else {
							boardMap[nextX][nextY] = 0
							continue
						}
					}
				}
			}
		}
	}
	return ""
}
