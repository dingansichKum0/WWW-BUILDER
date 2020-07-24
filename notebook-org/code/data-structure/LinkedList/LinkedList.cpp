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
