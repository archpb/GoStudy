package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

type tree struct {
	value       int
	left, right *tree
}

// Sort sorts values in place.
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = addNode(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func addNode(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = addNode(t.left, value)
	} else {
		t.right = addNode(t.right, value)
	}
	return t
}

/* 习题实现一： 将tree解析打印*/
func (t *tree) String() {
	if t == nil {
		fmt.Println("<nil>")
	}
	if t.left != nil {
		t.left.String()
	}
	fmt.Printf("%3d", t.value)
	if t.right != nil {
		t.right.String()
	}

}

/* 习题实现二： 将tree解析成字符串序列*/
func (t *tree) String2() string {
	buf := new(bytes.Buffer)

	//buf.WriteString("[")
	if t == nil {
		//buf.WriteString("<nil>]")
		return "<nil>"
	}
	if t.left != nil {
		buf.WriteString(t.left.String2())
	}
	buf.WriteString(strconv.Itoa(t.value))
	buf.WriteString(" ")
	if t.right != nil {
		buf.WriteString(t.right.String2())

	}
	return buf.String()

}

func (t *tree) Len() int {
	var count int
	if t == nil {
		return 0
	}
	count++ // count current node
	if t.left != nil {
		count += t.left.Len()
	}
	if t.right != nil {
		count += t.right.Len()
	}
	return count
}

func main() {
	data := make([]int, 20)
	var t *tree
	for i := range data {
		data[i] = rand.Int() % 20
		t = addNode(t, data[i])
	}
	fmt.Println("orgin data:", data)
	fmt.Println("len:", t.Len())
	t.String()
	fmt.Println("\n String2:", t.String2())
	strings.NewReader("asdf")

}

//!-
