package datastructures

type TrieNode struct {
	value    byte
	children [26]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func newTrieNode(val byte) *TrieNode {
	return &TrieNode{
		value: val,
	}
}

func (t *Trie) Insert(word string) {
	if t.root == nil {
		t.root = newTrieNode('/')
	}

	currTrieNode := t.root
	for i := 0; i < len(word); i++ {
		c := word[i] - 'a'
		if currTrieNode.children[c] == nil {
			currTrieNode.children[c] = newTrieNode(c + 'a')
		}
		currTrieNode = currTrieNode.children[c]
	}

	currTrieNode.isEnd = true
}

func (t *Trie) Search(word string) bool {
	currTrieNode := t.root
	for i := 0; i < len(word); i++ {
		c := word[i] - 'a'
		if currTrieNode.children[c] == nil {
			return false
		}
		currTrieNode = currTrieNode.children[c]
	}
	return currTrieNode.isEnd
}
