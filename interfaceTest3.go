package main

import "fmt"

type Duration int

func (d *Duration) show() {
	fmt.Println(d)
	fmt.Printf("value=%v\n", *d)
}

func main() {
	a := Duration(42)
	//fmt.Println("a=", a)
	//Duration(42).show()
	a.show()
	//(&a).show()
}
