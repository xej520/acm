package sort

import (
	"testing"
	"fmt"
)

func TestSort(t *testing.T) {
	arr := []int{
		2, 3, 4, 1,
	}
	arr = sortByMin(arr)
	print(arr)
}

func sortByMax(arr []int) []int {

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] < arr[j+1] {
				tmp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			}
		}
	}

	return arr
}

func sortByMin(arr []int) []int {

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				tmp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			}
		}
	}

	return arr
}

func print(arr []int) {

	for _, v := range arr {
		fmt.Printf("%d ", v)
	}

	fmt.Println()
}
