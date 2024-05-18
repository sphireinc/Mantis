package datastructures

import (
	"errors"
)

// source & credit: https://ieftimov.com/post/golang-datastructures-trees/

// PKTreeNode is a node within our tree
type PKTreeNode struct {
	id       string
	name     string
	value    any
	parent   *PKTreeNode
	children []*PKTreeNode
}

// NewPKTree creates a new tree
func NewPKTree(id string, name string, value any) *PKTreeNode {
	return &PKTreeNode{
		id:    id,
		name:  name,
		value: value,
	}
}

// Insert a new node into our tree under a given parent
func (tree *PKTreeNode) Insert(id string, name string, value any, parentID string) (bool, error) {
	parent := tree.FindByID(parentID)
	if parent == nil {
		return false, errors.New("parent is nil")
	}

	for _, child := range parent.children {
		if child.id == id {
			return false, errors.New("duplicate child id")
		}
	}

	child := &PKTreeNode{
		id:     id,
		name:   name,
		value:  value,
		parent: parent,
	}

	parent.children = append(parent.children, child)
	return true, nil
}

// FindByID finds a given node by its ID (BFS)
func (tree *PKTreeNode) FindByID(id string) *PKTreeNode {
	queue := append(make([]*PKTreeNode, 0), tree)
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
func (tree *PKTreeNode) FindByIDDFS(id string) *PKTreeNode {
	if tree.id == id {
		return tree
	}

	if len(tree.children) <= 0 {
		return nil
	}

	for _, child := range tree.children {
		return child.FindByIDDFS(id)
	}
	return nil
}

// Remove a given node from our tree
func (tree *PKTreeNode) Remove(node *PKTreeNode) {
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
