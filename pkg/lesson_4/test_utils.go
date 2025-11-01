package lesson_4

// Test utility functions shared across test files

// collectInorder performs inorder traversal and returns values as slice
func collectInorder(tree *Tree) []int {
	var result []int
	var inorder func(*Tree)
	inorder = func(n *Tree) {
		if n == nil {
			return
		}
		inorder(n.left)
		result = append(result, n.data)
		inorder(n.right)
	}
	inorder(tree)
	return result
}

// slicesEqual compares two int slices for equality
func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// isBST checks if tree maintains BST property (left < parent < right)
func isBST(node *Tree, min, max *int) bool {
	if node == nil {
		return true
	}

	if min != nil && node.data <= *min {
		return false
	}
	if max != nil && node.data >= *max {
		return false
	}

	return isBST(node.left, min, &node.data) && isBST(node.right, &node.data, max)
}
