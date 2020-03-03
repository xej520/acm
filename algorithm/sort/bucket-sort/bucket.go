package bucket_sort

import (
	"fmt"
)

var bucketScope = 10

type SortFunc func([]int)

var sortFuncSlice = []SortFunc{quickSort, countingSort}

func BucketSort(data []int) {
	bucket := newBucket(computerBucketNum(data, bucketScope))

	data2bucket(data, bucket)
	fmt.Println("--------------111------------")
	sortForBucket(bucket, sortFuncSlice)

	bucket2data(bucket, data)
	fmt.Println("--------------123------------")
}

func computerBucketNum(data []int, bucketScope int) int {
	min, max := countMaxMin(data)

	return (max - min) / bucketScope
}

func newBucket(bucketNum int) map[int][]int {
	var bucket = make(map[int][]int)
	for i := 0; i < bucketNum; i++ {
		bucket[i] = []int{}
	}

	return bucket
}

func data2bucket(data []int, bucket map[int][]int) {
	for i := 0; i < len(data); i++ {
		// 判断data[i] 存储到 哪个桶里
		index := data[i] / bucketScope
		fmt.Printf("桶号:\t%d, 值:\t%d\n", index, data[i])
		bucket[index] = append(bucket[index], data[i])
	}
}

func sortForBucket(bucket map[int][]int, sortFunc []SortFunc) {
	var index int
	for _, slice := range bucket {
		sortFunc[index](slice)
		index++
		if index == len(sortFunc) {
			index = 0
		}
	}
}

func bucket2data(bucket map[int][]int, data []int) {
	var index int
	for i := 0; i < len(bucket); i++ {
		slice := bucket[i]
		if len(slice) != 0 {
			for j := 0; j < len(slice); j++ {
				fmt.Printf("===%d===桶内数据:\t%d\n", j, slice[j])
				data[index] = slice[j]
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
	fmt.Println("--------------2------------")
	var index int
	for i := min; i < len(temp); i++ {
		for j := temp[i]; j > 0; j-- {
			data[index] = i
			index++
			fmt.Println("--------------2-1-----------")
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
