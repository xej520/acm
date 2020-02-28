package selection_sort

import (
	"testing"
	"fmt"
)

func TestSelectionSort(b *testing.T) {
	datas := []int{
		9,8,7,6,5,4,3,2,1,
	}
	//
	//datas := []int{
	//	6, 2, 1, 7, 4, 5, 9, 8, 3,
	//}

	SelectionSort(datas)

	for _, d := range datas{
		fmt.Printf("%d ", d)
	}
	fmt.Println()

}


