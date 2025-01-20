package main

import (
	"fmt"
	"sort"
)

// prereqs记录了每个课程的前置课程
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}
func topoSort(m map[string][]string) []string {
	var order []string                // 一个字符串切片，用于存储拓扑排序的结果。
	seen := make(map[string]bool)     //用于记录哪些课程已经被访问过，防止重复访问和陷入无限循环。
	var visitAll func(items []string) // 递归函数，用于访问所有课程及其前置课程。当匿名函数需要被递归调用时，
	// 我们必须首先声明一个变量（在上面的例子中，我们首先声明了 visitAll），
	// 再将匿名函数赋值给这个变量。如果不分成两部，函数字面量无法与visitAll绑定，我们 也无法递归调用该匿名函数
	visitAll = func(items []string) { // 这里只是定义递归调用函数， 真正的调用在后面visitAll(keys)
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item) // 所有前置课程都学习后， 会被append到排序切片中， 叶子节点将会首先append进去
			}
		}
	}
	var keys []string
	for key := range m {
		keys = append(keys, key) // 收集所有课程名称
	}
	sort.Strings(keys) //
	visitAll(keys)     // DFS深度优先排序
	return order
}
