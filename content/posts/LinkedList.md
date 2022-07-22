+++
title = "LinkedList"
author = ["zakudriver"]
date = 2020-03-05T00:00:00+08:00
draft = false
+++

## C++ {#c-plus-plus}

```c++
#include <iostream>
using namespace std;

typedef int Data;

struct Node {
  Data data;
  Node *next;
  Node(Data d = 0, Node *n = NULL) : data(d), next(n){};
};

class LinkedList {
public:
  LinkedList(int len) : len(len), head(NULL){};

  bool is_empty();
  Data get(int);
  void insert(int, Data);
  void remove(int);
  void modify(int, Data);
  void print_all();

private:
  Node *head;
  int len;
  void check_index(int);
  Node *get_by_index(int);
};

void LinkedList::check_index(int index)
{
  if (index > len || index < 0) {
    printf("error: 索引越界.");
    exit(-1);
  }
}

Node *LinkedList::get_by_index(int index)
{
  check_index(index);

  Node *p = head;
  int i = 0;
  while (i++ < index) {
    p = p->next;
  }

  return p;
}

bool LinkedList::is_empty() { return len == 0; }

Data LinkedList::get(int index) { return get_by_index(index)->data; }

void LinkedList::insert(int index, Data value)
{
  check_index(index);

  Node *n = new Node(value);
  Node *p = head;

  if (index == 0) {
    head = n;
    head->next = p;
    len++;
    return;
  }

  int i = 1;
  while (p != NULL && i++ < index) {
    p = p->next;
  }

  if (p == NULL) {
    printf("error: 链表中存在空指针");
    return;
  }

  n->next = p->next;
  p->next = n;
  len++;
}

void LinkedList::remove(int index)
{
  check_index(index);

  Node *p = head;
  while (index-- > 1) {
    p = p->next;
  }

  p->next = p->next->next;
  delete p->next;
  len--;
}

void LinkedList::modify(int index, Data value)
{
  get_by_index(index)->data = value;
}

void LinkedList::print_all()
{
  Node *p = head;
  int i = len;
  while (i-- > 0) {
    cout << p->data << endl;
    p = p->next;
  }
}

int main()
{
  LinkedList *l = new LinkedList(0);

  l->insert(0, 1);
  l->insert(1, 2);
  l->insert(1, 3);
  l->insert(0, 4);

  l->modify(3, 10);

  l->print_all();

  return 0;
}
```


## Go {#go}

```go
package main

import "fmt"

type Data int

type Maper interface {
  Is_empty() bool
  Get(int) Data
  Insert(int, Data)
  Remove(int)
  Modify(int, Data)
  Print_all()
}

type Node struct {
  Data
  next *Node
}

type LinkedList struct {
  Len  int
  head *Node
}

// check_index ...
func (l *LinkedList) check_index(index int) {
  if index > l.Len || index < 0 {
    panic("error: 索引越界.")
  }
}

// Is_empty ...
func (l *LinkedList) Is_empty() bool {
  return l.Len == 0
}

// Get ...
func (l *LinkedList) Get(index int) Data {

  return l.get(index).Data
}

// get ...
func (l *LinkedList) get(index int) *Node {
  l.check_index(index)

  p := l.head
  i := index
  for i < index {
    p = p.next
    i++
  }

  return p
}

// Insert ...
func (l *LinkedList) Insert(index int, value Data) {
  l.check_index(index)

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
  l.check_index(index)

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

// Print_all ...
func (l *LinkedList) Print_all() {
  p := l.head
  for i := l.Len; i > 0; i-- {
    fmt.Println(p.Data)
    p = p.next
  }

}

func main() {
  var l = new(LinkedList)

  l.Insert(0, 1)
  l.Insert(1, 2)
  l.Insert(1, 3)

  l.Print_all()

}
```
