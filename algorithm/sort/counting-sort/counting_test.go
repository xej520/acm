package counting_sort


import (
	"testing"
	"fmt"
)

func TestCountingSort(b *testing.T) {
	datas := []int{
		9,8,7,4,5,4,3,3,1,
	}

	//datas := []int{
	//	6, 2, 1, 7, 4, 5, 9, 8, 3,
	//}

	CountingSort(datas)

	for _, d := range datas{
		fmt.Printf("%d ", d)
	}
	fmt.Println()


}

