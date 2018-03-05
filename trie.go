// You can edit this code!
// Click here and start typing.
package main
/* Aaron Pitman */

import (
	"fmt"
)

type trie struct {
	root *node
}

type node struct {
	isWord bool
	next [26]*node
}

func New() *trie {
	return &trie{nil}
}

func (trieNode *trie) contains(word string) bool {
	if trieNode == nil {
		return false
	}
	current := trieNode.root
	for i := 0; i < len(word); i++ {
		if current.next[getIndex(word[i])] == nil {
			return false
		} else {
			current = current.next[getIndex(word[i])]
		}
	}
	return current.isWord
}

func (trieNode *trie) insert(word string) {
	if trieNode.root == nil {
		trieNode.root = newNode()
	}
	current := trieNode.root
	for i := 0; i < len(word); i++ {
		if current.next[getIndex(word[i])] != nil {
			current = current.next[getIndex(word[i])]
		} else {
			current.next[getIndex(word[i])] = newNode()
			current = current.next[getIndex(word[i])]
		}
	}
	current.isWord = true
}

func newNode() *node {
	var next [26]*node
	return &node{false,next}
}

func getIndex(charVal byte) int {
	return int(charVal - 'a')
}

func main() {
	fmt.Println("Hey, this is WIP. Stop judging.")
	trie := New()
	trie.insert("hello")
	fmt.Println(trie.contains("hello"))
}
