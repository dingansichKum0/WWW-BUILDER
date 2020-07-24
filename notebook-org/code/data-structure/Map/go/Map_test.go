package Map

import (
	"testing"
)

func TestLinkedListMap(t *testing.T) {
	m := new(LinkedListMap)

	sample := [][2]interface{}{
		{"a", "A"},
		{"b", "B"},
		{"c", "C"},
		{"d", "D"},
		{1, 1},
	}

	for _, v := range sample {
		m.Set(v[0], v[1])
	}

	if m.Get("a") != "A" {
		t.Error("a!=A")
	}

	if m.Get("b") != "B" {
		t.Error("b!=B")
	}

	if m.Get(1) != 1 {
		t.Error("1!=1")
	}

	m.Remove("c")

	if m.Get("c") != nil {
		t.Error("key c should be nil")
	}
}

func BenchmarkTestLinkedListMap(b *testing.B) {
	m := new(LinkedListMap)

	for i := 0; i < b.N; i++ {

		sample := [][2]interface{}{
			{"a", "A"},
			{"b", "B"},
			{"c", "C"},
			{"d", "D"},
			{1, 1},
		}

		for _, v := range sample {
			m.Set(v[0], v[1])
		}

		if m.Get("a") != "A" {
			b.Error("a!=A")
		}

		if m.Get("b") != "B" {
			b.Error("b!=B")
		}

		if m.Get(1) != 1 {
			b.Error("1!=1")
		}
	}
}
