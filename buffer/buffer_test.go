package buffer

import (
	"testing"
)

func TestNewBuffer(t *testing.T) {
  emptyBuffer := NewBuffer(5)
  contents := emptyBuffer.contents
  if len(contents) != 5 {
    t.Errorf("expected contents to have length 5, got %d", len(contents))
  }
}

func TestGetHead(t *testing.T) {
  emptyBuffer := NewBuffer(5)
  head := emptyBuffer.GetHead()
  if head != 0 {
    t.Errorf("expected head to be 0, got %d", head)
  }
}

func TestGetTail(t *testing.T) {
  emptyBuffer := NewBuffer(5)
  tail := emptyBuffer.GetTail()
  if tail != 0 {
    t.Errorf("expected tail to be 0, got %d", tail)
  }
}

func TestGetLength(t *testing.T) {
  emptyBuffer := NewBuffer(5)
  length := emptyBuffer.GetLength()
  if length != 0 {
    t.Errorf("expected length to be 0, got %d", length)
  }
}

func TestGetCapacity(t *testing.T) {
  emptyBuffer := NewBuffer(5)
  capacity := emptyBuffer.GetCapacity()
  if capacity != 5 {
    t.Errorf("expected length to be 5, got %d", capacity)
  }
}

func TestPut(t *testing.T) {
  oneItemBuffer := NewBuffer(5)
  oneItemBuffer.Put(1)
  val := oneItemBuffer.contents[0]
  if val != 1 {
    t.Errorf("expected val to be 1, got %d", val)
  }
  oneItemBuffer.Put(2)
  val = oneItemBuffer.contents[1]
  if val != 2 {
    t.Errorf("expected val to be 2, got %d", val)
  }
}

func TestPop(t *testing.T) {
  twoItemBuffer := NewBuffer(5)
  twoItemBuffer.Put(1)
  twoItemBuffer.Put(2)
  popped := twoItemBuffer.Pop()
  if popped != 1 {
    t.Errorf("expected popped to be 1, got %d", popped)
  }
  popped = twoItemBuffer.Pop()
  if popped != 2 {
    t.Errorf("expected popped to be 2, got %d", popped)
  }
}
