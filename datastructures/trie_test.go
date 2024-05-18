package datastructures

import (
	"testing"
)

func TestTrie_Insert(t *testing.T) {
	trie := &Trie{}

	// Insert a single word
	trie.Insert("hello")
	if !trie.Search("hello") {
		t.Errorf("Insert failed, expected to find 'hello' after insertion")
	}

	// Insert multiple words
	words := []string{"world", "trie", "algorithm", "data"}
	for _, word := range words {
		trie.Insert(word)
	}

	// Verify each word can be found
	for _, word := range words {
		if !trie.Search(word) {
			t.Errorf("Insert failed, expected to find '%s' after insertion", word)
		}
	}
}

func TestTrie_Search(t *testing.T) {
	trie := &Trie{}
	words := []string{"hello", "world", "trie", "algorithm", "data"}

	// Insert words into the trie
	for _, word := range words {
		trie.Insert(word)
	}

	// Search for existing words
	for _, word := range words {
		if !trie.Search(word) {
			t.Errorf("Search failed, expected to find '%s'", word)
		}
	}

	// Search for non-existent words
	nonExistentWords := []string{"helloo", "word", "tried", "algorithms", "datum"}
	for _, word := range nonExistentWords {
		if trie.Search(word) {
			t.Errorf("Search failed, did not expect to find '%s'", word)
		}
	}
}

func TestTrie_PrefixSearch(t *testing.T) {
	trie := &Trie{}

	// Insert words with common prefixes
	words := []string{"app", "apple", "application", "apt", "apex"}
	for _, word := range words {
		trie.Insert(word)
	}

	// Verify full words can be found
	for _, word := range words {
		if !trie.Search(word) {
			t.Errorf("Prefix search failed, expected to find '%s'", word)
		}
	}

	// Verify prefixes are not found as complete words
	prefixes := []string{"ap", "appl", "applic"}
	for _, prefix := range prefixes {
		if trie.Search(prefix) {
			t.Errorf("Prefix search failed, did not expect to find '%s' as a complete word", prefix)
		}
	}
}
