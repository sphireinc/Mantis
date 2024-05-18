package datastructures

import (
	"testing"
)

func TestNewPKTree(t *testing.T) {
	tree := NewPKTree("1", "root", "rootValue")

	if tree.id != "1" || tree.name != "root" || tree.value != "rootValue" {
		t.Errorf("NewPKTree failed, expected id: '1', name: 'root', value: 'rootValue', got id: '%s', name: '%s', value: '%v'", tree.id, tree.name, tree.value)
	}
}

func TestPKTInsert(t *testing.T) {
	tree := NewPKTree("1", "root", "rootValue")

	// Insert a child node
	success, err := tree.Insert("2", "child1", "childValue1", "1")
	if !success || err != nil {
		t.Errorf("Insert failed, expected success, got error: %v", err)
	}

	// Try inserting a duplicate node
	success, err = tree.Insert("2", "child1", "childValue1", "1")
	if success || err == nil || err.Error() != "duplicate child id" {
		t.Errorf("Insert failed, expected duplicate child id error, got success: %v, error: %v", success, err)
	}

	// Try inserting a node with a non-existent parent
	success, err = tree.Insert("3", "child2", "childValue2", "99")
	if success || err == nil || err.Error() != "parent is nil" {
		t.Errorf("Insert failed, expected parent is nil error, got success: %v, error: %v", success, err)
	}
}

func TestFindByID(t *testing.T) {
	tree := NewPKTree("1", "root", "rootValue")
	_, _ = tree.Insert("2", "child1", "childValue1", "1")
	_, _ = tree.Insert("3", "child2", "childValue2", "1")

	// Find existing node
	node := tree.FindByID("2")
	if node == nil || node.id != "2" {
		t.Errorf("FindByID failed, expected id: '2', got: %v", node)
	}

	// Find non-existent node
	node = tree.FindByID("99")
	if node != nil {
		t.Errorf("FindByID failed, expected nil, got: %v", node)
	}
}

func TestFindByIDDFS(t *testing.T) {
	tree := NewPKTree("1", "root", "rootValue")
	_, _ = tree.Insert("2", "child1", "childValue1", "1")
	_, _ = tree.Insert("3", "child2", "childValue2", "1")
	_, _ = tree.Insert("4", "child2", "childValue2", "1")

	// Find existing node
	node := tree.FindByIDDFS("2")
	if node == nil || node.id != "2" {
		t.Errorf("FindByIDDFS failed, expected id: '2', got: %v", node)
	}

	// Find non-existent node
	node = tree.FindByIDDFS("99")
	if node != nil {
		t.Errorf("FindByIDDFS failed, expected nil, got: %v", node)
	}
}

func TestPKTRemove(t *testing.T) {
	tree := NewPKTree("1", "root", "rootValue")
	_, _ = tree.Insert("2", "child1", "childValue1", "1")
	_, _ = tree.Insert("3", "child2", "childValue2", "1")
	childNode := tree.FindByID("2")

	// Remove existing node
	tree.Remove(childNode)
	node := tree.FindByID("2")
	if node != nil {
		t.Errorf("Remove failed, expected node to be removed, got: %v", node)
	}

	// Check if children of the removed node are orphaned
	if len(childNode.children) != 0 {
		t.Errorf("Remove failed, expected children to be nil, got: %v", childNode.children)
	}
}
