package lesson_3

import(
	"testing"
	"strings"
)

func BenchmarkSimpleConcat(b *testing.B) {
	str := "test"
	for i := 0; i < b.N; i++ {
		str += "_test"
	}
	_ = str
}

func BenchmarkStringBuilderConcat(b *testing.B) {
	builder := strings.Builder{}
	builder.WriteString("test")
	for i := 0; i < b.N; i++ {
		builder.WriteString("_test")
	}
	_ = builder.String()
}

func BenchmarkStringBuilderConcatOptimized(b *testing.B) {
	builder := strings.Builder{}
	builder.Grow(4 + b.N*5)
	builder.WriteString("test")
	for i := 0; i < b.N; i++ {
		builder.WriteString("_test")
	}
	_ = builder.String()
}