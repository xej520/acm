package fangshi1

import (
	"math"
)

type sortFunc func([]int)

func RadixSort(data []int) {
	// 1、求出数据中，最长位数是多少
	ml := maxLength(data)

	// 2、根据最长位数，生成一组桶
	b := newBucket(ml)

	// 3、将待测试数据灌入到桶里，
	// 灌入思路：数据位数是1位的，放入第一个桶里，数据位数是2位的，灌入第2个桶里
	// 依次类推
	data2Bucket(data, b)

	// 4、对每个桶，进行排序，每个桶是独立的，可以采取不同的排序算法
	sortForBucket(b, []sortFunc{quickSort, countingSort})

	// 5、将非空桶的数据，重新导入到原测试数据里
	bucket2data(data, b)
}

func maxLength(data []int) int {
	var ml int
	for _, v := range data {
		t := (int)(math.Log10(float64(v))) + 1
		if ml < t {
			ml = t
		}
	}

	return ml
}

func newBucket(bucketNum int) map[int][]int {

	var b  = make(map[int][]int)
	for i := 0; i < bucketNum; i++ {
		b[i] = []int{}
	}

	return b

}

func data2Bucket(data []int, b map[int][]int) {
	for _, v := range data {
		l := (int)(math.Log10(float64(v)))
		b[l] = append(b[l], v)
	}
}

func sortForBucket(b map[int][]int, f []sortFunc) {
	var index int
	for _, v := range b {
		f[index](v)
		index++
		if index == len(f) {
			index = 0
		}
	}
}

func bucket2data(data []int, b map[int][]int) {
	var index int
	for _, v := range b {
		if len(v) != 0 {
			for _, d := range v {
				data[index] = d
				index++
			}
		}
	}
}

func quickSort(data []int) {
	subQuickSort(data, 0, len(data)-1)
}

func subQuickSort(data []int, l, r int) {
	if r <= l {
		return
	}

	pivot := data[l]

	var left, right = l, r

	for left < right {

		for left < right && pivot <= data[right] {
			right--
		}

		for left < right && data[left] <= pivot {
			left++
		}

		if left < right {
			data[left], data[right] = data[right], data[left]
		}

	}

	data[left], data[l] = pivot, data[left]

	// 右边分区
	subQuickSort(data, right+1, r)

	// 左边分区
	subQuickSort(data, l, right-1)

}

func countingSort(data []int) {
	if len(data) <= 1 {
		return
	}

	min, max := countMaxMin(data)

	temp := make([]int, max+1)

	for i := 0; i < len(data); i++ {
		temp[data[i]]++
	}
	var index int
	for i := min; i < len(temp); i++ {
		for j := temp[i]; j > 0; j-- {
			data[index] = i
			index++
		}
	}
}

func countMaxMin(data []int) (int, int) {
	min, max := data[0], data[0]

	for i := 1; i < len(data); i++ {
		if min > data[i] {
			min = data[i]
		}

		if max < data[i] {
			max = data[i]
		}

	}

	return min, max
}
