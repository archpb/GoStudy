package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// 创建一个Scanner实例，用于从标准输入读取数据
	scanner := bufio.NewScanner(os.Stdin)

	// 设置Scanner的分隔符为按单词分割
	scanner.Split(bufio.ScanWords)

	fmt.Println("Enter some words (type 'exit' to quit):")
	for scanner.Scan() {
		word := scanner.Text()
		// 将输入的单词转换为小写，以便与"exit"进行比较
		if strings.ToLower(word) == "exit" {
			break
		}
		fmt.Println("You entered:", word)
	}

	// 检查扫描过程中是否发生错误
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}
