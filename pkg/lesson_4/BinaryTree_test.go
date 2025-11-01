package lesson_4

import (
	"testing"
)

func TestInsertNew(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		want   []int // expected inorder traversal
	}{
		{
			name:   "insert single value",
			values: []int{5},
			want:   []int{5},
		},
		{
			name:   "insert multiple values",
			values: []int{5, 3, 7, 1, 4, 6, 9},
			want:   []int{1, 3, 4, 5, 6, 7, 9},
		},
		{
			name:   "insert duplicates",
			values: []int{5, 3, 5, 3, 7},
			want:   []int{3, 5, 7},
		},
		{
			name:   "insert in ascending order",
			values: []int{1, 2, 3, 4, 5},
			want:   []int{1, 2, 3, 4, 5},
		},
		{
			name:   "insert in descending order",
			values: []int{5, 4, 3, 2, 1},
			want:   []int{1, 2, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var tree *Tree
			for _, val := range tt.values {
				tree = tree.Insert(val)
			}

			got := collectInorder(tree)
			if !slicesEqual(got, tt.want) {
				t.Errorf("InsertNew() got inorder = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_Contains(t *testing.T) {
	// Build test tree: 5, 3, 7, 1, 4, 6, 9
	var tree *Tree
	values := []int{5, 3, 7, 1, 4, 6, 9}
	for _, val := range values {
		tree = tree.Insert(val)
	}

	tests := []struct {
		name  string
		value int
		want  bool
	}{
		{"contains root", 5, true},
		{"contains left child", 3, true},
		{"contains right child", 7, true},
		{"contains leaf", 1, true},
		{"contains leaf", 9, true},
		{"does not contain smaller", 0, false},
		{"does not contain larger", 10, false},
		{"does not contain middle value", 2, false},
		{"does not contain middle value", 8, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tree.Contains(tt.value); got != tt.want {
				t.Errorf("Contains(%d) = %v, want %v", tt.value, got, tt.want)
			}
		})
	}
}

func TestTree_Contains_EmptyTree(t *testing.T) {
	var tree *Tree
	if tree.Contains(5) {
		t.Error("Contains on nil tree should return false")
	}
}

func TestTree_Delete(t *testing.T) {
	tests := []struct {
		name         string
		initialVals  []int
		deleteVal    int
		wantAfter    []int
		wantContains bool
	}{
		{
			name:         "delete leaf node",
			initialVals:  []int{5, 3, 7, 1, 9},
			deleteVal:    1,
			wantAfter:    []int{3, 5, 7, 9},
			wantContains: false,
		},
		{
			name:         "delete node with left child only",
			initialVals:  []int{5, 3, 1},
			deleteVal:    3,
			wantAfter:    []int{1, 5},
			wantContains: false,
		},
		{
			name:         "delete node with right child only",
			initialVals:  []int{5, 3, 4},
			deleteVal:    3,
			wantAfter:    []int{4, 5},
			wantContains: false,
		},
		{
			name:         "delete node with two children",
			initialVals:  []int{5, 3, 7, 1, 4, 6, 9},
			deleteVal:    7,
			wantAfter:    []int{1, 3, 4, 5, 6, 9},
			wantContains: false,
		},
		{
			name:         "delete root with two children",
			initialVals:  []int{5, 3, 7, 1, 4, 6, 9},
			deleteVal:    5,
			wantAfter:    []int{1, 3, 4, 6, 7, 9},
			wantContains: false,
		},
		{
			name:         "delete root leaf",
			initialVals:  []int{5},
			deleteVal:    5,
			wantAfter:    []int{},
			wantContains: false,
		},
		{
			name:         "delete non-existent value",
			initialVals:  []int{5, 3, 7},
			deleteVal:    10,
			wantAfter:    []int{3, 5, 7},
			wantContains: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var tree *Tree
			for _, val := range tt.initialVals {
				tree = tree.Insert(val)
			}

			tree = tree.Delete(tt.deleteVal)

			got := collectInorder(tree)
			if !slicesEqual(got, tt.wantAfter) {
				t.Errorf("After Delete(%d), inorder = %v, want %v", tt.deleteVal, got, tt.wantAfter)
			}

			if tree != nil && tree.Contains(tt.deleteVal) != tt.wantContains {
				t.Errorf("After Delete(%d), Contains(%d) = %v, want %v",
					tt.deleteVal, tt.deleteVal, tree.Contains(tt.deleteVal), tt.wantContains)
			}
		})
	}
}

func TestTree_Delete_EmptyTree(t *testing.T) {
	var tree *Tree
	result := tree.Delete(5)
	if result != nil {
		t.Error("Delete on nil tree should return nil")
	}
}

func TestInsertNew_BST_Property(t *testing.T) {
	// Verify BST property is maintained
	var tree *Tree
	values := []int{50, 30, 70, 20, 40, 60, 80, 10, 25, 35, 45}
	for _, val := range values {
		tree = tree.Insert(val)
	}

	if !isBST(tree, nil, nil) {
		t.Error("Tree does not maintain BST property after insertions")
	}
}

func TestDelete_BST_Property(t *testing.T) {
	// Verify BST property is maintained after deletions
	var tree *Tree
	values := []int{50, 30, 70, 20, 40, 60, 80, 10, 25, 35, 45}
	for _, val := range values {
		tree = tree.Insert(val)
	}

	// Delete various nodes
	deleteVals := []int{20, 30, 50}
	for _, val := range deleteVals {
		tree = tree.Delete(val)
		if !isBST(tree, nil, nil) {
			t.Errorf("Tree does not maintain BST property after deleting %d", val)
		}
	}
}

func TestTree_Count(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		want   int
	}{
		{
			name:   "empty tree",
			values: []int{},
			want:   0,
		},
		{
			name:   "single node",
			values: []int{5},
			want:   1,
		},
		{
			name:   "three nodes",
			values: []int{5, 3, 7},
			want:   3,
		},
		{
			name:   "seven nodes",
			values: []int{5, 3, 7, 1, 4, 6, 9},
			want:   7,
		},
		{
			name:   "with duplicates",
			values: []int{5, 3, 7, 3, 5, 1},
			want:   4, // duplicates not counted
		},
		{
			name:   "ascending order",
			values: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want:   10,
		},
		{
			name:   "descending order",
			values: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			want:   10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var tree *Tree
			for _, val := range tt.values {
				tree = tree.Insert(val)
			}

			got := tree.Count()
			if got != tt.want {
				t.Errorf("Count() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestTree_Count_EmptyTree(t *testing.T) {
	var tree *Tree
	if count := tree.Count(); count != 0 {
		t.Errorf("Count() on nil tree = %d, want 0", count)
	}
}

func TestTree_Count_AfterDelete(t *testing.T) {
	var tree *Tree
	values := []int{5, 3, 7, 1, 4, 6, 9}
	for _, val := range values {
		tree = tree.Insert(val)
	}

	if count := tree.Count(); count != 7 {
		t.Errorf("Initial Count() = %d, want 7", count)
	}

	// Delete one element
	tree = tree.Delete(3)
	if count := tree.Count(); count != 6 {
		t.Errorf("After Delete(3), Count() = %d, want 6", count)
	}

	// Delete another element
	tree = tree.Delete(7)
	if count := tree.Count(); count != 5 {
		t.Errorf("After Delete(7), Count() = %d, want 5", count)
	}

	// Delete root
	tree = tree.Delete(5)
	if count := tree.Count(); count != 4 {
		t.Errorf("After Delete(5), Count() = %d, want 4", count)
	}
}

func TestTree_Count_AfterDeleteNonExistent(t *testing.T) {
	var tree *Tree
	values := []int{5, 3, 7}
	for _, val := range values {
		tree = tree.Insert(val)
	}

	initialCount := tree.Count()

	// Delete non-existent element
	tree = tree.Delete(10)

	if count := tree.Count(); count != initialCount {
		t.Errorf("After deleting non-existent element, Count() = %d, want %d", count, initialCount)
	}
}

func TestTree_Count_Progressive(t *testing.T) {
	var tree *Tree

	// Start with empty tree
	if count := tree.Count(); count != 0 {
		t.Errorf("Empty tree Count() = %d, want 0", count)
	}

	// Add elements one by one
	for i := 1; i <= 5; i++ {
		tree = tree.Insert(i * 10)
		if count := tree.Count(); count != i {
			t.Errorf("After inserting %d elements, Count() = %d, want %d", i, count, i)
		}
	}

	// Remove elements one by one
	deleteVals := []int{10, 30, 50}
	expectedCount := 5
	for _, val := range deleteVals {
		tree = tree.Delete(val)
		expectedCount--
		if count := tree.Count(); count != expectedCount {
			t.Errorf("After deleting %d, Count() = %d, want %d", val, count, expectedCount)
		}
	}
}

func TestTree_Count_VsInorder(t *testing.T) {
	// Count should match the length of inorder traversal
	var tree *Tree
	values := []int{50, 30, 70, 20, 40, 60, 80, 10, 25, 35, 45}
	for _, val := range values {
		tree = tree.Insert(val)
	}

	count := tree.Count()
	inorder := collectInorder(tree)

	if count != len(inorder) {
		t.Errorf("Count() = %d, but inorder traversal length = %d", count, len(inorder))
	}
}
