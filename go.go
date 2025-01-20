package main

import "fmt"

func main() {
	var a, b int = 10, 1
	var p *int = &a
	a = *p + b
	println(a)

	var str1, str2 string
	fmt.Scanf("%s %s", &str1, &str2)
	println(str1, str2)

	var aaa *int = new(int) * 8
	*aaa = 100
	println(aaa)

	var bbb []int = new(int)

}
