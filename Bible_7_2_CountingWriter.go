package main

import (
	"fmt"
	"io"
	"os"
)

type CountWriter struct {
	WrittenCnt int64
	//WrittenCntPtr *int64
	OrigWriter io.Writer
}

func (w *CountWriter) Write(p []byte) (n int, err error) {
	w.WrittenCnt = 0
	n, err = w.OrigWriter.Write(p)
	w.WrittenCnt = int64(n)
	return
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	newWriter := CountWriter{OrigWriter: w}
	newWriter.WrittenCnt = 0
	//newWriter.WrittenCntPtr = &newWriter.WrittenCnt
	return &newWriter, &newWriter.WrittenCnt
}

func main() {
	//var p *int
	writer, p := CountingWriter(os.Stdout)
	//fmt.Fprintln(writer, "hello world.")
	fmt.Fprintf(writer, "CountingWriter: write:%s\n", "hhhhsssdd asdf")
	fmt.Println("Bytes written:", *p)
}
