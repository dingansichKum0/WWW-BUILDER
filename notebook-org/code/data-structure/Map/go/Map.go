package Map

type K interface{}

type V interface{}

type Maper interface {
	Remove(K)
	Contains(K) bool
	Get(K) V
	Set(K, V)
	IsEmpty() bool
	Range(func(K, V))
}

type Node struct {
	Key   K
	Value V
	next  *Node
}

type LinkedListMap struct {
	head *Node
	Size int
}

// IsEmpty ...
func (m *LinkedListMap) IsEmpty() bool {
	return m.Size == 0
}

// get ...
func (m *LinkedListMap) get(key K) *Node {
	p := m.head
	for p != nil {
		if key == p.Key {
			return p
		}
		p = p.next
	}
	return nil
}

// Contains ...
func (m *LinkedListMap) Contains(key K) bool {
	return m.get(key) != nil
}

// Get ...
func (m *LinkedListMap) Get(key K) V {
	p := m.get(key)
	if p != nil {
		return p.Value
	}
	return nil
}

// Set ...
func (m *LinkedListMap) Set(key K, value V) {
	p := m.get(key)
	if p != nil {
		p.Value = value
		return
	}

	if m.head == nil {
		m.head = new(Node)
		m.head.next = &Node{key, value, nil}
	} else {
		m.head.next = &Node{key, value, m.head.next}
	}
	m.Size++
}

// Remove ...
func (m *LinkedListMap) Remove(key K) {
	p := m.head
	if p == nil {
		return
	}

	for p.next != nil {
		if key == p.next.Key {
			break
		}
		p = p.next
	}

	if p.next != nil {
		del := p.next
		p.next = del.next
		del = nil
		m.Size--
	}

}

// Range ...
func (m *LinkedListMap) Range(fn func(K, V)) {
	p := m.head.next
	for p != nil {
		fn(p.Key, p.Value)
		p = p.next
	}
}
