package lesson_4

import (
	"testing"
)

func TestNewOrderedMap(t *testing.T) {
	om := NewOrderedMap()

	if om == nil {
		t.Fatal("NewOrderedMap() returned nil")
	}

	if om.Size() != 0 {
		t.Errorf("NewOrderedMap().Size() = %d, want 0", om.Size())
	}

	if om.Contains(1) {
		t.Error("NewOrderedMap() should not contain any keys")
	}
}

func TestOrderedMap_Insert(t *testing.T) {
	om := NewOrderedMap()

	om.Insert(5, 50)
	if !om.Contains(5) {
		t.Error("After Insert(5, 50), Contains(5) should be true")
	}

	if om.Size() != 1 {
		t.Errorf("After one Insert, Size() = %d, want 1", om.Size())
	}
}

func TestOrderedMap_Insert_Multiple(t *testing.T) {
	om := NewOrderedMap()

	pairs := []struct {
		key   int
		value int
	}{
		{5, 50},
		{3, 30},
		{7, 70},
		{1, 10},
		{9, 90},
	}

	for _, p := range pairs {
		om.Insert(p.key, p.value)
	}

	if om.Size() != len(pairs) {
		t.Errorf("After %d inserts, Size() = %d, want %d", len(pairs), om.Size(), len(pairs))
	}

	for _, p := range pairs {
		if !om.Contains(p.key) {
			t.Errorf("Contains(%d) = false, want true", p.key)
		}
	}
}

func TestOrderedMap_Insert_Duplicates(t *testing.T) {
	om := NewOrderedMap()

	om.Insert(5, 50)
	om.Insert(5, 500) // Duplicate key

	// Size should still be 1 (not 2) when inserting duplicate key
	if om.Size() != 1 {
		t.Errorf("After inserting duplicate key, Size() = %d, want 1", om.Size())
	}
}

func TestOrderedMap_Erase(t *testing.T) {
	om := NewOrderedMap()

	om.Insert(5, 50)
	om.Insert(3, 30)
	om.Insert(7, 70)

	om.Erase(5)

	if om.Contains(5) {
		t.Error("After Erase(5), Contains(5) should be false")
	}

	if om.Size() != 2 {
		t.Errorf("After erasing 1 of 3 elements, Size() = %d, want 2", om.Size())
	}

	// Other elements should still exist
	if !om.Contains(3) || !om.Contains(7) {
		t.Error("Erase(5) affected other elements")
	}
}

func TestOrderedMap_Erase_NonExistent(t *testing.T) {
	om := NewOrderedMap()

	om.Insert(5, 50)
	originalSize := om.Size()

	om.Erase(10) // Erase key that doesn't exist

	// Size should not change when erasing non-existent key
	if om.Size() != originalSize {
		t.Errorf("After erasing non-existent key, Size() = %d, want %d", om.Size(), originalSize)
	}
}

func TestOrderedMap_Erase_All(t *testing.T) {
	om := NewOrderedMap()

	keys := []int{5, 3, 7, 1, 9}
	for _, k := range keys {
		om.Insert(k, k*10)
	}

	// Erase all elements
	for _, k := range keys {
		om.Erase(k)
	}

	if om.Size() != 0 {
		t.Errorf("After erasing all elements, Size() = %d, want 0", om.Size())
	}

	for _, k := range keys {
		if om.Contains(k) {
			t.Errorf("After erasing all, Contains(%d) should be false", k)
		}
	}
}

func TestOrderedMap_Contains(t *testing.T) {
	om := NewOrderedMap()

	om.Insert(5, 50)
	om.Insert(3, 30)
	om.Insert(7, 70)

	tests := []struct {
		key  int
		want bool
	}{
		{5, true},
		{3, true},
		{7, true},
		{1, false},
		{10, false},
		{4, false},
	}

	for _, tt := range tests {
		if got := om.Contains(tt.key); got != tt.want {
			t.Errorf("Contains(%d) = %v, want %v", tt.key, got, tt.want)
		}
	}
}

func TestOrderedMap_Contains_Empty(t *testing.T) {
	om := NewOrderedMap()

	if om.Contains(1) {
		t.Error("Empty map should not contain any keys")
	}
}

func TestOrderedMap_Size(t *testing.T) {
	om := NewOrderedMap()

	if om.Size() != 0 {
		t.Errorf("Empty map Size() = %d, want 0", om.Size())
	}

	for i := 1; i <= 5; i++ {
		om.Insert(i, i*10)
		if om.Size() != i {
			t.Errorf("After %d inserts, Size() = %d, want %d", i, om.Size(), i)
		}
	}

	for i := 1; i <= 5; i++ {
		om.Erase(i)
		expectedSize := 5 - i
		if om.Size() != expectedSize {
			t.Errorf("After erasing %d elements, Size() = %d, want %d", i, om.Size(), expectedSize)
		}
	}
}

func TestOrderedMap_ForEach(t *testing.T) {
	om := NewOrderedMap()

	// Insert in random order
	pairs := map[int]int{
		5: 50,
		3: 30,
		7: 70,
		1: 10,
		9: 90,
		4: 40,
		6: 60,
	}

	for k, v := range pairs {
		om.Insert(k, v)
	}

	// Collect keys via ForEach
	var keys []int

	om.ForEach(func(key, value int) {
		keys = append(keys, key)
	})

	// Keys should be in sorted order (inorder traversal of BST)
	expectedKeys := []int{1, 3, 4, 5, 6, 7, 9}
	if !slicesEqual(keys, expectedKeys) {
		t.Errorf("ForEach keys = %v, want %v (sorted order)", keys, expectedKeys)
	}

	// Verify all keys were visited
	if len(keys) != om.Size() {
		t.Errorf("ForEach visited %d keys, but Size() = %d", len(keys), om.Size())
	}
}

func TestOrderedMap_ForEach_Empty(t *testing.T) {
	om := NewOrderedMap()

	callCount := 0
	om.ForEach(func(key, value int) {
		callCount++
	})

	if callCount != 0 {
		t.Errorf("ForEach on empty map called action %d times, want 0", callCount)
	}
}

func TestOrderedMap_ForEach_Order(t *testing.T) {
	om := NewOrderedMap()

	// Insert in descending order
	for i := 10; i >= 1; i-- {
		om.Insert(i, i*100)
	}

	// ForEach should visit in ascending order
	previousKey := -1
	om.ForEach(func(key, value int) {
		if key <= previousKey {
			t.Errorf("ForEach not in ascending order: got key %d after %d", key, previousKey)
		}
		previousKey = key
	})
}

func TestOrderedMap_ForEach_AfterErase(t *testing.T) {
	om := NewOrderedMap()

	// Insert elements
	for i := 1; i <= 10; i++ {
		om.Insert(i, i*10)
	}

	// Erase some elements
	om.Erase(3)
	om.Erase(7)
	om.Erase(9)

	// Collect keys
	var keys []int
	om.ForEach(func(key, value int) {
		keys = append(keys, key)
	})

	expectedKeys := []int{1, 2, 4, 5, 6, 8, 10}
	if !slicesEqual(keys, expectedKeys) {
		t.Errorf("ForEach after erases, keys = %v, want %v", keys, expectedKeys)
	}

	// Erased keys should not appear
	for _, key := range keys {
		if key == 3 || key == 7 || key == 9 {
			t.Errorf("ForEach visited erased key %d", key)
		}
	}
}

func TestOrderedMap_ComplexWorkflow(t *testing.T) {
	om := NewOrderedMap()

	// Insert several elements
	om.Insert(5, 50)
	om.Insert(3, 30)
	om.Insert(7, 70)
	om.Insert(1, 10)
	om.Insert(9, 90)

	if om.Size() != 5 {
		t.Errorf("After 5 inserts, Size() = %d, want 5", om.Size())
	}

	// Delete some elements
	om.Erase(3)
	om.Erase(7)

	if om.Size() != 3 {
		t.Errorf("After 2 erases, Size() = %d, want 3", om.Size())
	}

	// Verify remaining elements
	if !om.Contains(5) || !om.Contains(1) || !om.Contains(9) {
		t.Error("Remaining elements not found")
	}

	if om.Contains(3) || om.Contains(7) {
		t.Error("Erased elements still found")
	}

	// ForEach should only visit remaining elements
	var keys []int
	om.ForEach(func(key, value int) {
		keys = append(keys, key)
	})

	expectedKeys := []int{1, 5, 9}
	if !slicesEqual(keys, expectedKeys) {
		t.Errorf("ForEach after erases, keys = %v, want %v", keys, expectedKeys)
	}
}

func TestOrderedMap_StressTest(t *testing.T) {
	om := NewOrderedMap()

	// Insert many elements
	n := 100
	for i := 0; i < n; i++ {
		om.Insert(i, i*10)
	}

	if om.Size() != n {
		t.Errorf("After inserting %d elements, Size() = %d, want %d", n, om.Size(), n)
	}

	// Verify all present
	for i := 0; i < n; i++ {
		if !om.Contains(i) {
			t.Errorf("Contains(%d) = false, want true", i)
		}
	}

	// Verify order via ForEach
	expectedKey := 0
	om.ForEach(func(key, value int) {
		if key != expectedKey {
			t.Errorf("ForEach: expected key %d, got %d", expectedKey, key)
		}
		expectedKey++
	})

	// Erase half the elements
	for i := 0; i < n; i += 2 {
		om.Erase(i)
	}

	if om.Size() != n/2 {
		t.Errorf("After erasing half, Size() = %d, want %d", om.Size(), n/2)
	}

	// Verify erased keys are gone
	for i := 0; i < n; i += 2 {
		if om.Contains(i) {
			t.Errorf("Contains(%d) = true after erase, want false", i)
		}
	}

	// Verify remaining keys still present
	for i := 1; i < n; i += 2 {
		if !om.Contains(i) {
			t.Errorf("Contains(%d) = false, want true", i)
		}
	}
}
