package main

import "fmt"

type ss struct {
	name string
	age  int
}

func (s *ss) show() string {
	return s.name
}

func main() {
	var a ss = ss{"aaabb", 11}
	fmt.Println(a.show())
	fmt.Println((&a).show())
	b := ss{"aaabb", 22}.show()
}
