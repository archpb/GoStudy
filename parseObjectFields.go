package main

import (
	"fmt"
	"reflect"
)

// ParseFields parses the fields of any input object, including nested and anonymous struct fields.
func ParseFields(input interface{}) {
	if input == nil {
		fmt.Println("Input is nil.")
		return
	}

	inputValue := reflect.ValueOf(input)
	inputType := reflect.TypeOf(input)

	// If the input is a pointer, dereference it.
	if inputValue.Kind() == reflect.Ptr {
		inputValue = inputValue.Elem()
		inputType = inputType.Elem()
	}

	fmt.Printf("Type: %s\n", inputType.Name())

	// Call the recursive parsing function.
	parseFieldsRecursive(inputValue, inputType, "")
}

// parseFieldsRecursive handles recursion for nested and anonymous fields.
func parseFieldsRecursive(value reflect.Value, typ reflect.Type, prefix string) {
	if value.Kind() == reflect.Struct {
		for i := 0; i < value.NumField(); i++ {
			field := value.Field(i)
			fieldType := typ.Field(i)

			fieldName := fieldType.Name
			if prefix != "" {
				fieldName = prefix + "." + fieldName
			}

			if fieldType.Anonymous {
				// Handle anonymous fields (embedded structs).
				parseFieldsRecursive(field, field.Type(), fieldName)
			} else {
				fmt.Printf("Field: %s, Type: %s, Value: %v\n", fieldName, field.Type(), field.Interface())
			}
		}
	} else {
		// If the value is not a struct, print its type and value.
		fmt.Printf("Value: %v, Type: %s\n", value.Interface(), value.Type())
	}
}

func main() {
	type NamedStruct struct {
		Skin int
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
		HiNamedStruct NamedStruct
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
		HiNamedStruct: NamedStruct{99},
	}

	ParseFields(obj)
	ParseFields(1)
}
