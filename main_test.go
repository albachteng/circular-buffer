package main

import (
  "testing"
)

func TestNewNode(t *testing.T) {
  emptyNode := *NewNode(0)
  val := emptyNode.val
  if val != 0 {
    t.Errorf("expected val to be 0, got %d", val)
  }
}

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
