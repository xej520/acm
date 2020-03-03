package quicksort

import (
	"testing"
	"fmt"
)

func TestQuicksort(b *testing.T) {
	//datas := []int{
	//	9,8,7,6,5,4,3,2,1,
	//}

	//datas := []int{
	//	1,2,3,4,5,6,7,8,9,
	//}

	datas := []int{
		9,8,7,4,5,4,3,94,3,1,22,23,28,25,85,33,39,34,31,30,55,54,59,69,68,61,62,67,77,71,72,79,73,11,13,18,12,19,88,78,89,83,84,81,99,92,93,94,98,96,
	}

	//datas := []int{
	//	6, 2, 1, 7, 4, 5, 9, 8, 3,
	//}

	//datas := []int{
	//	9,9,9,8,8,
	//}

	quickSort(datas)

	for _, d := range datas{
		fmt.Printf("%d ", d)
	}
	fmt.Println()


}

