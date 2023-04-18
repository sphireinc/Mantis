package datastructures

// TrieNode represents a node within our Trie
type TrieNode struct {
	value    byte
	children [26]*TrieNode
	isEnd    bool
}

// Trie is our core Trie
type Trie struct {
	root *TrieNode
}

// newTrieNode creates a new Trie
func newTrieNode(val byte) *TrieNode {
	return &TrieNode{
		value: val,
	}
}

// Insert a word onto our Trie
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

// Search for a word within our Trie
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
