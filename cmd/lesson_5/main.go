package main

import(
	"fmt"
	fp "deepgo/pkg/lesson_5"
)

func main() {
	
}

func filterCheck() {
	s := []int{1, 2, 3, 4, 5}
	filteredS := fp.Filter(s, func(el int) bool { 
		return el % 2 == 0
	})
	fmt.Println("s:", s)
	fmt.Println("filteredS:", filteredS)
}

func mapCheck() {
	s := []int{1, 2, 3, 4, 5}
	mappedS := fp.Map(s, func(el int) int { 
		return el * 2 
	})
	fmt.Println("s:", s)
	fmt.Println("mappedS:", mappedS)
}

func opt1() {
	defer fmt.Println("1")
	var f func()
	defer f()

	fmt.Println("2")
}

func opt2() {
	var f func()
	defer f()

	f = func ()  {
		fmt.Println("1")
	}
	fmt.Println("2")
}

func opt3() {
	fn := compose(wrapPrint(sqr), wrapPrint(neg), wrapPrint(sqr))
	fn(2)
}

func sqr(number int) int {
	return number * number
}

func neg(number int) int {
	return -number
}

func wrapPrint(fn func(int) int) func (int) int  {
	return func(value int) int {
		res := fn(value)
		fmt.Println(res)
		return res
	}
}

func compose(fns...func(int) int) func(int) int {
	return func(value int) int {
		for _, v := range fns {
			value = v(value)
		}

		return value
	}
}