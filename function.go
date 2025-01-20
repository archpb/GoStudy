package main

import "fmt"

func main() {

	var arrayA [3]int                                               // empty
	var arrayB [4]int = [4]int{1, 2, 3, 4}                          // full init
	arrayC := [10]byte{1, 2, 3, 4, 5, 6}                            // partially init
	arrayD := [...]string{"1st", "2nd", "3rd", "4th", "5th", "6th"} // init with length ...

	fmt.Println(arrayA, arrayB, arrayC, arrayD)
	f1("aaa", 1)
	fmt.Printf("func type is %T\n", f1)
	fmt.Printf("var f1 is %p\n", f1)

	var funcVar func(string, int) bool = f1
	funcVar("use func var to print", 111)

	//var func2 := f1
	//func2("func2 := f1", 155)
	func2 := func(a, b int) int {
		fmt.Printf("in 匿名函数，a=%d, b:=%d return:%d\n", a, b, a*b)
		func3 := func() {
			fmt.Println("this a 匿名函数 inside func2")
		}
		func3()
		return a * b
	}
	func2(2, 4)

	oper(10, 20, add)
	oper(10, 20, sub)
	oper(10, 20, func(a, b int) int {
		c := a * b
		fmt.Println("in anonymous func used as func para %d", c)
		return c
	})

}

func f1(s string, a int) bool {
	fmt.Println(s, a)
	return true
}

func add(a, b int) int {
	return a + b
}
func sub(a, b int) int {
	return a - b
}
func oper(a, b int, f func(int, int) int) int {
	fmt.Println(f(a, b))
	return f(a, b)
}
