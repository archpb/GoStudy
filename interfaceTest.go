package main

import "fmt"

// 定义一个接口
type Speaker interface {
	Speak() // 定义一个方法
	//Speak(int) string	// 报错，接口中不能重名
}

type BigSpeaker interface {
	Speaker()
	Speak() // 可以定义和Speaker中同名的方法
	//Speak(int)	// 同Speaker中，报错
}

type Talker interface {
	Speak()
}

// 类型实现接口：值接收者
type Person struct {
	Name string
}

// 值接收者方法
func (p Person) Speak() {
	fmt.Println("Hi, I'm", p.Name)
}

// 类型实现接口：指针接收者
type Animal struct {
	Kind string
}

// 指针接收者方法
func (a *Animal) Speak() {
	fmt.Println("I am a", a.Kind)
}

func main() {
	// 值接收者实现的接口
	var sp Speaker
	p := Person{Name: "Alice"}
	sp = p     // 值赋值，OK
	sp.Speak() // 输出：Hi, I'm Alice
	sp = &p    // 指针赋值，OK
	sp.Speak() // 输出：Hi, I'm Alice
	//var q *Person
	//q = &p
	//q.Speak()	// ok 输出：Hi, I'm Alice
	//(*q).Speak()	// ok 输出：Hi, I'm Alice
	//p.Speak()	// ok 输出：Hi, I'm Alice
	//(&p).Speak()	// ok 输出：Hi, I'm Alice

	// 指针接收者实现的接口
	a := Animal{Kind: "Dog"}
	sp = &a    // 指针赋值，OK
	sp.Speak() // 输出：I am a Dog
	//var b *Animal
	//b = &a
	//a.Speak()    // ok, 输出：I am a Dog
	//(&a).Speak() // ok 输出：I am a Dog
	//b.Speak()    // ok 输出：I am a Dog
	//(*b).Speak() // ok 输出：I am a Dog

	/* 注意 */
	//sp = a     // 值赋值，报错：Animal没有实现Speaker接口

	var tp Talker
	tp = &a
	fmt.Println("Here is a talker:")
	tp.Speak()
}
