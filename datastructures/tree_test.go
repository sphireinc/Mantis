package datastructures

import (
	"testing"
)

func TestTree_Insert(t *testing.T) {
	tree := &Tree{}

	// Insert a single value
	tree.Insert(10)
	if tree.root == nil || tree.root.value != 10 {
		t.Errorf("Insert failed, expected root value 10, got %v", tree.root)
	}

	// Insert values to form a balanced tree
	tree.Insert(5)
	tree.Insert(15)
	if tree.root.left == nil || tree.root.left.value != 5 {
		t.Errorf("Insert failed, expected left child value 5, got %v", tree.root.left)
	}
	if tree.root.right == nil || tree.root.right.value != 15 {
		t.Errorf("Insert failed, expected right child value 15, got %v", tree.root.right)
	}

	// Insert additional values
	tree.Insert(2)
	tree.Insert(7)
	tree.Insert(12)
	tree.Insert(17)

	if tree.root.left.left == nil || tree.root.left.left.value != 2 {
		t.Errorf("Insert failed, expected left-left child value 2, got %v", tree.root.left.left)
	}
	if tree.root.left.right == nil || tree.root.left.right.value != 7 {
		t.Errorf("Insert failed, expected left-right child value 7, got %v", tree.root.left.right)
	}
	if tree.root.right.left == nil || tree.root.right.left.value != 12 {
		t.Errorf("Insert failed, expected right-left child value 12, got %v", tree.root.right.left)
	}
	if tree.root.right.right == nil || tree.root.right.right.value != 17 {
		t.Errorf("Insert failed, expected right-right child value 17, got %v", tree.root.right.right)
	}
}

func TestTree_Search(t *testing.T) {
	tree := &Tree{}

	// Insert values
	values := []int{10, 5, 15, 2, 7, 12, 17}
	for _, val := range values {
		tree.Insert(val)
	}

	// Search for existing values
	for _, val := range values {
		found := tree.Search(val)
		if !found {
			t.Errorf("Search failed, expected to find value %d", val)
		}
	}

	// Search for non-existent values
	nonExistentValues := []int{1, 3, 6, 8, 11, 13, 16, 18}
	for _, val := range nonExistentValues {
		found := tree.Search(val)
		if found {
			t.Errorf("Search failed, did not expect to find value %d", val)
		}
	}
}
