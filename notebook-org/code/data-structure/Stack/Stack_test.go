package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	s := new(Stack)

	if !s.IsEmpty() {
		t.Error("stack should be empty.")
	}

	s.Push(1)
	s.Push(2)
	s.Push(3)

	if len(s.data) != 3 {
		t.Error("stack length should be 2.")
	}

	s.Pop()

	if s.Top() != 2 {
		t.Error("3 has be removed.")
	}

	ms := Constructor()

	ms.Push(-3)
	ms.Push(1)
	ms.Push(-2)
	ms.Push(-1)

	ms.Pop()

	if ms.Top() != -2 {
		t.Error("ms.Top() should be -2")
	}

	if ms.GetMin() != -3 {
		t.Error("ms.GetMin() should be -3")
	}

}
