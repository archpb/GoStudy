package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Num   int
	Name  string
	Sex   string
	Score int
}

func main() {
	var mapA map[interface{}]interface{}
	mapA = make(map[interface{}]interface{}, 3)
	mapA["a"] = 1
	mapA[2] = 'B'
	mapA[false] = "bool"
	mapA["array"] = []int{1, 2, 3}

	fmt.Println(mapA)
	fmt.Println("=================")
	type person struct {
		name string
		age  int
	}
	aa := 100
	var sliceA []interface{} = make([]interface{}, 8)
	sliceA = []interface{}{111, "hello", true, 1.2121, 'A',
		struct{ a int }{a: 1},
		person{
			name: "adf",
			age:  13,
		},
		&aa,
	}

	fmt.Println(sliceA)
	for i, v := range sliceA {
		if value, ok := v.(int); ok {
			fmt.Printf("sliceA[%d] type:(int), value=%v\n", i, value)
		} else if value, ok := v.(string); ok {
			fmt.Printf("sliceA[%d] type:(string), value=%v\n", i, value)
		} else if value, ok := v.(float64); ok {
			fmt.Printf("sliceA[%d] type:(float64), value=%v\n", i, value)
		} else if value, ok := v.(map[string]interface{}); ok {
			fmt.Printf("sliceA[%d] type:(map[string]interface{}), value=%v\n", i, value)
		} else if value, ok := v.(person); ok {
			fmt.Printf("sliceA[%d] type:(person), value=%v\n", i, value)
		} else if value, ok := v.(bool); ok {
			fmt.Printf("sliceA[%d] type:(bool), value=%v\n", i, value)
		} else if value, ok := v.(*int); ok {
			fmt.Printf("sliceA[%d] type:(*int), value=%v\n", i, *value)
		} else {
			fmt.Printf("default: sliceA[%d] type:(%s), value=%v\n", i, reflect.TypeOf(v).Name(), reflect.ValueOf(v))
		}
	}

	stu1 := Student{Num: 01, Name: "Xiaoming", Sex: "male", Score: 100}
	fmt.Println(stu1)
	fmt.Println("========change student==========")
	ptrStu := &stu1
	t := reflect.TypeOf(ptrStu).Elem()
	v := reflect.ValueOf(ptrStu).Elem()
	for i := 0; i < t.NumField(); i++ {
		field := v.Field(i)
		typ := t.Field(i)
		fmt.Println(field.Interface(), typ.Name, typ.Type.Name())
		if field.CanSet() && typ.Type.Kind() == reflect.Int {
			if typ.Name == "Num" {
				field.SetInt(11)
			} else if typ.Name == "Score" {
				field.SetInt(99)
			}
		} else if field.CanSet() && typ.Type.Kind() == reflect.String && typ.Name == "Name" {
			field.SetString("HanMeimei")
		}

	}
	fmt.Println("after changne:", stu1)
}
