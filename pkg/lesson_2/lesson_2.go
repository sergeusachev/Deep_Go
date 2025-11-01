package lesson_2

import(
	"fmt"
	"strings"
	"unsafe"
	"runtime"
)

func main() {
	gcSliceTest()
}

func clearTest() {
	
}

var Fake []byte //nil

func gcSliceTest() {
	data := make([]byte, 1<<30)
	// let's imagine that data was read from a file

	sequence := findSequenceCopy(data)
	_ = sequence // using of sequence later

	printAllocs()
	runtime.GC()
	printAllocs()

	runtime.KeepAlive(sequence)
}

func findSequenceCopy(data []byte) []byte {
	for i := 0; i < len(data)-1; i++ {
		if data[i] == 0x00 && data[i+1] == 0x00 {
			partData := make([]byte, 20)
			copy(partData, data[i:])
			return partData
		}
	}
	return nil
}

func findSequence(data []byte) []byte {
	for i := 0; i < len(data)-1; i++ {
		if data[i] == 0x00 && data[i+1] == 0x00 {
			return data[i:i+20]
		}
	}
	return nil
}

func printAllocs() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d MB\n", m.Alloc/1024/1024)
}

func sliceReducing() {
//                                           9 | 10
	s := []int{1, 2, 3, 4, 5, 6, 7 , 8 , 9, 10}
	//s = s[:len(s) - 1]
	s = append(s[:0], s[1:]...)
	fmt.Println(s)
}

func foo11() {
	var arr *[4]int // = nil
	
	fmt.Println("arr == nil:", arr == nil)
	
}

func foo10() {
	a := [3]int{5, 7, 0}
	b := a
	a[2] = 13
	
	fmt.Println(a)
	fmt.Println(b)
}

func foo9() {
	a := [3]int{1, 2, 3}
	b := [3]int{4, 5, 6}
	
	fmt.Println("BEFORE a:", a)
	fmt.Printf("BEFORE &a: %p\n", &a)
	fmt.Println("BEFORE &a[0]:", &a[0])
	for i, v := range a {
		if i == 0 {
			fmt.Println("LOOP_1 a:", a)
			fmt.Println("LOOP_1 &a[0]:", &a[0])
			fmt.Println("LOOP_1 v:", v)
			fmt.Printf("LOOP_1 &v: %p\n", &v)
		}
		
		if i == 1 {
			fmt.Println("LOOP_2 a:", a)
			fmt.Println("LOOP_2 &a[1]:", &a[1])
			fmt.Println("LOOP_2 v:", v)
			fmt.Printf("LOOP_2 &v: %p\n", &v)
			fmt.Println("a = b")
			a = b
			fmt.Println("LOOP_2 a:", a)
			fmt.Println("LOOP_2 &a[1]:", &a[1])
			fmt.Println("LOOP_2 v:", v)
			fmt.Printf("LOOP_2 &v: %p\n", &v)
		}
		
		if i == 2 {
			fmt.Println("LOOP_3 a:", a)
			fmt.Println("LOOP_3 &a[2]:", &a[2])
			fmt.Println("LOOP_3 v:", v)
			fmt.Printf("LOOP_3 &v: %p\n", &v)
		}
		fmt.Println(v)
		fmt.Println(a[i])
		fmt.Println()
	}
}

func foo8() {
	//indices:	0  1  2  3  4
	a := [5]int{1, 2, 3, 4, 5}
    b := a[5:] // It's OK
    fmt.Println("b:", b)
    fmt.Println("len(b)", len(b))
    fmt.Println("cap(b)", cap(b))
    b = append(b, 13)
    fmt.Println("b[0]:", b[0])	
   
    //fmt.Println("a[5]:", a[5]) // compiler error
}

func foo7() {
	s := []int{0, 1, 2, 3, 4, 5}
	s1 := s[:4] 	// {0, 1, 2, 3}
	s2 := s[2:5]	// {      2, 3, 4}
	copy(s1, s2)
	
	fmt.Println(s1)
}

func foo6() {
	s := make([]byte, 5)
	s = s[2:4]
	fmt.Println("s:", s)
	fmt.Println("len(s):", len(s))
	fmt.Println("cap(s):", cap(s))
	
	fmt.Println()	
	
	s = s[:cap(s)]
	fmt.Println("s:", s)
	fmt.Println("len(s):", len(s))
	fmt.Println("cap(s):", cap(s))
}

func foo5() {
	//b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	//fmt.Println(b[1:4] == []byte{'o', 'l', 'a'})
}

func foo4() {
	var x = []string{"A", "M", "C"}
	//fmt.Println("(BEFORE range) x:", x)
	//fmt.Printf("(BEFORE range) &x: %p\n", &x)
	fmt.Println("(BEFORE) len(x):", len(x))
	fmt.Println("(BEFORE) cap(x):", cap(x))
	
	for i, s := range x {
		fmt.Println(i, s)
		
		fmt.Println("(IN range) x:", x)
		fmt.Printf("(IN range) &x: %p\n", &x)
		
		x[i+1] = "M"
		//fmt.Println("(IN tange BEFORE append) len(x):", len(x))
		//fmt.Println("(IN tange BEFORE append) cap(x):", cap(x))
		x = append(x, "Z")
		//fmt.Println("(IN tange AFTER append) len(x):", len(x))
		//fmt.Println("(IN tange AFTER append) cap(x):", cap(x))
		
		fmt.Println("(IN range AFTER append) x:", x)
		fmt.Printf("(IN range AFTER append) &x: %p\n", &x)

		x[i+1] = "Z"
	}
}

func foo3() {
	sliceA := make([]byte, 10 << 20)
	fmt.Println("len(sliceA):", len(sliceA))
	fmt.Println("cap(sliceA):", cap(sliceA))
	
	fmt.Println()
	
	sliceB := makeDirty(10<<20)
	fmt.Println("sliceB:", sliceB)
	fmt.Println("len(sliceB):", len(sliceB))
	fmt.Println("cap(sliceB):", cap(sliceB))
}

func makeDirty(size int) []byte {
	var sb strings.Builder
	sb.Grow(size)
	
	pointer := unsafe.StringData(sb.String())
	return unsafe.Slice(pointer, size)
}

func foo2() {
	slice := make([]int, 3, 6)
	fmt.Println("slice:", slice)
	fmt.Printf("&slice: %p\n", &slice)
	fmt.Println("len(slice):", len(slice))
	fmt.Println("cap(slice):", cap(slice))
	fmt.Println()
	
	for i := 0; i < 3; i++ {
		extend(&slice, i+1)
	}
	
	fmt.Println("slice:", slice)
	fmt.Printf("&slice: %p\n", &slice)
	fmt.Println("len(slice):", len(slice))
	fmt.Println("cap(slice):", cap(slice))
}

func extend(s *[]int, n int) {
	*s = append(*s, n)
	fmt.Println("s:", *s)
	fmt.Printf("&s: %p\n", s)
	fmt.Println("len(s):", len(*s))
	fmt.Println("cap(s):", cap(*s))
	fmt.Println()	
}

func foo1() {
	arr := [10]int{}
	fmt.Println("arr:", arr)
	fmt.Println("len(arr):", len(arr))
	fmt.Println("cap(arr):", cap(arr))
}