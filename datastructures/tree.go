package datastructures

type Node struct {
	value int
	left  *Node
	right *Node
}

type Tree struct {
	root *Node
}

func (t *Tree) Insert(val int) {
	if t.root == nil {
		t.root = &Node{value: val}
		return
	}

	currNode := t.root
	for {
		if val < currNode.value {
			if currNode.left == nil {
				currNode.left = &Node{value: val}
				return
			}
			currNode = currNode.left
		} else {
			if currNode.right == nil {
				currNode.right = &Node{value: val}
				return
			}
			currNode = currNode.right
		}
	}
}

func (t *Tree) Search(val int) bool {
	currNode := t.root
	for currNode != nil {
		if val == currNode.value {
			return true
		} else if val < currNode.value {
			currNode = currNode.left
		} else {
			currNode = currNode.right
		}
	}
	return false
}
