package lesson_1

import (
	"unsafe"
)

func ToLittleEndian(number uint32) uint32 {
	var littleEndian uint32
	littleEndianPtr := unsafe.Pointer(&littleEndian)
	bigEndianPtr := unsafe.Pointer(&number)
	size := int(unsafe.Sizeof(number))
	
	for i := 0; i < size; i++ {
		*(*uint8)(unsafe.Add(littleEndianPtr, i)) = *(*uint8)(unsafe.Add(bigEndianPtr, size-1-i))
	}
	return littleEndian
}



