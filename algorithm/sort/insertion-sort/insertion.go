package insertion_sort

func InsertionSort(data []int) {
	if len(data) == 1 {
		return
	}

	subInsertionSort(data, 1)
}

// insertIndex，插入元素的下标
func subInsertionSort(data []int, insertIndex int) {
	// 结束递归过程，说明所有的元素都已经排序好了
	if insertIndex > len(data)-1 {
		return
	}

	// 将元素data[insertIndex]已经插入到inserIndex前面了
	nexInsertIndex := frontSort(data, insertIndex)

	// 对剩下的元素，继续插入
	subInsertionSort(data, nexInsertIndex)
}

// 将data[insertIndex]元素插入到insertIndex前面，并排序好
// 返还insertIndex+1 下标，也就说，该insertIndex+1这个元素，插入到前面，进行排序了
func frontSort(data []int, insertIndex int) int {
	temp := data[insertIndex]
	i := insertIndex - 1

	for i >= 0 {
		if data[i] <= temp {
			break
		}

		if data[i] > temp {
			data[i+1] = data[i]
			i--
		}
	}

	data[i+1] = temp

	return insertIndex + 1
}
