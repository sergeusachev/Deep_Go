package lesson_1

import (
    "fmt"
    "math"
    "errors"
    "unsafe"
    "runtime"
)

func HasDuplicatesFrom1To7WithBits(data []int) bool {
    var lookup int8
    for _, number := range data {
        if lookup & (1<<number) == 1 {
            return true
        }
        lookup = lookup | (1<<number)
    }
    return false
}

func HasDuplicatesFrom1To7WithArray(data []int) bool {
    var lookup[8]int8
    for _, number := range data {
        if lookup[number] == 1 {
            return true
        }
        lookup[number] = 1
    }
    return false
}

func HasDuplicatesFrom1To7WithHashTable(data []int) bool {
    lookup := make(map[int]struct{}, 8)
    for _, number := range data {
        _, found := lookup[number]
        if found {
            return true
        }
        lookup[number] = struct{}{}
    }
    return false
}

func isPowerofTwo(value int) bool {
    return value > 0 && value & (value-1) == 0
}

func bitShift() {
    var x int8 = 0b00000001
    fmt.Println("decimal x:", x)
    fmt.Printf("binary x: %08b\n", x)
    
    fmt.Println("decimal x << 1:", x << 1)
    fmt.Printf("binary x << 1: %08b\n", x << 1)
    
    fmt.Println("decimal x >> 1:", x >> 1)
    fmt.Printf("binary x >> 1: %08b\n", x >> 1)
}

func isLittleEndian() bool {
    var number int16 = 0x0001   
    pointer := (*int8)(unsafe.Pointer(&number))
    return *pointer == 1
}

func isBigEndian() bool {
    return !isLittleEndian()
}

func endianessTrick() {
    var number int32 = 0x12345678
    pointer := unsafe.Pointer(&number)
    
    fmt.Printf("0x")
    for i := 0; i < 4; i++ {
        byteValue := *(*int8)(unsafe.Add(pointer, i))
        fmt.Printf("%x", byteValue)
    }
    fmt.Println()
}

func uintptrTest() {
    x := new(int)
    y := new(int)
    
    ptrX := unsafe.Pointer(x)
    addressY := uintptr(unsafe.Pointer(y))
    
    runtime.GC()
    
    *(*int)(ptrX) = 100
    *(*int)(unsafe.Pointer(addressY)) = 300 //dangerous
}

func sizeofCheck() {
    var value1 int8 = 10
    fmt.Println("size1:", unsafe.Sizeof(value1)) //compile time
    
    var value2 int32 = 10
    fmt.Println("size2:", unsafe.Sizeof(value2)) //compile time
}

func unsafePointer() {
    var value uint32 = 0xFFFFFFFF
    
    pointer := unsafe.Pointer(&value)
    bytePointer := (*uint8)(pointer)
    fmt.Println("value1:", *bytePointer)
    
    pointer = unsafe.Add(pointer, 2)
    twoBytePointer := (*uint16)(pointer)
    fmt.Println("value2:", *twoBytePointer)
}

func pointersTricky() {
    var v1 int32 = 200
    p := &v1
    fmt.Println("value:", *p)
    fmt.Println("address:", p)
     
    process(&p)
    fmt.Println("value:", *p)
    fmt.Println("address:", p)
}

func process(temp **int32) {
    var v2 int32 = 100
    *temp = &v2
}

func pointers() {
    var value int32 = 100
    var pointer *int32 = &value
    
    fmt.Println("address:", pointer)
    fmt.Println("value:", *pointer)
    
    *pointer = 500
    fmt.Println("address:", pointer)
    fmt.Println("value:", *pointer)
}

func strangeSum() {
    sum := 100 + 010
    fmt.Println(sum)
}

func checkIncrementOverflow() {
    counter := math.MaxInt
    counter, err := inc(counter)
    fmt.Println("counter =", counter, " err =", err)
}

func inc(counter int) (int, error) {
    if counter == math.MaxInt {
        return 0, errors.New("int overflow")
    }
    return counter + 1, nil
}

func intsOverflows() {
    var signed int8 = math.MaxInt8
    signed++
    fmt.Println(signed)
    
    var unsigned uint8 = math.MaxUint8
    unsigned++
    fmt.Println(unsigned)
}