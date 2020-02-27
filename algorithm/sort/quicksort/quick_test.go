package quicksort

import (
	"testing"
	"fmt"
)

func TestQuicksort(b *testing.T) {
	//datas := []int{
	//	9,8,7,6,5,4,3,2,1,
	//}

	datas := []int{
		6, 2, 1, 7, 4, 5, 9, 8, 3,
	}

	quickSort(datas)

	for _, d := range datas{
		fmt.Printf("%d ", d)
	}
	fmt.Println()


}

