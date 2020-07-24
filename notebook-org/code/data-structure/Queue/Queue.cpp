#include <iostream>
#include <vector>
using namespace std;

typedef int Data;

class ArrayQueue {
private:
  vector<Data> data;
  int p_start;

public:
  ArrayQueue() { p_start = 0; }
  bool enQueue(int x)
  {
    data.push_back(x);
    return true;
  }

  bool deQueue()
  {
    if (isEmpty()) {
      return false;
    }
    p_start++;
    return true;
  };

  int Front() { return data[p_start]; };

  bool isEmpty() { return p_start >= data.size(); }
};

class CircularQueue {
private:
  vector<Data> data;
  int head;
  int tail;
  int size;

public:
  CircularQueue(int k)
  {
    data.resize(k);
    head = -1;
    tail = -1;
    size = k;
  }

  bool enQueue(int value)
  {
    if (isFull()) {
      return false;
    }
    if (isEmpty()) {
      head = 0;
    }
    tail = (tail + 1) % size;
    data[tail] = value;
    return true;
  }

  bool deQueue()
  {
    if (isEmpty()) {
      return false;
    }
    if (head == tail) {
      head = -1;
      tail = -1;
      return true;
    }
    head = (head + 1) % size;
    return true;
  }

  int Front()
  {
    if (isEmpty()) {
      return -1;
    }
    return data[head];
  }

  int Rear()
  {
    if (isEmpty()) {
      return -1;
    }
    return data[tail];
  }

  bool isEmpty() { return head == -1; }

  bool isFull() { return ((tail + 1) % size) == head; }
};

int main()
{
  ArrayQueue q;
  q.enQueue(5);
  q.enQueue(3);
  if (!q.isEmpty()) {
    cout << q.Front() << endl;
  }
  q.deQueue();
  if (!q.isEmpty()) {
    cout << q.Front() << endl;
  }
  q.deQueue();
  if (!q.isEmpty()) {
    cout << q.Front() << endl;
  }
}
