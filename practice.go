package main

import "fmt"

func main() {
	a := 4
	fmt.Printf("Hello World, fac(%d)=%v\n", a, fac(a))

	fmt.Printf("maxs(xxx)=%d\n", maxs(1, 3, 45, 452, 33, 62, 34, 5, 1, -2320, 234))

	var test []int = []int{2, 3, 4, 34, 6, 7, 8, 9, 10, 3, 3, 34, 6, 23, 34, 34, 2, 2, 2, 3, 3, 6, 34}
	c, d := findMaxMinIndexOfArray(test)
	fmt.Printf("The max index: %v\nThe min index: %v\n", c, d)

	var ff []float64 = []float64{2.23, 3.5423, 12.5, 0.656, 13, 48.4}
	fmt.Printf("average of float slice is: %v\n", averageFloat(ff))
}

func fac(n int) int {
	if n == 1 {
		return 1
	} else {
		return n * fac(n-1)
	}
}
func maxs(args ...int) int {
	var r int = 0
	for _, v := range args {
		if r < v {
			r = v
		}
	}
	return r
}
func findMaxMinIndexOfArray(arr []int) (maxIdxSlice, minIdxSlice []int) {
	maxIdx, minIdx := 0, 0
	maxIdxSlice = make([]int, 0)
	minIdxSlice = make([]int, 0)

	for i, _ := range arr {
		if i == len(arr)-1 {
			break
		}
		if arr[i+1] >= arr[maxIdx] {
			maxIdx = i + 1
		}
		if arr[i+1] <= arr[minIdx] {
			minIdx = i + 1
		}
	}
	fmt.Printf("maxIdx=%d(v=%d), minIdx=%d(v=%d)\n", maxIdx, arr[maxIdx], minIdx, arr[minIdx])

	for i, v := range arr {
		if v == arr[maxIdx] {
			maxIdxSlice = append(maxIdxSlice, i)
			continue
		}
		if v == arr[minIdx] {
			minIdxSlice = append(minIdxSlice, i)
		}
	}
	return maxIdxSlice, minIdxSlice
}
func averageFloat(f []float64) float64 {
	var sum float64
	for _, v := range f {
		sum += v
	}
	return sum / float64(len(f))
}
