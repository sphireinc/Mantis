package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func createNewTree() *TreeNode {
	tree := NewTree("1", "first", "hello world")
	_, _ = tree.Insert("2", "two", "hello world two", "1")
	_, _ = tree.Insert("3", "three", "hello world three", "2")
	_, _ = tree.Insert("4", "four", "hello world four", "2")
	_, _ = tree.Insert("5", "five", "hello world five", "4")
	_, _ = tree.Insert("6", "six", "hello world six", "5")
	_, _ = tree.Insert("7", "seven", "hello world seven", "5")
	return tree
}

func TestTree_NewTree(t *testing.T) {
	tree := NewTree("1", "first", "hello world")
	assert.NotNil(t, tree)
}

func TestTree_Insert(t *testing.T) {
	tree := createNewTree()

	_, err := tree.Insert("2", "two", "hello world two", "1")
	assert.Equal(t, "duplicate child id", err.Error())

	_, err = tree.Insert("2", "two", "hello world two", "10000")
	assert.Equal(t, "parent is nil", err.Error())

	_, _ = tree.Insert("8", "two", "hello world two", "1")
	assert.Equal(t, 2, len(tree.FindByID("1").children))

	_, _ = tree.Insert("9", "two", "hello world two", "1")
	assert.Equal(t, 3, len(tree.FindByID("1").children))
}

func TestTree_FindById(t *testing.T) {
	tree := createNewTree()
	assert.Equal(t, 2, len(tree.FindByID("5").children))
}

func TestTree_FindByIdDFS(t *testing.T) {
	tree := createNewTree()
	five := tree.FindByIDDFS("5")
	assert.Equal(t, 2, len(five.children))
	assert.Equal(t, "five", five.name)
}

func TestTree_Remove(t *testing.T) {
	tree := createNewTree()
	five := tree.FindByID("5")
	assert.NotNil(t, five)
	tree.Remove(five)
	fiveNew := tree.FindByID("5")
	assert.Nil(t, fiveNew)
}
