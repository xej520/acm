package main

import "fmt"

type sortF func(int, int) bool

func sort(arr []int, f sortF) []int {

	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if f(arr[j], arr[j+1]) {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	return arr
}

// 如果按照从小到大排序
// 那么，相邻两个数(前后两个数)交换的条件是
// 前面的数，大于了， 后面的数，
// 需要进行交换了
func sortBySmall(a, b int) bool {
	return a > b
}

func sortByBig(a, b int) bool {
	return a < b
}

func main() {
	data := []int{6, 5, 4, 2, 9, 1, 0, 4, 2, 3, 11}

	sort(data, sortByBig)

	for _, d := range data {
		fmt.Printf("%d ", d)
	}

	fmt.Println()
}
