package main

import (
	"fmt"
	"math"
)

func main() {
	//var a, b, c float64 = 40, 50, 30
	//
	//defer func() {
	//	if err := recover(); err != nil {
	//		fmt.Printf("error info:%v\n", err)
	//	}
	//}()
	//area := areaTriangle(a, b, c)
	//fmt.Printf("Area = %.2f\n", area)
	//
	for i := 0; i < 5; i++ {
		defer fmt.Printf("defer in for[%d], i=%d\n", i, i)
	}

	j := 10
	defer fmt.Println("defer1(now j=10): j=", j)
	j = 20
	defer fmt.Println("defer2(now j=20): j=", j)

	jPointer := &j
	defer fmt.Printf("defer3(now *jPointer=j=%d): j=%d\n", *jPointer, j)
	*jPointer = 30
	defer fmt.Printf("defer4(now *jPointer=j=%d): j=%d\n", *jPointer, j)

	jj := 10
	defer fmt.Printf("Defer 5: jj=%d\n", jj) // 显式传值
	jj = 20
	defer func() { fmt.Printf("Defer 6: jj=%d\n", j) }() // 捕获闭包，引用 j
	jjPointer := &jj
	*jjPointer = 30
	defer fmt.Printf("Defer 7: jjPointer=%d, jj=%d\n", *jjPointer, jj) // 引用指针

	fmt.Println("main ends. j=", j)
}

//func main() {
//	var a, b, c float64 = 40, 50, 30
//	var area float64
//	defer func() {
//		if r := recover(); r != nil {
//			fmt.Printf("error info:%v\n", r) // 这里 r 是 interface{} 类型，可以安全地打印
//		} else {
//			fmt.Printf("Area = %.2f\n", area)
//		}
//	}()
//	area = areaTriangle(a, b, c) // 注意：如果 panic 发生，这行代码将不会被执行
//}

func areaTriangle(a, b, c float64) float64 {
	if a <= 0 || b <= 0 || c <= 0 ||
		math.Max(math.Max(a, b), c) >= ((a+b+c)/2) {
		//a+b <= c || a+c <= b || b+c <= a {
		panic("This triangle is not valid")
	}
	fmt.Printf("max side = %v\n, s=%v\n", math.Max(math.Max(a, b), c), ((a + b + c) / 2))

	fmt.Printf("Start calculating the area with sides: %.2f %.2f %.2f\n", a, b, c)
	s := (a + b + c) / 2
	return math.Sqrt(s * (s - a) * (s - b) * (s - c))
}
