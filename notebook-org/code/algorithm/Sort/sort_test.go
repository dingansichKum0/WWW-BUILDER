package sort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	a := []int{43, 2, 22, 15, 66, 0, 214}
	QuickSort(a, 0, len(a)-1)

	fmt.Println(a)
}

func TestMergeSort(t *testing.T) {
	a := []int{43, 2, 22, 15, 66, 0, 214}
	MergeSort(a)

	fmt.Println(a)
}
