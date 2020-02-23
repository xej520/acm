package sort

import (
	"testing"
	"fmt"
)

func TestSort(t *testing.T) {
	arr := []int{
		6,8,5,4, 3, 2, 1,0,
	}
	arr = sortByMin(arr)
	print(arr)
}

func Benchmark_sort(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		arr := []int{
			6,8,5,4, 3, 2, 1,0,
		}
		arr = sortByMin2(arr)
		print(arr)
	}
}


func sortByMax(arr []int) []int {

	for i := 0; i < len(arr)-1; i++ {
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

	for i := 0; i < len(arr)-1; i++ {
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

func sortByMin2(arr []int) []int {
	for i:=0;i< len(arr)-1; i++ {
		for j:=len(arr)-1; j-1 - i>= 0; j--  {
			if arr[j] > arr[j-1] {
				tmp := arr[j]
				arr[j] = arr[j-1]
				arr[j-1] = tmp
			}
		}
	}

	return arr
}

// 利用Go的多重赋值特性，可以很简单的完成两个变量的交换
func sortByMin3(arr []int) []int {
	for i:=0;i< len(arr)-1; i++ {
		for j:=len(arr)-1; j > i ; j--  {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
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
