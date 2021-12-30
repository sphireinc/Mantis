package data

import (
	"errors"
)

// source & credit: https://ieftimov.com/post/golang-datastructures-trees/

type treeNode struct {
	id       string
	name     string
	value    any
	parent   *treeNode
	children []*treeNode
}

// NewTree creates a new tree
func NewTree(id string, name string, value any) *treeNode {
	return &treeNode{
		id:    id,
		name:  name,
		value: value,
	}
}

// Insert a new node into our tree under a given parent
func (tree *treeNode) Insert(id string, name string, value any, parentId string) (bool, error) {
	parent := tree.FindById(parentId)
	if parent == nil {
		return false, errors.New("parent is nil")
	}

	for _, child := range parent.children {
		if child.id == id {
			return false, errors.New("duplicate child id")
		}
	}

	child := &treeNode{
		id:     id,
		name:   name,
		value:  value,
		parent: parent,
	}

	parent.children = append(parent.children, child)
	return true, nil
}

// FindById finds a given node by its ID (BFS)
func (tree *treeNode) FindById(id string) *treeNode {
	queue := append(make([]*treeNode, 0), tree)
	for len(queue) > 0 {
		next := queue[0]
		queue = queue[1:]
		if next.id == id {
			return next
		}
		if len(next.children) > 0 {
			for _, child := range next.children {
				queue = append(queue, child)
			}
		}
	}
	return nil
}

// FindByIdDFS finds a given node by its ID (DFS)
func (tree *treeNode) FindByIdDFS(id string) *treeNode {
	if tree.id == id {
		return tree
	}

	if len(tree.children) > 0 {
		for _, child := range tree.children {
			tree = child.FindByIdDFS(id)
		}
	}
	return tree
}

// Remove a given node from our tree
func (tree *treeNode) Remove(node *treeNode) {
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
