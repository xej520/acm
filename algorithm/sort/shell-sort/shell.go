package shell_sort

func ShellSort(data []int) {
	subShellSort(data, len(data)/2)
}

func subShellSort(data []int, interval int) {
	// 当间隔等于0时，说明结束了，退出递归
	if interval == 0 {
		return
	}

	// 在间隔一定的时候，对每个分组，都进行插入排序
	for startIndex := 0; startIndex < len(data)-interval; startIndex++ {
		insertSort(data, startIndex, interval)
	}

	// 更新间隔，进入下一次递归
	subShellSort(data, interval/2)
}

func insertSort(data []int, startIndex, interval int) {

	nextInsertIndex := startIndex + interval
	temp := data[nextInsertIndex]

	index := startIndex

	for index >= 0 && temp < data[index] {

		data[index+interval] = data[index]
		index = index - interval

	}

	data[index+interval] = temp

}
