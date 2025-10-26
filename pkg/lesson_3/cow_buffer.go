// Предположим, что все будут производить копирование
// буффера только с использованием метода Clone()
/*
type COWBuffer struct { ... }

func NewCOWBuffer(data []byte)                         // создать буффер с определенными данными
func (b *COWBuffer) Clone() COWBuffer                  // создать новую копию буфера
func (b *COWBuffer) Close()                            // перестать использовать копию буффера
func (b *COWBuffer) Update(index int, value byte) bool // изменить определенный байт в буффере
func (b *COWBuffer) String() string                    // сконвертировать буффер в строку 
*/

package lesson_3

import(
	"unsafe"
)

type COWBuffer struct {
	buffer []byte
	refCount int
}

func NewCOWBuffer(data []byte) *COWBuffer {
	return &COWBuffer{
		buffer: data,
		refCount: 0,
	}
}

func (b *COWBuffer) Clone() COWBuffer {
	return COWBuffer{
		buffer: b.buffer,
		refCount: b.refCount++,
	}
}
                  
func (b *COWBuffer) Close() {
	b.copyBuffer()
}   
                         
func (b *COWBuffer) Update(index int, value byte) bool {
	if b.refCount > 0 {
		b.copyBuffer()
	}
	
	b.buffer[index] = value
	
	return true
} 

func (b *COWBuffer) copyBuffer() {
	newBuffer := make([]byte, len(b.buffer))
	copy(newBuffer, b.buffer)
	b.buffer = newBuffer
	b.refCount = 0
}

func (b *COWBuffer) String() string {
	if len(b.buffer) == 0 {
		return ""
	}
	
	return unsafe.String(unsafe.SliceData(b.buffer), len(b.buffer))
}                     


