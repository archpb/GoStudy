package main

import (
	"fmt"
	"io"
	"os"
)

type LimitRead struct {
	r io.Reader // 被封装的原始 Reader
	n int64     // 剩余可读取的字节数
}

// LimitReader 函数：返回一个新的 io.Reader，它读取 r 中最多 n 字节
func LimitReader(r io.Reader, n int64) io.Reader {
	return &LimitRead{r: r, n: n}
}

// 实现 io.Reader 接口
func (lr *LimitRead) Read(p []byte) (n int, err error) {
	if lr.n <= 0 {
		// 如果没有剩余字节可以读取，返回 EOF
		return 0, io.EOF
	}

	// 如果要读取的字节数超过剩余字节数，只读取剩余字节数
	if int64(len(p)) > lr.n {
		p = p[:lr.n]
	}

	// 从原始 Reader 读取数据
	n, err = lr.r.Read(p)

	// 更新剩余可读取字节数, 需要多次读取的case
	lr.n -= int64(n)

	// 如果已经读取到 n 字节，返回 EOF
	if lr.n <= 0 && err == nil {
		err = io.EOF // 安装limitReader的要求， 最多值读取n个字节， ，读满时， 并io.EOF作为err返回
	}

	return n, err
}

func main() {
	// 创建一个示例 Reader (可以是任何实现了 io.Reader 接口的类型)
	//data := "Hello, Go!"
	//reader := LimitReader((data), 5)
	file, err := os.Open("README.md")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// 创建一个LimitReader，限制最多读取10个字节
	reader := io.LimitReader(file, 10)

	// 使用新的 Reader 读取数据
	buf := make([]byte, 100) // 设置一个足够小的缓冲区
	for {
		n, err := reader.Read(buf)
		if err == io.EOF {
			break
		}
		fmt.Printf("Read %d bytes: %s\n", n, buf[:n])
	}

	// 输出结果：
	// Read 3 bytes: Hel
	// Read 2 bytes: lo
}
