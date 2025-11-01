package lesson_4

import (
	"fmt"
)

type Tree struct {
	data  int
	left  *Tree
	right *Tree
}

func (t *Tree) Insert(data int) *Tree {
	if t == nil {
		return createNewNode(data)
	}

	if data < t.data {
		t.left = t.left.Insert(data)
	} else if data > t.data {
		t.right = t.right.Insert(data)
	}

	return t
}

func (tree *Tree) Contains(value int) bool {
	if tree == nil {
		return false
	}

	if value > tree.data {
		return tree.right.Contains(value)
	} else if value < tree.data {
		return tree.left.Contains(value)
	} else {
		return true
	}
}

func (tree *Tree) Delete(value int) *Tree {
	if tree == nil {
		return nil
	}

	if value > tree.data {
		tree.right = tree.right.Delete(value)
	} else if value < tree.data {
		tree.left = tree.left.Delete(value)
	} else {
		if tree.right == nil {
			return tree.left
		} else if tree.left == nil {
			return tree.right
		} else {
			successor := tree.right.findMin()
			tree.data = successor.data
			tree.right = tree.right.Delete(successor.data)
		}
	}

	return tree
}

func (tree *Tree) findMin() *Tree {
	current := tree
	for current.left != nil {
		current = current.left
	}
	return current
}

func (tree *Tree) InorderTraversal(action func(int, int)) {
	if tree == nil {
		return
	}

	tree.left.InorderTraversal(action)
	action(tree.data, 0)
	tree.right.InorderTraversal(action)
}

func (tree *Tree) Count() int {
	if tree == nil {
		return 0
	}

	return 1 + tree.left.Count() + tree.right.Count()
}

func (n *Tree) PrintNode() {
	if n == nil {
		fmt.Println("Receiver is nil. Skipping output.")
		return
	}
	fmt.Printf("Node{\n data: %d\n left: %v\n right: %v\n}\n\n", n.data, n.left, n.right)
}

func createNewNode(data int) *Tree {
	return &Tree{
		data:  data,
		left:  nil,
		right: nil,
	}
}
