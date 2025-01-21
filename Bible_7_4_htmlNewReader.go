package main

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

type MyHtmlReader struct {
	len int
	s   string
}

func NewReader(r string) io.Reader {
	reader := MyHtmlReader{s: r, len: len(r)}
	return &reader
}

func (p *MyHtmlReader) Read(b []byte) (n int, err error) {
	docPtr, err := html.Parse(strings.NewReader(p.s))
	if err != nil {
		fmt.Println("html Parse fail:", err)
		return 0, err
	}
	buf := new(bytes.Buffer)
	// traverse nodes tree, add data and attr to []byte
	for n := range docPtr.Descendants() {
		// data name
		buf.WriteString(n.Data)
		// type
		buf.WriteString(",Type:")
		buf.WriteString(strconv.Itoa(int(n.Type)))
		buf.WriteString(";")
		// attr
		for _, a := range n.Attr {
			buf.WriteString(a.Key)
			buf.WriteString("=")
			buf.WriteString(a.Val)
			buf.WriteString(",")
		}
		buf.WriteString("\n")
	}
	fmt.Println("Read finished:\n", buf.String())
	copy(b, buf.Bytes())
	return len(buf.Bytes()), nil
}

func main() {

	//s := `<p>Links:</p><ul><li><a href="foo">Foo</a><li><a href="/bar/baz">BarBaz</a></ul>`
	io.LimitReader(NewReader("Hello World!"), 50)
	s := `
	<html>
		<head><title>Example</title></head>
		<body>
			<h1>Hello, World!</h1>
			<p>This is a paragraph.</p>
		</body>
	</html>
	`
	reader := NewReader(s)
	b := make([]byte, 1024)
	reader.Read(b)

	//	doc, err := html.Parse(os.Stdin)
	//	if err != nil {
	//		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
	//		os.Exit(1)
	//	}
	//	for _, link := range visit(nil, doc) {
	//		fmt.Println(link)
	//	}
}

//// visit appends to links each link found in n and returns the result.
//func visit(links []string, n *html.Node) []string {
//	if n.Type == html.ElementNode && n.Data == "a" {
//		for _, a := range n.Attr {
//			if a.Key == "href" {
//				links = append(links, a.Val)
//			}
//		}
//	}
//	for c := n.FirstChild; c != nil; c = c.NextSibling {
//		links = visit(links, c)
//	}
//	return links
//}

/*
//!+html
package html

type Node struct {
	Type                    NodeType
	Data                    string
	Attr                    []Attribute
	FirstChild, NextSibling *Node
}

type NodeType int32

const (
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
)

type Attribute struct {
	Key, Val string
}

func Parse(r io.Reader) (*Node, error)
//!-html
*/
