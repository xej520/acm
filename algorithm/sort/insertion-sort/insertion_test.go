package insertion_sort

import (
	"testing"
	"fmt"
)

func TestInsertSort(b *testing.T) {
	datas := []int{
		9,8,7,6,5,4,3,2,1,4,
	}
	//
	//datas := []int{
	//	6, 2, 1, 7, 4, 5, 9, 8, 3,
	//}

	InsertionSort(datas)

	for _, d := range datas{
		fmt.Printf("%d ", d)
	}
	fmt.Println()

}

