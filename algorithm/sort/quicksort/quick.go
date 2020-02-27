package quicksort

func quickSort(data []int) {

	subQuickSort(data, 0, len(data)-1)
}

func subQuickSort(data []int, l, r int) {
	// 这种情况是为了防止，数据进行递归运算时，中间数据状态可能满足一定的条件，如5，2，1，3，4，9，8，7时
	// 左分区[5,2,1,3，4],右分界线[9,8,7]
	// 数据[[5,2,1,3，4]],经过下面的计算后，分界线是最后一位，然后，此时再递归调用时，就会出现死循环了
	// 因为目前subQuickSort函数递归调用自己的时候，调两次，根据分界线，左分区和右分区
	// 很有可能，分界线刚好时，开头或者结尾，此时再去调用自己的时候，就不会满足具体业务，如left < right，会进入死循环了
	if r < l {
		return
	}

	if r-l == 1 {
		if data[l] > data[r] {
			data[l], data[r] = data[r], data[l]
		}
		return
	}

	if r-l == 2 {
		compareForThree(data, l, r)
		return
	}

	pivotIndex := l

	pivot := data[pivotIndex]

	var left, right = pivotIndex+1, r

	for left < right {
		for ; left < right; right-- {
			if data[right] < pivot {
				break
			}
		}

		for ; left < right; left++ {
			if data[left] > pivot {
				break
			}
		}

		if left < right {
			data[left], data[right] = data[right], data[left]
			left++
			right--
		}

		if left == right && data[pivotIndex] > data[right] {
			data[pivotIndex], data[right] = data[right], data[pivotIndex]
			// 交换后，下标也需要交换
			pivotIndex = right
			break
		}
	}

	// 右边分区
	subQuickSort(data, pivotIndex+1, r)

	// 左边分区
	subQuickSort(data, l, pivotIndex-1)

}

func compareForThree(datas []int, i, j int) {
	if datas[i] > datas[i+1] {
		datas[i], datas[i+1] = datas[i+1], datas[i]
	}

	if datas[i] > datas[j] {
		datas[i], datas[j] = datas[j], datas[i]
	}

	if datas[i+1] > datas[j] {
		datas[i+1], datas[j] = datas[j], datas[i+1]
	}
}
