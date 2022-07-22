+++
title = "Stack"
author = ["zakudriver"]
date = 2020-03-13T00:00:00+08:00
draft = false
+++

## c++ {#c-plus-plus}

```c++
#include <iostream>
#include <vector>
using namespace std;

class Stack {
private:
  vector<int> data;

public:
  void push(int x) { data.push_back(x); }

  bool isEmpty() { return data.empty(); }

  int top() { return data.back(); }

  bool pop()
  {
    if (isEmpty()) {
      return false;
    }
    data.pop_back();
    return true;
  }
};

int main()
{
  Stack s;
  s.push(1);
  s.push(2);
  s.push(3);
  for (int i = 0; i < 4; ++i) {
    if (!s.isEmpty()) {
      cout << s.top() << endl;
    }
    cout << (s.pop() ? "true" : "false") << endl;
  }
}
```


## go {#go}

```go
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
```
