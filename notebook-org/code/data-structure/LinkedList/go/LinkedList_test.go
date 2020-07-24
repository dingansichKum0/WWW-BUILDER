package LinkedList

import (
	"fmt"
	"testing"
)

func TestLinkedList(t *testing.T) {
	sample := [][2]interface{}{
		{1, 1},
		{2, 2},
		{3, 3},
	}

	var l LinkedLister = new(LinkedList)

	// for k, v := range sample {
	// 	l.Insert(k, v[0])
	// }

	l.Insert(0, 0)
	l.Insert(0, 1)
	l.Insert(0, 2)
	fmt.Println(l.Get(1))

	for k, _ := range sample {

		fmt.Println(l.Get(k))

		// if l.Get(k) != v[1] {
		// 	// t.Errorf("%s != %s", v[0], v[1])
		// }

	}

}
