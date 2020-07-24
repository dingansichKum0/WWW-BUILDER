#include <iostream>
using namespace std;

typedef int Data;

class List {
public:
  List(int len = 10) : len(len)
  {
    used = 0;
    p = (Data *)malloc(sizeof(Data) * len);
  }
  ~List() { cout << "delete" << endl; };
  int len;
  int used;
  int *p;

  bool is_full();
  bool is_empty();
  int get(int);
  void insert(int, Data);
  void remove(int);
  void print_all();

private:
  void check_index(int);
};

bool List::is_full() { return len == used; }

bool List::is_empty() { return used == 0; }

void List::check_index(int index)
{
  if (index >= len || index < 0) {
    printf("error: 索引越界.");
    exit(-1);
  }
}

int List::get(int index)
{
  check_index(index);

  return p[index];
}

void List::insert(int index, Data value)
{
  check_index(index);
  if (index > used || index < 0) {
    printf("error: 索引越界.");
    return;
  }

  int i = used;
  while (i - index > 0) {
    i--;
    p[i] = p[i - 1];
  }
  p[index] = value;
  used++;
}

void List::remove(int index)
{
  check_index(index);
  if (index >= used) {
    printf("error: 索引越界.");
    return;
  }

  used--;

  for (int i = index; i < used; i++) {
    p[i] = p[i + 1];
  }
}

void List::print_all()
{
  for (int i = 0; i < used; i++) {
    cout << get(i) << endl;
  }
}

int main()
{
  List *l = new List(5);

  l->insert(0, 5);
  l->insert(1, 4);
  l->insert(2, 3);
  l->insert(0, 1);

  l->print_all();

  // delete l;
  cout << l->used << endl;

  return 0;
}
