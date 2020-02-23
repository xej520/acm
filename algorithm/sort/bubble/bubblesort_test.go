package bubble

import (
	"testing"
)

const num = 1000

func small2big() []int {
	da := make([]int, num)

	for i := 0; i < num; i++ {
		da[i] = i
	}

	return da
}

func big2small() []int {
	da := make([]int, num)

	for i := 0; i < num; i++ {
		da[num-i-1] = i
	}

	return da
}

func createData(f func() []int) []int {
	return f()
}

func Benchmark_SortByMin2(b *testing.B) {
	var arr = createData(small2big)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		arr = sortByMin(arr)
	}

	b.StopTimer()
	//print(arr)
}
