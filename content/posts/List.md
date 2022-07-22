+++
title = "List"
author = ["zakudriver"]
date = 2020-03-03T00:00:00+08:00
draft = false
+++

```c++
#include <iostream>
#include <stack>
#include <stdlib.h>
using namespace std;

class List {
public:
  List(int len = 10) : len(len) {
    used = 0;
    p = (int *)malloc(sizeof(int) * len);
  }
  int len;
  int used;
  int *p;

  bool is_full();
  bool is_empty();
  void push(int);
  int get(int);
  void insert(int, int);
  void remove(int);
  void print_all();

private:
  void check_index(int);
};

bool List::is_full() { return len == used; }

bool List::is_empty() { return used == 0; }

void List::check_index(int index) {
  if (index > len - 1 || index < 0) {
    printf("error: 索引越界.");
    exit(-1);
  }
}

void List::push(int value) {
  if (used >= len) {
    printf("error: 链表已满.");
  }

  p[used] = value;
  used += 1;
};

int List::get(int index) {
  check_index(index);

  return p[index];
}

void List::insert(int index, int value) {
  check_index(index);
  if (index > used) {
    printf("error: 索引越界.");
    exit(-1);
  }

  while (used - index > 0) {
    used--;
    p[used] = p[used - 1];
  }
  p[index] = value;
  used += 1;
}

void List::remove(int index) {
  check_index(index);
  if (index >= used) {
    printf("error: 索引越界.");
    exit(-1);
  }

  used--;

  for (int i = index; i < used; i++) {
    p[i] = p[i + 1];
  }
}

void List::print_all() {
  for (int i = 0; i < used; i++) {
    cout << get(i) << endl;
  }
}

int main() {
  List l(5);

  l.push(10);
  l.push(20);

  l.insert(2, 5);
  l.remove(0);

  l.print_all();

  return 0;
}
```
