package buffer

type RingBuffer struct {
  contents []byte
  head uint8
  tail uint8
  length uint8
  capacity uint8
}

func NewBuffer(size uint8) *RingBuffer {
  b := RingBuffer{
    contents: make([]byte, size),
    head: 0,
    tail: 0,
    length: 0,
    capacity: size,
  }
  return &b
}

// returns the full slice of byte, including empty spots
func (b *RingBuffer) GetContents() []byte {
  return b.contents
}

// returns the index of the current head
func (b *RingBuffer) GetHead() uint8 {
  return b.head
}

// returns the index of the current tail
func (b *RingBuffer) GetTail() uint8 {
  return b.tail
}

// returns the number of currently filled slots in the buffer
func (b *RingBuffer) GetLength() uint8 {
  return b.length
}

// returns the number of currently filled slots in the buffer
func (b *RingBuffer) GetCapacity() uint8 {
  return b.capacity
}

func (b *RingBuffer) Put(n byte) uint8 {
  b.contents[b.tail] = n
  b.tail++
  b.length++
  return b.tail
}

func (b *RingBuffer) Pop() byte {
  n := b.contents[b.head]
  b.contents[b.head] = 0
  b.head++
  b.length--
  return n
}
