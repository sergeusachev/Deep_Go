package lesson_2

import(
	"testing"
)

var Result []byte

//go test -bench=. 
func Benchmark_Make(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Result = make([]byte, 0, 10<<20)
	}
}

func Benchmark_MakeDirty(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Result = makeDirty(10<<20)
	}
}