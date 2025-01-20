package main

import "fmt"

type emptyIf interface{}

type myStruct struct {
	//n int
}

func main() {

	var e []emptyIf = make([]emptyIf, 5)
	my := myStruct{}
	c := make(chan int)
	e[0], e[1], e[2], e[3], e[4] = my, 121, "test", false, c

	// commit-ok check
	for i, element := range e {
		if value, ok := element.(int); ok {
			fmt.Printf("e[%d] type is (int), value=%v\n", i, value)
		} else if value, ok := element.(string); ok {
			fmt.Printf("e[%d] type is (string), value=%v\n", i, value)
		} else if value, ok := element.(chan int); ok {
			fmt.Printf("e[%d] type is (chan int), value=%v\n", i, value)
		} else if value, ok := element.(bool); ok {
			fmt.Printf("e[%d] type is (bool), value=%v\n", i, value)
		} else if value, ok := element.(myStruct); ok {
			fmt.Printf("e[%d] type is (myStruct), value=%v\n", i, value)
		} else {
			fmt.Printf("e[%d] not matched!\n", i)
		}
	}

	fmt.Println("==========================")
	// switch match
	for i, element := range e {
		switch value := element.(type) {
		//case emptyIf: // 这里所有都会匹配到空接口
		//	fmt.Printf("e[%d] type is (emptyIf), value=%v\n", i, value)
		case myStruct:
			fmt.Printf("e[%d] type is (myStruct), value=%v\n", i, value)
		case chan int:
			fmt.Printf("e[%d] type is (chan int), value=%v\n", i, value)
		case bool:
			fmt.Printf("e[%d] type is (bool), value=%v\n", i, value)
		case int:
			fmt.Printf("e[%d] type is (int), value=%v\n", i, value)
		case string:
			fmt.Printf("e[%d] type is (string), value=%v\n", i, value)
		default:
			fmt.Printf("e[%d] not matched!\n", i)
		}
	}

}
