package bytes

func NewBuffer(size int) *Buffer {
	return &Buffer{
		B: make([]byte, 0, size),
	}
}

type Buffer struct {
	B []byte
}

func (this *Buffer) Len() int {
	return len(this.B)
}

func (this *Buffer) Cap() int {
	return cap(this.B)
}

func (this *Buffer) String() string {
	if this == nil { // Special case, useful in debugging.
		return "<nil>"
	}
	return string(this.B)
}

func (this *Buffer) Bytes() []byte {
	return this.B
}

func (this *Buffer) Reset() {
	this.B = this.B[:0]
}

func (this *Buffer) Write(p []byte) (n int, err error) {
	this.B = append(this.B, p...)
	return len(p), nil
}

func (this *Buffer) WriteByte(c byte) error {
	this.B = append(this.B, c)
	return nil
}
