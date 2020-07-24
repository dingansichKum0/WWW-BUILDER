package LinkedList

type Data interface{}

type LinkedLister interface {
	IsEmpty() bool
	Contains(int) bool
	Get(int) Data
	Insert(int, Data)
	Remove(int)
	Modify(int, Data)
	PrintAll(func(int, Data))
}

type Node struct {
	Data
	next *Node
}

type LinkedList struct {
	Len  int
	head *Node
}

// checkIndex ...
func (l *LinkedList) checkIndex(index int) {
	if index > l.Len || index < 0 {
		panic("error: 索引越界.")
	}
}

// IsEmpty ...
func (l *LinkedList) IsEmpty() bool {
	return l.Len == 0
}

// Get ...
func (l *LinkedList) Get(index int) Data {
	return l.get(index).Data
}

// get ...
func (l *LinkedList) get(index int) *Node {
	l.checkIndex(index)

	p := l.head
	i := index
	for i < index {
		p = p.next
		i++
	}

	return p
}

// Contains ...
func (l *LinkedList) Contains(index int) bool {
	return l.get(index) != nil
}

// Insert ...
func (l *LinkedList) Insert(index int, value Data) {
	l.checkIndex(index)

	n := new(Node)
	n.Data = value

	p := l.head

	if index == 0 {
		l.head = n
		l.head.next = p
		l.Len++
		return
	}

	i := 1
	for p != nil && i < index {
		p = p.next
		i++
	}

	if p == nil {
		panic("error: 链表中存在空指针")
	}

	n.next = p.next
	p.next = n
	l.Len++
}

// Remove ...
func (l *LinkedList) Remove(index int) {
	l.checkIndex(index)

	p := l.head
	for index > 1 {
		p = p.next
		index--
	}

	p.next = p.next.next
	l.Len--
}

// Modify ...
func (l *LinkedList) Modify(index int, value Data) {
	l.get(index).Data = value
}

// PrintAll ...
func (l *LinkedList) PrintAll(fn func(index int, value Data)) {
	p := l.head
	idx := 0
	for i := l.Len; i > 0; i-- {
		fn(idx, p.Data)
		p = p.next
		idx++
	}
}
