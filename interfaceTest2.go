package main

import (
	"fmt"
	"reflect"
)

// 定义一个结构体类型
type Person struct {
	Name string
	Age  int
}

// 定义一个方法
func (p Person) Greet() string {
	return "Hello, " + p.Name
}
func (p Person) SayHi(s string) {
	fmt.Printf("Hi, I'm %s\n", p.Name)
}
func (p *Person) ChangeInfo(name string, age int) {
	p.Name = name
	p.Age = age
	fmt.Printf("Info changed: => Name: %s, Age: %d\n", p.Name, p.Age)
}

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
func main() {
	// 创建一个 Person 实例
	p := Person{Name: "Alice", Age: 30}
	obj := Outer{
		A: 1,
		B: "B",
		Inner: Inner{
			C:    2,
			D:    "D",
			Base: Base{"Your Dad"},
		},
	}
	t2 := reflect.TypeOf(obj)
	//v2 := reflect.ValueOf(obj)
	for i := 0; i < t2.NumMethod(); i++ {
		m := t2.Method(i)
		fmt.Println(m)
	}
	// 使用反射获取类型和值
	t := reflect.TypeOf(p)
	v := reflect.ValueOf(p)

	// 打印类型信息
	fmt.Println("Type:", t)
	fmt.Println("Kind:", t.Kind()) // 应该是 reflect.Struct

	// 遍历结构体的字段
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		fmt.Printf("Field name: %s, Field type: %s, Field value: %v\n", field.Name, field.Type, value)
	}

	// 注意：由于 p 是通过值传递的，因此下面的 Set 操作会失败（panic）
	// 为了让 Set 操作成功，我们需要一个指向 p 的指针
	pv := reflect.ValueOf(&p).Elem()

	// 修改 Age 字段的值（要求字段可设置）
	if pv.FieldByName("Age").CanSet() {
		pv.FieldByName("Age").SetInt(35)
	}

	// 打印修改后的 p
	fmt.Printf("Modified Person: %+v\n", p)

	// 调用 Greet 方法
	method := pv.MethodByName("Greet")
	results := method.Call(nil) // Greet 方法没有参数，所以传入 nil
	fmt.Println("Greet result:", results[0].String())

	// 遍历sturct de 所有method
	t = reflect.TypeOf(&p)
	v = reflect.ValueOf(&p)

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Println(m)
	}
}
