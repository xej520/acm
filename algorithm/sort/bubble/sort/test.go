package main

import (
	"fmt"
)

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

// 利用Go的多重赋值特性，可以很简单的完成两个变量的交换
func sortByMax2(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := len(arr) - 1; j > i; j-- {
			if arr[j-1] < arr[j] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}

	return arr
}

func sortByMin(arr []int) []int {

	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}

	return arr
}

func sortByMin2(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := len(arr) - 1; j-1-i >= 0; j-- {
			if arr[j-1] > arr[j] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}

	return arr
}

func print(arr []int) {

	fmt.Println()
	for _, v := range arr {
		fmt.Printf("%d ", v)
	}

	fmt.Println()
}

func main() {

}
