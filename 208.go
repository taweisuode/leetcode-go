package LeetCode

import (
	"fmt"
)

func Code208() {
	trie := Constructor()
	trie.Insert("hello")
	trie.Insert("world")

	res := trie.Search("hel")
	fmt.Println(res, trie)
}

/**
实现一个 Trie (前缀树)，包含 insert, search, 和 startsWith 这三个操作。

示例:

Trie trie = new Trie();

trie.insert("apple");
trie.search("apple");   // 返回 true
trie.search("app");     // 返回 false
trie.startsWith("app"); // 返回 true
trie.insert("app");
trie.search("app");     // 返回 true

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/implement-trie-prefix-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
type Trie struct {
	IsWord   bool
	Children map[rune]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{Children: make(map[rune]*Trie), IsWord: false}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	root := this
	for _, item := range word {
		n := item - 'a'
		if root.Children[n] == nil {
			create := &Trie{Children: make(map[int32]*Trie), IsWord: false}
			root.Children[n] = create
		}
		root = root.Children[n]

	}
	root.IsWord = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	root := this

	for _, item := range word {
		n := item - 'a'
		if root.Children[n] == nil {
			return false
		}
		root = root.Children[n]
	}
	return root.IsWord
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	root := this
	for _, item := range prefix {
		n := item - 'a'
		if root.Children[n] == nil {
			return false
		}
		root = root.Children[n]
	}
	return root != nil
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
