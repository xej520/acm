package shell_sort

func ShellSort(data []int) {
	subShellSort(data, len(data) / 2)
}

func subShellSort(data []int, interval int) {
	if interval == 0{
		return
	}

	for startIndex := 0; startIndex < len(data); startIndex++ {
		insertSort(data, startIndex, interval)
	}


	subShellSort(data, interval/2)
}

func insertSort(data []int, startIndex, interval int) {

}





