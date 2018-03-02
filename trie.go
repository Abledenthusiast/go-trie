/* Aaron Pitman */

import (
	"fmt"
)

type trie struct {
	root *node
	children []trie
}

type node struct {
	key string
	isWord bool
	next [26]*node
}

func (trieNode trie) insert(word string) *node {
	if trieNode is nil {
		trieNode.root = newNode()
	}
	current := trieNode.root
	for i := 0; i < len(word); i++ {
		if(current.next[word[i]] != nil){
			current = current.next
		}
		else {
			current.next[word[i]] = newNode()
			current = current.next[word[i]]
		}
	}
	current.isWord = true
}

func newNode() *node {
	var next [256]*node
	return &node{nil,nil,next}
}

func main() {
	fmt.Println("Hey, this is WIP. Stop judging.")
}