package fangshi2

import (
	"math"
	"fmt"
	"strconv"
)

func RadixSort(data []int) {

	ml := maxLength(data)

	data2Bucket(data, 1, ml)
}

func newBucket() map[int][]int {
	var b = make(map[int][]int)

	for i := 0; i < 10; i++ {
		b[i] = []int{}
	}

	return b
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

func data2Bucket(data []int, i, ml int) {
	if i > ml {
		return
	}

	sortForBucket(data, i)

	data2Bucket(data, i+1, ml)
}

func sortForBucket(data []int, i int) {
	b := newBucket()

	for _, v := range data {
		vs := strconv.Itoa(v)
		l := len(vs)

		if i <= l {
			vd := fmt.Sprintf("%c", vs[l-i])
			k, e := strconv.Atoi(vd)
			if e == nil {
				b[k] = append(b[k], v)
			}

		} else {
			b[0] = append(b[0], v)
		}
	}

	var index int
	for j:=0; j<10; j++{
		v := b[j]
		if len(v) != 0 {
			for i := 0; i < len(v); i++ {
				data[index] = v[i]
				index++
			}
		}
	}
}
