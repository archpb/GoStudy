package main

import (
	"fmt"
)

// test user self-defined pointer type as a normal type
type IIP []byte

func (i IIP) changeI() {
	i = []byte("newone") // 这里是将临时切片指向了一个新的内存地址，所以原切片x实际上没有被修改,
	i[0] = 'S'           // z这样才是真正的修改了i和x指向共享内存的值
	i[1] = 'B'
	fmt.Println("in changI(): ", string(i))
}

func main() {
	z := "hello world"
	var x IIP = []byte(z)
	fmt.Println(string(x))
	x.changeI()
	fmt.Println(string(x))

	var sliceA []int = make([]int, 8, 10) // make 会分配一个长度为8，容量为10的切片
	var tmp []int = sliceA
	fmt.Println(tmp)
	var sliceAA []string                     // empty slice
	arrayA := [8]int{1, 2, 3, 4, 5, 6, 7, 8} // init from a existed array
	sliceBB := arrayA[2:5]

	fmt.Println(sliceA, sliceAA, sliceBB)

	sliceA[0] = 1
	sliceA[1] = 2
	sliceA[2] = 3
	sliceA[3] = 4
	for i, _ := range sliceA {
		sliceA[i] = i + 10
	}

	fmt.Println(sliceA)
	fmt.Printf("len: %d, cap: %d\n", len(sliceA), cap(sliceA))
	fmt.Printf("pointer sliceA = %p\n", sliceA)

	sliceA2 := append(sliceA, 100, 110, 34)
	fmt.Println(sliceA)
	fmt.Println(sliceA2)
	fmt.Printf("len: %d, cap: %d\n", len(sliceA), cap(sliceA))
	fmt.Printf("pointer sliceA = %p\n", sliceA)
	fmt.Printf("pointer sliceA[1] = %p\n", &sliceA[1])

	arrA := []int{1, 2, 3, 4, 5, 6, 7, 8}
	sliceB := arrA[2:5]
	fmt.Printf("len: %d, cap: %d\n", len(sliceB), cap(sliceB))

	sliceStr := make([]string, 3, 6)
	sliceStr[0] = "aaa"
	sliceStr[1] = "bbb"
	sliceStr[2] = "cccc"
	//sliceStr[3] = "dddd"

	fmt.Println(sliceStr)
	fmt.Printf("len: %d, cap: %d\n", len(sliceStr), cap(sliceStr))

	sliceStr = append(sliceStr, "ddd", "哈喽试试", "eeeeee")
	fmt.Println(sliceStr)

	//traverse slice
	for i, _ := range sliceStr {
		fmt.Printf("sliceStr[%d]=%s, len=%d\n", i, sliceStr[i], len(sliceStr[i]))
	}
	for i := 0; i < len(sliceStr)-1; i++ {
		sliceStr[i] = sliceStr[i] + sliceStr[i+1]
	}
	fmt.Println(sliceStr)

	// slice copy

}
