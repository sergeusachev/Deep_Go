/*

1. func Map(data []int, action func(int) int) []int
2. func Filter(data []int, action func(int) bool) []int
3. func Reduce(data []int, initial int, action func(int, int) int) int

*/

package lesson_5

func Map(data []int, action func(int) int) []int {
	mapped := make([]int, len(data))
	for i, v := range data {
		mapped[i] = action(v)
	}
	return mapped
}

func Filter(data []int, action func(int) bool) []int {
	var filtered []int
	for _, v := range data {
		if action(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func Reduce(data []int, initial int, action func(int, int) int) int {
	accumulator := initial
	for _, v := range data {
		accumulator = action(accumulator, v)
	}
	return accumulator
}





