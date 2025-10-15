package main

import (
    "testing"
)

func BenchmarkHasDuplicatesFrom1To7WithHashTable(b *testing.B) {
    data := []int{1, 2, 3, 4, 5 , 6, 7}
    for i := 1; i < b.N; i++ {
        _ = HasDuplicatesFrom1To7WithHashTable(data)
    }
}

func BenchmarkHasDuplicatesFrom1To7WithArray(b *testing.B) {
    data := []int{1, 2, 3, 4, 5 , 6, 7}
    for i := 1; i < b.N; i++ {
        _ = HasDuplicatesFrom1To7WithArray(data)
    }
}

func BenchmarkHasDuplicatesFrom1To7WithBits(b *testing.B) {
    data := []int{1, 2, 3, 4, 5 , 6, 7}
    for i := 1; i < b.N; i++ {
        _ = HasDuplicatesFrom1To7WithBits(data)
    }
}