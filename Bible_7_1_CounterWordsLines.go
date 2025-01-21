/* 计算一个文件中，words， lines的数量*/
package main

import (
	"bufio"
	"bytes"
	"fmt"
	//"io"
	"os"
)

type WordsLinesCounter struct {
	CntWords int
	CntLines int
}

func (c *WordsLinesCounter) Read(r []byte) (n int, err error) {
	c.CntWords = 0
	buf := bytes.NewBuffer(r)
	scanner := bufio.NewScanner(buf)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		c.CntWords++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	// reset buf pointer
	c.CntLines = 0
	//buf.Reset()	// wrong: this will empty the buffer
	buf = bytes.NewBuffer(r)
	scanner = bufio.NewScanner(buf)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		c.CntLines++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	fmt.Printf("Words: %d, Lines: %d\n", c.CntWords, c.CntLines)
	return c.CntLines*10000 + c.CntWords, nil
}

func main() {
	// show current path
	curPath, err := os.Getwd()
	if err != nil {
		fmt.Println("Failed to open Current path:")
	}
	fmt.Printf("Current Path: %v\n", curPath)

	// list files in current path
	curDirFiles, err2 := os.ReadDir(curPath)
	if err2 != nil { // get all the entries in current dir
		fmt.Println("Failed to Get Current directory")
	}
	for _, itemFile := range curDirFiles {
		fmt.Println(itemFile)
	}

	// open target file
	/* 当您以0644作为perm参数调用os.OpenFile时，如果指定的文件不存在，Go将创建一个新文件，并设置其权限为所有者可以读写，所属组和其他用户可以读。
	如果文件已经存在，并且您没有使用os.O_TRUNC标志（该标志会清空文件内容），则文件的现有内容将保持不变，并且其权限不会受到perm参数的影响。
	*/
	scanFile, err3 := os.OpenFile("README.md", os.O_RDONLY, 0644)
	if err3 != nil {
		fmt.Println("Failed to Open README.md")
	}
	defer scanFile.Close()
	//var n, l int

	var res WordsLinesCounter
	var buffer []byte = make([]byte, 200)
	if _, err := scanFile.Read(buffer); err != nil {
		fmt.Println(err)
		return
	}
	_, err = res.Read(buffer)
	if err != nil {
		fmt.Println(err)
	}

	//
	//
	//// scan the file context
	//fmt.Println("============start working============")
	//scanner := bufio.NewScanner(scanFile) // os.file 实现了io.Reader接口，直接使用生成scanner
	//scanner.Split(bufio.ScanWords)
	//for scanner.Scan() {
	//	n++ // words计数器+1
	//}
	//fmt.Println("Workds:", n) // 注意此时文件指针已经指向了文件末尾，重新扫描需要重置文件指针
	//
	//// reset the file pointer to the beginning
	//// Seek设置下一次读/写的位置。offset为相对偏移量，而whence决定相对位置：0为相对文件开头，1为相对当前位置，2为相对文件结尾。它返回新的偏移量（相对开头）和可能的错误。
	//if _, err := scanFile.Seek(0, io.SeekStart); err != nil {
	//	fmt.Println("Failed to reset file pointer:", err)
	//	return
	//}
	//
	//scanner = bufio.NewScanner(scanFile) // os.file 实现了io.Reader接口，直接使用生成scanner
	//scanner.Split(bufio.ScanLines)
	//for scanner.Scan() {
	//	l++ // lines 计数器+1
	//}
	//if err := scanner.Err(); err != nil { // check if error is generated
	//	fmt.Println("Error scanning for lines:", err)
	//	return
	//}
	//
	//fmt.Println("lines:", l)

}
