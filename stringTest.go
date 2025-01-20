package main

import "fmt"

func main() {
	str := "hello你好啊"

	for i, value := range str { //str = "hello你好啊"
		fmt.Printf("str[%d]=%c\n", i, value)
	}
	fmt.Println("=============================")

	for j := 0; j < len(str); j++ {
		fmt.Printf("str[%d]=%c\n", j, str[j])
	}

	fmt.Printf("str[%d]=%c\n", 8, str[8])
	fmt.Println("=============================")

	ss := "helloabckdeeef"
	sPointer := &ss
	//var sPointerIndex
	//用%s输出string类型指针， *sPointer不能直接sPointer
	fmt.Printf("s_pointer type:%T, p=%p, %s\n", sPointer, sPointer, *sPointer)
	sliceByte := ss[3:]

	fmt.Printf("sliceByte type:%T", sliceByte)

	//var byteArray byte[]
	fmt.Println("=============================")

	const (
		_      = iota   //0, 忽略不用
		Monday = iota   // 1
		Tue    = iota   // 2
		Wed             // 3
		Thur            // 4
		Fri    = 'A'    // 'A'
		Sat    = iota   // 6
		Sun             // 7
		Md     = "asdf" // "asdf"
	)
	fmt.Println(Monday, Tue, Wed, Thur, Fri, Sat, Md)
	const (
		a = iota // 0
		b = iota // 1
	)
	fmt.Println(a, b)
	const (
		_  = iota
		KB = 1 << (iota * 10) // 1024
		MB
		GB
	)

	fmt.Println(KB, MB, GB)

}
