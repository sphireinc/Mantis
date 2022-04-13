package data

import (
	"errors"
)

// source & credit: https://ieftimov.com/post/golang-datastructures-trees/

// TreeNode is a node within our tree
type TreeNode struct {
	id       string
	name     string
	value    any
	parent   *TreeNode
	children []*TreeNode
}

// NewTree creates a new tree
func NewTree(id string, name string, value any) *TreeNode {
	return &TreeNode{
		id:    id,
		name:  name,
		value: value,
	}
}

// Insert a new node into our tree under a given parent
func (tree *TreeNode) Insert(id string, name string, value any, parentID string) (bool, error) {
	parent := tree.FindByID(parentID)
	if parent == nil {
		return false, errors.New("parent is nil")
	}

	for _, child := range parent.children {
		if child.id == id {
			return false, errors.New("duplicate child id")
		}
	}

	child := &TreeNode{
		id:     id,
		name:   name,
		value:  value,
		parent: parent,
	}

	parent.children = append(parent.children, child)
	return true, nil
}

// FindByID finds a given node by its ID (BFS)
func (tree *TreeNode) FindByID(id string) *TreeNode {
	queue := append(make([]*TreeNode, 0), tree)
	for len(queue) > 0 {
		next := queue[0]
		queue = queue[1:]
		if next.id == id {
			return next
		}
		if len(next.children) > 0 {
			queue = append(queue, next.children...)
		}
	}
	return nil
}

// FindByIDDFS finds a given node by its ID (DFS)
func (tree *TreeNode) FindByIDDFS(id string) *TreeNode {
	if tree.id == id {
		return tree
	}

	if len(tree.children) <= 0 {
		return tree
	}
	for _, child := range tree.children {
		tree = child.FindByIDDFS(id)
	}
	return tree
}

// Remove a given node from our tree
func (tree *TreeNode) Remove(node *TreeNode) {
	for idx, sibling := range node.parent.children {
		if sibling == node {
			node.parent.children = append(
				node.parent.children[:idx],
				node.parent.children[idx+1:]...,
			)
		}
	}

	if len(node.children) != 0 {
		for _, child := range node.children {
			child.parent = nil
		}
		node.children = nil
	}
}
