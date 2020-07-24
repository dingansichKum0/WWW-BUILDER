package stack

type Data interface{}

type Stack struct {
	data []Data
}

func (s *Stack) Push(value Data) {
	s.data = append(s.data, value)
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack) Top() Data {
	if s.IsEmpty() {
		return -1
	}
	return s.data[len(s.data)-1]
}

func (s *Stack) Pop() bool {
	if s.IsEmpty() {
		return false
	}

	s.data = s.data[:len(s.data)-1]
	return true
}

type MinStack struct {
	data []Data
	min  *Stack
}

func Constructor() MinStack {
	return MinStack{
		data: make([]Data, 0),
		min:  new(Stack),
	}
}

func (s *MinStack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *MinStack) Push(value Data) {
	if s.min.IsEmpty() {
		s.min.Push(value)
	} else {
		if value.(int) <= s.min.Top().(int) {
			s.min.Push(value)
		}
	}

	s.data = append(s.data, value)
}

func (s *MinStack) Pop() {
	if s.Top() == s.min.Top() {
		s.min.Pop()
	}

	if s.IsEmpty() {
		return
	}

	s.data = s.data[:len(s.data)-1]
}

func (s *MinStack) Top() Data {
	if s.IsEmpty() {
		return -1
	}
	return s.data[len(s.data)-1]
}

func (s *MinStack) GetMin() Data {
	return s.min.Top()
}

// "{[]}"
// "()[]{}"
// "([)]"
func isValid(s string) bool {
	l := len(s)
	sk := new(Stack)

	for i := 0; i < l; i++ {
		ss := string(s[i])
		if ss == "[" || ss == "{" || ss == "(" {
			sk.Push(ss)
		}

		if ss == "]" {
			if sk.Top() == "[" {
				sk.Pop()
				continue
			}
			return false

		}

		if ss == "}" {
			if sk.Top() == "{" {
				sk.Pop()
				continue
			}
			return false
		}

		if ss == ")" {
			if sk.Top() == "(" {
				sk.Pop()
				continue
			}
			return false
		}
	}

	return sk.IsEmpty()
}
