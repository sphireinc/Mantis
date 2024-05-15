package datastructures

// TreeNode defines our tree node
type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

// Tree defines our tree
type Tree struct {
	root *TreeNode
}

// Insert a value into our tree
func (t *Tree) Insert(val int) {
	if t.root == nil {
		t.root = &TreeNode{value: val}
		return
	}

	currNode := t.root
	for {
		if val < currNode.value {
			if currNode.left == nil {
				currNode.left = &TreeNode{value: val}
				return
			}
			currNode = currNode.left
		} else {
			if currNode.right == nil {
				currNode.right = &TreeNode{value: val}
				return
			}
			currNode = currNode.right
		}
	}
}

// Search for a value within our tree
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
