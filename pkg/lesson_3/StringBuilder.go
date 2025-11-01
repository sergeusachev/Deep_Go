package lesson_3

type Builder struct {
	buffer []byte
}

func NewBuilder() Builder {
	return Builder{}
}

func (b *Builder) Grow(capacity int) {
	if capacity < 0 {
		return
	}
	
	if capacity < len(b.buffer) {
		b.buffer = b.buffer[:capacity]
		return
	}
	
	buffer := make([]byte, len(b.buffer), capacity)
	copy(buffer, b.buffer)
	b.buffer = buffer
}

func (b *Builder) Write(symbol byte) {
	b.buffer = append(b.buffer, symbol)
}

func (b *Builder) String() string {
	return string(b.buffer)
}