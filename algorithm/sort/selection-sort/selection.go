package selection_sort

func SelectionSort(data []int) {
	subSelectionSort(data, 0)
}

func subSelectionSort(data []int, startIndex int) {

	if startIndex == len(data)-1 {
		return
	}

	tempStartIndex := evaluationValue(data, startIndex)

	subSelectionSort(data, tempStartIndex)

}

// 每一轮，只交换一次数据，而冒泡排序，交换次数多
func evaluationValue(data []int, startIndex int) int {

	// 临时最小值下标
	tempMinIndex := startIndex

	// 找到data数据从startIndex后的最小值下标
	for j := startIndex + 1; j < len(data); j++ {
		if data[tempMinIndex] > data[j] {
			tempMinIndex = j
		}
	}

	// 将最小值跟data数据中startIndex的元素，进行交换
	data[startIndex], data[tempMinIndex] = data[tempMinIndex], data[startIndex]

	return startIndex + 1
}
