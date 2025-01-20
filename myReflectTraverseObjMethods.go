package main

import (
	"fmt"
	"reflect"
)

type Base struct {
	RootName string
}

func (a Base) Method1() {
	fmt.Println("Root Method1")
}

/* ------------------Inner -------------------*/
type Inner struct {
	Base
	C int
	D string
}

func (a Inner) SayHi() {
	fmt.Println("Inner: Hi ")
}

func (a *Inner) SetInnerC(c int) {
	fmt.Println("SetInnerC: C=", c)
	a.C = c

}

/* ------------------Inner -------------------*/

/* ------------------Outer -------------------*/
type Outer struct {
	A int
	B string
	Inner
}

func (a *Outer) SetOuterA(b int) {
	fmt.Println("SetOuterA: A=", b)
	a.A = b
}
func (a *Outer) SetOuterB(b string) {
	fmt.Println("SetOuterB: B=", b)
	a.B = b
}
func (a Outer) GetOuterA() int {
	fmt.Println("GetOuterA: ", a.A)
	return a.A
}
func (a Outer) GetOuterB() string {
	fmt.Println("GetOuterB: ", a.B)
	return a.B
}

/* ------------------Outer -------------------*/

func TraverseObjMethods(obj interface{}) {
	// 使用反射获取类型和值
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	// 如果传入的是指针，先解引用
	// 需要区分 值接收者方法和指针接受者方法 TBD

	//if t.Kind() == reflect.Ptr {
	//	t = t.Elem()
	//	v = v.Elem()
	//}
	// 打印obj实际类型信息
	fmt.Println("Type:", t)
	fmt.Println("Kind:", t.Kind())
	traversMethodRecursive(t, v, t.Name(), 0)
}

func traversMethodRecursive(type_ reflect.Type, value_ reflect.Value, prefix_ string, level_ uint) {
	// 自定义的基础类型也可能有绑定method TBD
	//if value_.Kind() == reflect.Ptr {
	//	// 解引用，如果是函数指针，直接打印类型和值
	//} else if value_.Kind() == reflect.Struct {
	//	// 遍历methods
	//	for i := 0; i < value_.NumMethod(); i++ {
	//		method := value_.Method(i)
	//		//fmt.Printf("Method[%d]: %s %v\n", i, method.Name, method.Type())
	//		fmt.Println("Method[%d]:", method)
	//
	//	}
	//} else if value_.Kind() == reflect.Func {
	//	// 如果是函数类型， 直接打印类型
	//	//reflect.
	//} else if value_.Kind() == reflect.Map {
	//	// map 的value如果又是函数指针或者函数， TBD
	//} else {
	for i := 0; i < type_.NumMethod(); i++ {
		method := type_.Method(i)
		//fmt.Printf("Method[%d]: %s %v\n", i, method.Name, method.Type())
		fmt.Println("Method[%d]:", method)

	}
	//}

}

func main() {
	obj := Outer{
		A: 1,
		B: "B",
		Inner: Inner{
			C:    2,
			D:    "D",
			Base: Base{"Your Dad"},
		},
	}
	TraverseObjMethods(&obj)

}
