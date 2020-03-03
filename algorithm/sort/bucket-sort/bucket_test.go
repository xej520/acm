package bucket_sort

import (
	"testing"
	"fmt"
)

func TestBucketSort(b *testing.T) {
	datas := []int{
		9,8,7,4,5,4,3,3,1,22,23,28,77,78,57,25,33,39,34,44,89,22,31,30,55,54,66,88,77,66,89,59,69,68,67,63,61,62,67,77,71,72,79,73,77,88,73,11,13,18,12,19,88,78,89,83,84,81,99,92,93,94,98,96,
	}

	BucketSort(datas)

	for _, d := range datas{
		fmt.Printf("%d ", d)
	}
	fmt.Println()
}


