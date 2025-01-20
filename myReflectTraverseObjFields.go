package main

import (
	"fmt"
	"reflect"
)

func TraverseObjFields(obj interface{}) {
	// 使用反射获取类型和值
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	// 如果传入的是指针，先解引用
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}
	// 打印obj实际类型信息
	fmt.Println("Type:", t)
	fmt.Println("Kind:", t.Kind())
	traversRecursive(t, v, t.Name(), 0)
}

func traversRecursive(type_ reflect.Type, value_ reflect.Value, prefix_ string, level_ uint) {

	if value_.Kind() == reflect.Struct {
		for i := 0; i < value_.NumField(); i++ {
			fieldType := type_.Field(i) // return structField structure, see above comments of StructField
			field := value_.Field(i)    // return Value

			fieldName := fieldType.Name
			if fieldType.Type.Kind() == reflect.Ptr { // 如果是指针，在field name前面加个‘*’
				fieldName = fmt.Sprintf("%s[%d].*%s", prefix_, i, fieldName)
			} else {
				fieldName = fmt.Sprintf("%s[%d].%s", prefix_, i, fieldName)
			}

			if fieldType.Type.Kind() == reflect.Ptr && fieldType.Type.Elem().Kind() == reflect.Struct {
				if field.IsNil() { //空指针直接打印 nil
					fmt.Printf("%s[%d].*%s: Type=%v, Value=%v, (Offset=%v)\n",
						prefix_, i, fieldType.Name, fieldType.Type, field, fieldType.Offset)
				} else {
					// 解引用指针
					traversRecursive(fieldType.Type.Elem(), field.Elem(), fieldName, level_+1)
				}

			} else if fieldType.Anonymous || field.Kind() == reflect.Struct {
				// Anonymous field,
				traversRecursive(fieldType.Type, field, fieldName, level_+1)

			} else {
				// normal field
				fmt.Printf("%s[%d].%s: Type=%v, Value=%v, (Offset=%v)\n",
					prefix_, i, fieldType.Name, fieldType.Type, field, fieldType.Offset)
			}
		}
	}

}

type FuncPointer func(string) string

func main() {

	type NamedStruct struct {
		Skin  int
		Heart bool
	}
	type Base struct {
		Root         string
		IntPointer   *int
		BaseMap      map[int]string
		ChanIntSlice []chan int
		ArrayByte    []byte
	}
	type Inner struct {
		Base
		C int
		D string
	}

	type Outer struct {
		A int
		B string
		Inner
		Anonymous struct {
			E float64
			F bool
		}
		HiNamedStruct *NamedStruct
		IfVar         interface{}
		FuncP         *FuncPointer
	}

	obj := Outer{
		A: 1,
		B: "test",
		Inner: Inner{
			C: 2,
			D: "inner",
			Base: Base{
				Root:         "root",
				IntPointer:   new(int),
				BaseMap:      map[int]string{1: "one", 2: "two"},
				ChanIntSlice: []chan int{make(chan int), make(chan int), make(chan int)},
				ArrayByte:    []byte{'a', 'b', 'c', 'd', 'e', 'f'},
			},
		},
		Anonymous: struct {
			E float64
			F bool
		}{
			E: 3.14,
			F: true,
		},
		HiNamedStruct: nil,
		//HiNamedStruct: &NamedStruct{Skin: 99, Heart: true},
		IfVar: nil,
		FuncP: new(FuncPointer),
	}

	TraverseObjFields(obj)
	fmt.Println("=======================")

	type Inner2 struct {
		C int
		D string
	}
	type Outer2 struct {
		A int
		B string
		Inner2
		Anonymous struct {
			E float64
			F bool
		}
	}
	obj2 := Outer2{
		A: 1,
		B: "test",
		Inner2: Inner2{
			C: 2,
			D: "inner",
		},
		Anonymous: struct {
			E float64
			F bool
		}{
			E: 3.14,
			F: true,
		},
	}
	TraverseObjFields(obj2)

}
