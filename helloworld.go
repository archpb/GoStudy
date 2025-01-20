package main

import (
	"fmt"
	uuu "goProj1/utils"
	//utils2 "goProj1/utils"
)

//import "utils/tools"

/* sss*/
func main() {

	uuu.ShowMe()
	var a, b int
	a = 11
	b = 22
	var c int
	c, _ = add(&a, b)
	fmt.Printf("add(a,b)=%d", c)
	fmt.Printf("after add, a = %d\n", a)

	total, _ := sum("eee", 1, 23, 45, 65, -87)
	fmt.Println("total:%d", total)

	// array
	//arrStatic := [4]int{1, 2, 3, 4}
	//arrFrag := []int{11, 22, 33, 44}
	b := 1
	if b = 11; b < 10 {
		fmt.Printf("b is available\n b:%d", b)
	} else {
		fmt.Printf("b inn else avaliable %d\n", b)
	}
}

// func f1
func add(a *int, b int) (int, int) {
	*a = 111
	return *a + b + 100, b
}

func sum(name string, para ...int) (int, string) {
	sum := 0
	lenParas := len(para)
	for i := 0; i < lenParas; i++ {
		fmt.Println(para[i])
		sum += para[i]
	}
	return sum, name
}
