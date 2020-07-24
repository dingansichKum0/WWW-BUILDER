package queue

type Data int

// ArrayQueue
type ArrayQueue struct {
	Data []Data
	p    int
}

func (q *ArrayQueue) IsEmpty() bool {
	return q.p >= len(q.Data)
}

func (q *ArrayQueue) EnQueue(value Data) {
	q.Data = append(q.Data, value)
}

func (q *ArrayQueue) DeQueue() bool {
	if q.IsEmpty() {
		return false
	}

	q.p++
	return true
}

func (q *ArrayQueue) Front() Data {
	return q.Data[q.p]
}

// CircularQueue
type CircularQueue struct {
	data []Data
	head int
	tail int
	size int
}

func Constructor(k int) CircularQueue {
	return CircularQueue{
		data: make([]Data, k),
		head: -1,
		tail: -1,
		size: k,
	}
}

func (q *CircularQueue) EnQueue(value Data) bool {
	if q.IsFull() {
		return false
	}

	if q.IsEmpty() {
		q.head = 0
	}

	q.tail = (q.tail + 1) % q.size
	q.data[q.tail] = value
	return true
}

func (q *CircularQueue) DeQueue() bool {
	if q.IsEmpty() {
		return false
	}

	if q.head == q.tail {
		q.head = -1
		q.tail = -1
		return true
	}

	q.head = (q.head + 1) % q.size
	return true
}

func (q *CircularQueue) Front() int {
	if q.IsEmpty() {
		return -1
	}

	return q.data[q.head]
}

func (q *CircularQueue) Rear() int {
	if q.IsEmpty() {
		return -1
	}

	return q.data[q.tail]
}

func (q *CircularQueue) IsEmpty() bool {
	return q.head == -1
}

func (q *CircularQueue) IsFull() bool {
	return (q.tail+1)%q.size == -1
}
