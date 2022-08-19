package main

type Node struct {
  val uint8
  next *Node
}

type RingBuffer struct {
  contents []Node
  head uint8
  tail uint8
  length uint8
}

func NewNode(val uint8) *Node {
  n := Node{
    val: val,
    next: nil,
  }
  return &n
}

func NewBuffer(size uint8) *RingBuffer {
  buf := RingBuffer{
    contents: make([]Node, size),
    head: 0,
    tail: 1,
    length: size,
  }
  return &buf
}

func (b RingBuffer) GetContents() []Node {
  return b.contents
}

func (b RingBuffer) GetHead() uint8 {
  return b.head
}

func (b RingBuffer) GetTail() uint8 {
  return b.tail
}

func (b RingBuffer) getLength() uint8 {
  return b.length
}
