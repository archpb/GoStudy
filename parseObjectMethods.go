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

// TraverseMethods 解析对象的所有方法，包括匿名结构体的方法
func TraverseMethods(obj interface{}) {
	// 获取对象的类型和值
	t := reflect.TypeOf(obj)

	// 如果是指针类型，解引用
	//if t.Kind() == reflect.Ptr {
	//	t = t.Elem()
	//}

	// 遍历类型及其所有方法
	traverseTypeMethods(t, "")
}

func traverseTypeMethods(t reflect.Type, prefix string) {
	if t.Kind() != reflect.Struct {
		return
	}

	// 遍历当前结构体的字段
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		// 如果是匿名字段，递归解析
		if field.Anonymous {
			fieldType := field.Type
			if fieldType.Kind() == reflect.Ptr {
				fieldType = fieldType.Elem()
			}
			traverseTypeMethods(fieldType, prefix+"."+field.Name)
		}
	}

	// 打印当前类型的方法
	printMethods(t, prefix)
}

func printMethods(t reflect.Type, prefix string) {
	methodCount := t.NumMethod()
	for i := 0; i < methodCount; i++ {
		method := t.Method(i)
		fmt.Printf("Method: %s%s.%s, Type: %s\n", prefix, t.Name(), method.Name, method.Type)
	}
}

func main() {

	// 初始化对象
	obj := Outer{}

	// 调用 TraverseMethods
	TraverseMethods(&obj)
}
