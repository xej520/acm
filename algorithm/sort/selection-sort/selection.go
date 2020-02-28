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

func evaluationValue(data []int, startIndex int) int {
	// 临时最小值
	tempMin := data[startIndex]
	// 临时最小值下标
	tempMinIndex := startIndex

	for j := startIndex + 1; j < len(data); j++ {
		if tempMin > data[j] {
			tempMinIndex, tempMin = j, data[j]
		}
	}

	data[startIndex], data[tempMinIndex] = data[tempMinIndex], data[startIndex]

	return startIndex + 1
}
