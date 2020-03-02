package counting_sort

// 计数排序
func CountingSort(data []int) {
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
		for j := temp[i]; j > 0 ;j--{
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
