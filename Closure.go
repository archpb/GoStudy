package main

import "fmt"

func main() {
	funcs := []func(){}
	for i := 0; i < 5; i++ {
		//i := i
		funcs = append(funcs, func() {
			fmt.Println(i)
		})
	}

	// 打印的结果是：0 1 2 3 4
	for _, f := range funcs {
		f()
	}
}
