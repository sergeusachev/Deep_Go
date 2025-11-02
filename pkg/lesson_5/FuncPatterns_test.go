package lesson_5

import (
	"testing"
)

func TestMap(t *testing.T) {
	tests := []struct {
		name   string
		data   []int
		action func(int) int
		want   []int
	}{
		{
			name:   "double each element",
			data:   []int{1, 2, 3, 4, 5},
			action: func(x int) int { return x * 2 },
			want:   []int{2, 4, 6, 8, 10},
		},
		{
			name:   "square each element",
			data:   []int{1, 2, 3, 4},
			action: func(x int) int { return x * x },
			want:   []int{1, 4, 9, 16},
		},
		{
			name:   "add constant",
			data:   []int{1, 2, 3},
			action: func(x int) int { return x + 10 },
			want:   []int{11, 12, 13},
		},
		{
			name:   "negate each element",
			data:   []int{1, -2, 3, -4},
			action: func(x int) int { return -x },
			want:   []int{-1, 2, -3, 4},
		},
		{
			name:   "identity function",
			data:   []int{1, 2, 3},
			action: func(x int) int { return x },
			want:   []int{1, 2, 3},
		},
		{
			name:   "constant function",
			data:   []int{1, 2, 3, 4},
			action: func(x int) int { return 42 },
			want:   []int{42, 42, 42, 42},
		},
		{
			name:   "empty slice",
			data:   []int{},
			action: func(x int) int { return x * 2 },
			want:   []int{},
		},
		{
			name:   "single element",
			data:   []int{5},
			action: func(x int) int { return x * 10 },
			want:   []int{50},
		},
		{
			name:   "zero values",
			data:   []int{0, 0, 0},
			action: func(x int) int { return x + 1 },
			want:   []int{1, 1, 1},
		},
		{
			name:   "negative numbers",
			data:   []int{-5, -10, -15},
			action: func(x int) int { return x * -1 },
			want:   []int{5, 10, 15},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Map(tt.data, tt.action)
			if !slicesEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMap_OriginalUnchanged(t *testing.T) {
	original := []int{1, 2, 3, 4, 5}
	originalCopy := make([]int, len(original))
	copy(originalCopy, original)

	_ = Map(original, func(x int) int { return x * 2 })

	// Original slice should not be modified
	if !slicesEqual(original, originalCopy) {
		t.Errorf("Map() modified original slice: got %v, want %v", original, originalCopy)
	}
}

func TestMap_DifferentTransformations(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}

	// Apply multiple transformations
	doubled := Map(data, func(x int) int { return x * 2 })
	squared := Map(data, func(x int) int { return x * x })
	incremented := Map(data, func(x int) int { return x + 1 })

	expectedDoubled := []int{2, 4, 6, 8, 10}
	expectedSquared := []int{1, 4, 9, 16, 25}
	expectedIncremented := []int{2, 3, 4, 5, 6}

	if !slicesEqual(doubled, expectedDoubled) {
		t.Errorf("doubled = %v, want %v", doubled, expectedDoubled)
	}
	if !slicesEqual(squared, expectedSquared) {
		t.Errorf("squared = %v, want %v", squared, expectedSquared)
	}
	if !slicesEqual(incremented, expectedIncremented) {
		t.Errorf("incremented = %v, want %v", incremented, expectedIncremented)
	}
}

func TestMap_ChainedOperations(t *testing.T) {
	data := []int{1, 2, 3}

	// Chain multiple Map operations
	result := Map(Map(data, func(x int) int { return x * 2 }), func(x int) int { return x + 1 })

	expected := []int{3, 5, 7} // (1*2)+1=3, (2*2)+1=5, (3*2)+1=7
	if !slicesEqual(result, expected) {
		t.Errorf("chained Map() = %v, want %v", result, expected)
	}
}

func TestMap_LargeSlice(t *testing.T) {
	// Test with larger slice
	n := 1000
	data := make([]int, n)
	for i := 0; i < n; i++ {
		data[i] = i
	}

	result := Map(data, func(x int) int { return x * 2 })

	if len(result) != n {
		t.Errorf("Map() returned slice of length %d, want %d", len(result), n)
	}

	// Spot check some values
	if result[0] != 0 {
		t.Errorf("result[0] = %d, want 0", result[0])
	}
	if result[100] != 200 {
		t.Errorf("result[100] = %d, want 200", result[100])
	}
	if result[999] != 1998 {
		t.Errorf("result[999] = %d, want 1998", result[999])
	}
}

func TestMap_PreservesOrder(t *testing.T) {
	data := []int{5, 3, 8, 1, 9, 2}
	result := Map(data, func(x int) int { return x * 10 })

	expected := []int{50, 30, 80, 10, 90, 20}
	if !slicesEqual(result, expected) {
		t.Errorf("Map() = %v, want %v (order not preserved)", result, expected)
	}
}

func TestMap_ComplexFunction(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Complex transformation: if even, double it; if odd, square it
	result := Map(data, func(x int) int {
		if x%2 == 0 {
			return x * 2
		}
		return x * x
	})

	expected := []int{1, 4, 9, 8, 25, 12, 49, 16, 81, 20}
	if !slicesEqual(result, expected) {
		t.Errorf("Map() = %v, want %v", result, expected)
	}
}

func TestFilter(t *testing.T) {
	tests := []struct {
		name   string
		data   []int
		action func(int) bool
		want   []int
	}{
		{
			name:   "filter even numbers",
			data:   []int{1, 2, 3, 4, 5, 6},
			action: func(x int) bool { return x%2 == 0 },
			want:   []int{2, 4, 6},
		},
		{
			name:   "filter odd numbers",
			data:   []int{1, 2, 3, 4, 5, 6},
			action: func(x int) bool { return x%2 != 0 },
			want:   []int{1, 3, 5},
		},
		{
			name:   "filter positive numbers",
			data:   []int{-3, -1, 0, 1, 2, 3},
			action: func(x int) bool { return x > 0 },
			want:   []int{1, 2, 3},
		},
		{
			name:   "filter negative numbers",
			data:   []int{-3, -1, 0, 1, 2, 3},
			action: func(x int) bool { return x < 0 },
			want:   []int{-3, -1},
		},
		{
			name:   "filter greater than threshold",
			data:   []int{1, 5, 3, 8, 2, 10},
			action: func(x int) bool { return x > 5 },
			want:   []int{8, 10},
		},
		{
			name:   "filter all pass",
			data:   []int{2, 4, 6, 8},
			action: func(x int) bool { return true },
			want:   []int{2, 4, 6, 8},
		},
		{
			name:   "filter none pass",
			data:   []int{1, 2, 3, 4},
			action: func(x int) bool { return false },
			want:   []int{},
		},
		{
			name:   "empty slice",
			data:   []int{},
			action: func(x int) bool { return x > 0 },
			want:   []int{},
		},
		{
			name:   "single element matches",
			data:   []int{5},
			action: func(x int) bool { return x == 5 },
			want:   []int{5},
		},
		{
			name:   "single element doesn't match",
			data:   []int{5},
			action: func(x int) bool { return x == 3 },
			want:   []int{},
		},
		{
			name:   "filter by divisibility",
			data:   []int{1, 2, 3, 4, 5, 6, 9, 12, 15},
			action: func(x int) bool { return x%3 == 0 },
			want:   []int{3, 6, 9, 12, 15},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Filter(tt.data, tt.action)
			if !slicesEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilter_OriginalUnchanged(t *testing.T) {
	original := []int{1, 2, 3, 4, 5, 6}
	originalCopy := make([]int, len(original))
	copy(originalCopy, original)

	_ = Filter(original, func(x int) bool { return x%2 == 0 })

	// Original slice should not be modified
	if !slicesEqual(original, originalCopy) {
		t.Errorf("Filter() modified original slice: got %v, want %v", original, originalCopy)
	}
}

func TestFilter_PreservesOrder(t *testing.T) {
	data := []int{5, 3, 8, 1, 9, 2, 4, 7, 6}
	result := Filter(data, func(x int) bool { return x%2 == 0 })

	// Even numbers in original order
	expected := []int{8, 2, 4, 6}
	if !slicesEqual(result, expected) {
		t.Errorf("Filter() = %v, want %v (order not preserved)", result, expected)
	}
}

func TestFilter_ComplexPredicate(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	// Filter numbers that are even AND greater than 5
	result := Filter(data, func(x int) bool {
		return x%2 == 0 && x > 5
	})

	expected := []int{6, 8, 10, 12}
	if !slicesEqual(result, expected) {
		t.Errorf("Filter() = %v, want %v", result, expected)
	}
}

func TestFilter_ChainedOperations(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// First filter even numbers, then filter those > 5
	result := Filter(Filter(data, func(x int) bool { return x%2 == 0 }), func(x int) bool { return x > 5 })

	expected := []int{6, 8, 10}
	if !slicesEqual(result, expected) {
		t.Errorf("chained Filter() = %v, want %v", result, expected)
	}
}

func TestFilter_WithMap(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6}

	// Filter even numbers, then double them
	filtered := Filter(data, func(x int) bool { return x%2 == 0 })
	result := Map(filtered, func(x int) int { return x * 2 })

	expected := []int{4, 8, 12}
	if !slicesEqual(result, expected) {
		t.Errorf("Filter+Map = %v, want %v", result, expected)
	}
}

func TestFilter_MapThenFilter(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}

	// Double all numbers, then filter those > 5
	doubled := Map(data, func(x int) int { return x * 2 })
	result := Filter(doubled, func(x int) bool { return x > 5 })

	expected := []int{6, 8, 10}
	if !slicesEqual(result, expected) {
		t.Errorf("Map+Filter = %v, want %v", result, expected)
	}
}

func TestFilter_LargeSlice(t *testing.T) {
	// Test with larger slice
	n := 1000
	data := make([]int, n)
	for i := 0; i < n; i++ {
		data[i] = i
	}

	// Filter even numbers
	result := Filter(data, func(x int) bool { return x%2 == 0 })

	if len(result) != 500 {
		t.Errorf("Filter() returned slice of length %d, want 500", len(result))
	}

	// Spot check some values
	if result[0] != 0 {
		t.Errorf("result[0] = %d, want 0", result[0])
	}
	if result[1] != 2 {
		t.Errorf("result[1] = %d, want 2", result[1])
	}
	if result[499] != 998 {
		t.Errorf("result[499] = %d, want 998", result[499])
	}
}

func TestFilter_ZeroValues(t *testing.T) {
	data := []int{0, 1, 0, 2, 0, 3}

	// Filter out zeros
	result := Filter(data, func(x int) bool { return x != 0 })

	expected := []int{1, 2, 3}
	if !slicesEqual(result, expected) {
		t.Errorf("Filter() = %v, want %v", result, expected)
	}
}

func TestFilter_OnlyZeros(t *testing.T) {
	data := []int{0, 1, 0, 2, 0, 3}

	// Filter only zeros
	result := Filter(data, func(x int) bool { return x == 0 })

	expected := []int{0, 0, 0}
	if !slicesEqual(result, expected) {
		t.Errorf("Filter() = %v, want %v", result, expected)
	}
}

func TestFilter_NegativePositiveSplit(t *testing.T) {
	data := []int{-5, -2, -1, 0, 1, 2, 5, 10}

	negatives := Filter(data, func(x int) bool { return x < 0 })
	positives := Filter(data, func(x int) bool { return x > 0 })

	expectedNeg := []int{-5, -2, -1}
	expectedPos := []int{1, 2, 5, 10}

	if !slicesEqual(negatives, expectedNeg) {
		t.Errorf("negatives = %v, want %v", negatives, expectedNeg)
	}
	if !slicesEqual(positives, expectedPos) {
		t.Errorf("positives = %v, want %v", positives, expectedPos)
	}
}

func TestReduce(t *testing.T) {
	tests := []struct {
		name    string
		data    []int
		initial int
		action  func(int, int) int
		want    int
	}{
		{
			name:    "sum all elements",
			data:    []int{1, 2, 3, 4, 5},
			initial: 0,
			action:  func(acc, x int) int { return acc + x },
			want:    15,
		},
		{
			name:    "product of all elements",
			data:    []int{1, 2, 3, 4, 5},
			initial: 1,
			action:  func(acc, x int) int { return acc * x },
			want:    120,
		},
		{
			name:    "find maximum",
			data:    []int{3, 7, 2, 9, 1, 5},
			initial: 0,
			action:  func(acc, x int) int {
				if x > acc {
					return x
				}
				return acc
			},
			want:    9,
		},
		{
			name:    "find minimum",
			data:    []int{3, 7, 2, 9, 1, 5},
			initial: 999,
			action:  func(acc, x int) int {
				if x < acc {
					return x
				}
				return acc
			},
			want:    1,
		},
		{
			name:    "count elements",
			data:    []int{1, 2, 3, 4, 5},
			initial: 0,
			action:  func(acc, x int) int { return acc + 1 },
			want:    5,
		},
		{
			name:    "sum with non-zero initial",
			data:    []int{1, 2, 3},
			initial: 10,
			action:  func(acc, x int) int { return acc + x },
			want:    16,
		},
		{
			name:    "empty slice returns initial",
			data:    []int{},
			initial: 42,
			action:  func(acc, x int) int { return acc + x },
			want:    42,
		},
		{
			name:    "single element",
			data:    []int{5},
			initial: 0,
			action:  func(acc, x int) int { return acc + x },
			want:    5,
		},
		{
			name:    "subtraction (order matters)",
			data:    []int{10, 3, 2},
			initial: 0,
			action:  func(acc, x int) int { return acc - x },
			want:    -15, // 0 - 10 - 3 - 2 = -15
		},
		{
			name:    "build number from accumulation",
			data:    []int{1, 10, 100},
			initial: 0,
			action:  func(acc, x int) int { return acc*10 + x },
			want:    300, // 0*10+1=1, 1*10+10=20, 20*10+100=300
		},
		{
			name:    "all zeros",
			data:    []int{0, 0, 0, 0},
			initial: 0,
			action:  func(acc, x int) int { return acc + x },
			want:    0,
		},
		{
			name:    "negative numbers sum",
			data:    []int{-1, -2, -3, -4},
			initial: 0,
			action:  func(acc, x int) int { return acc + x },
			want:    -10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Reduce(tt.data, tt.initial, tt.action)
			if got != tt.want {
				t.Errorf("Reduce() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestReduce_Sum(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := Reduce(data, 0, func(acc, x int) int { return acc + x })

	if sum != 55 {
		t.Errorf("sum = %d, want 55", sum)
	}
}

func TestReduce_Product(t *testing.T) {
	data := []int{2, 3, 4}
	product := Reduce(data, 1, func(acc, x int) int { return acc * x })

	if product != 24 {
		t.Errorf("product = %d, want 24", product)
	}
}

func TestReduce_Max(t *testing.T) {
	data := []int{3, 7, 2, 9, 1, 5, 8}
	max := Reduce(data, data[0], func(acc, x int) int {
		if x > acc {
			return x
		}
		return acc
	})

	if max != 9 {
		t.Errorf("max = %d, want 9", max)
	}
}

func TestReduce_Min(t *testing.T) {
	data := []int{3, 7, 2, 9, 1, 5, 8}
	min := Reduce(data, data[0], func(acc, x int) int {
		if x < acc {
			return x
		}
		return acc
	})

	if min != 1 {
		t.Errorf("min = %d, want 1", min)
	}
}

func TestReduce_EmptySlice(t *testing.T) {
	data := []int{}
	result := Reduce(data, 100, func(acc, x int) int { return acc + x })

	if result != 100 {
		t.Errorf("Reduce on empty slice = %d, want 100 (initial value)", result)
	}
}

func TestReduce_OrderMatters(t *testing.T) {
	data := []int{1, 2, 3}

	// Subtraction is not commutative - order matters
	// 0 - 1 - 2 - 3 = -6
	result := Reduce(data, 0, func(acc, x int) int { return acc - x })

	if result != -6 {
		t.Errorf("Reduce with subtraction = %d, want -6", result)
	}
}

func TestReduce_WithMapFilter(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Filter even numbers, map to squares, then sum
	evens := Filter(data, func(x int) bool { return x%2 == 0 })
	squares := Map(evens, func(x int) int { return x * x })
	sum := Reduce(squares, 0, func(acc, x int) int { return acc + x })

	// Evens: 2, 4, 6, 8, 10
	// Squares: 4, 16, 36, 64, 100
	// Sum: 220
	if sum != 220 {
		t.Errorf("Map+Filter+Reduce pipeline = %d, want 220", sum)
	}
}

func TestReduce_LargeSlice(t *testing.T) {
	// Sum of numbers from 0 to 999 = 999 * 1000 / 2 = 499500
	n := 1000
	data := make([]int, n)
	for i := 0; i < n; i++ {
		data[i] = i
	}

	sum := Reduce(data, 0, func(acc, x int) int { return acc + x })
	expected := 499500

	if sum != expected {
		t.Errorf("Reduce large slice = %d, want %d", sum, expected)
	}
}

func TestReduce_BuildNumber(t *testing.T) {
	// Build a number from digits: [1, 2, 3] -> 123
	data := []int{1, 2, 3}
	number := Reduce(data, 0, func(acc, x int) int { return acc*10 + x })

	if number != 123 {
		t.Errorf("Build number from digits = %d, want 123", number)
	}
}

func TestReduce_CountEvens(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Count even numbers using Reduce
	count := Reduce(data, 0, func(acc, x int) int {
		if x%2 == 0 {
			return acc + 1
		}
		return acc
	})

	if count != 5 {
		t.Errorf("Count evens = %d, want 5", count)
	}
}

func TestReduce_AbsoluteSum(t *testing.T) {
	data := []int{-3, 5, -2, 8, -1}

	// Sum of absolute values
	absSum := Reduce(data, 0, func(acc, x int) int {
		if x < 0 {
			return acc - x
		}
		return acc + x
	})

	// |-3| + |5| + |-2| + |8| + |-1| = 3 + 5 + 2 + 8 + 1 = 19
	if absSum != 19 {
		t.Errorf("Absolute sum = %d, want 19", absSum)
	}
}

func TestReduce_Factorial(t *testing.T) {
	// Calculate 5! = 1*2*3*4*5 = 120
	data := []int{1, 2, 3, 4, 5}
	factorial := Reduce(data, 1, func(acc, x int) int { return acc * x })

	if factorial != 120 {
		t.Errorf("5! = %d, want 120", factorial)
	}
}

func TestReduce_MapReducePattern(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}

	// Square each number, then sum (classic MapReduce)
	sumOfSquares := Reduce(
		Map(data, func(x int) int { return x * x }),
		0,
		func(acc, x int) int { return acc + x },
	)

	// 1 + 4 + 9 + 16 + 25 = 55
	if sumOfSquares != 55 {
		t.Errorf("Sum of squares = %d, want 55", sumOfSquares)
	}
}

// Helper function
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
