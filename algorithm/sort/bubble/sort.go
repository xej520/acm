package main

import (
	"fmt"
	"flag"
	"os"
	"runtime/pprof"
	"time"
	"runtime"
)

func sortByMax(arr []int) []int {

	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] < arr[j+1] {
				tmp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			}
		}
	}

	return arr
}

func sortByMin(arr []int) []int {

	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				tmp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			}
		}
	}

	return arr
}

func sortByMin2(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := len(arr) - 1; j-1-i >= 0; j-- {
			if arr[j] > arr[j-1] {
				tmp := arr[j]
				arr[j] = arr[j-1]
				arr[j-1] = tmp
			}
		}
	}

	return arr
}

// 利用Go的多重赋值特性，可以很简单的完成两个变量的交换
func sortByMin3(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := len(arr) - 1; j > i; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}

	return arr
}

func print(arr []int) {

	for _, v := range arr {
		fmt.Printf("%d ", v)
	}

	fmt.Println()
}

var cpuProfile = flag.String("cpuprofile", "", "write cpu profile to file")
var memProfile = flag.String("memprofile", "", "write cpu profile to file")
var memProfileRate = flag.Int("memprofilerate", 1, "write cpu profile to file")

func startCPUProfile() {
	if *cpuProfile != "" {
		f, err := os.Create(*cpuProfile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Can not create cpu profile output file: %s",
				err)
			return
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			fmt.Fprintf(os.Stderr, "Can not start cpu profile: %s", err)
			f.Close()
			return
		}
	}
}

func stopCPUProfile() {
	if *cpuProfile != "" {
		pprof.StopCPUProfile() // 把记录的概要信息写到已指定的文件
	}
}

func startMemProfile() {
	if *memProfile != "" && *memProfileRate > 0 {
		runtime.MemProfileRate = *memProfileRate
	}
}

func stopMemProfile() {
	if *memProfile != "" {
		f, err := os.Create(*memProfile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Can not create mem profile output file: %s", err)
			return
		}
		if err = pprof.WriteHeapProfile(f); err != nil {
			fmt.Fprintf(os.Stderr, "Can not write %s: %s", *memProfile, err)
		}
		f.Close()
	}
}


func main() {
	flag.Parse()
	startMemProfile()
	defer stopMemProfile()

	arr := []int{
		6, 8, 5, 4, 3, 2, 1, 0,
	}

	for i := 0; i < 10000; i++ {
		time.Sleep(time.Millisecond)
		arr = sortByMin(arr)
	}

	print(arr)

}
