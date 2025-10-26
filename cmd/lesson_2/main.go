package main

import(
	"fmt"
)

func main() {
	s := []int{1, 2, 3, 4}
	fmt.Printf("len(s)=%d, cap(s)=%d, s=%v\n\n", len(s), cap(s), s)
	s = append(s, 5)
	s = append(s, 7)
	s = append(s, 8)
	s = append(s, 8)
	s = append(s, 8)
	fmt.Printf("len(s)=%d, cap(s)=%d, s=%v\n", len(s), cap(s), s)
}
