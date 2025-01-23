package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

const listTemplateHtml string = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Map to Table</title>
    <style>
        table {
            width: 50%;
            border-collapse: collapse;
        }
        th, td {
            border: 1px solid black;
            padding: 8px;
            text-align: left;
        }
        th {
            background-color: #f2f2f2;
        }
    </style>
</head>
<body>
    <h2>My Shop Goods: </h2>
    <table>
        <thead>
            <tr>
                <th>Item</th>
                <th>Value</th>
            </tr>
        </thead>
        <tbody>
            {{range $key, $value := .}}
            <tr>
                <td>{{$key}}</td>
                <td>${{$value}}</td>
            </tr>
            {{end}}
        </tbody>
    </table>
</body>
</html>
`

// define custom handler
func myHttpHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
func helloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello baby")
}

// define custom handle
type MyMux map[string]float32

// Bible 7_12 Use html/template to show
func (m MyMux) showAllitems(w http.ResponseWriter, req *http.Request) {
	tmpl, err := template.New("list").Parse(listTemplateHtml)
	if err != nil {
		fmt.Println("Parse(listTemplateHtml) error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, m)
	if err != nil {
		fmt.Println("Execute(listTemplateHtml) error:", err)
		w.WriteHeader(http.StatusInternalServerError)

	}

}
func (mux MyMux) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Println("-----------------dump req--------------------")
	fmt.Println("URL:", req.URL)
	fmt.Println("RequestURI:", req.RequestURI)
	fmt.Println("URL.Path:", req.URL.Path)
	fmt.Println("Host:", req.Host)
	fmt.Println("method:", req.Method)
	fmt.Println("Header:", req.Header)
	fmt.Println("Body:", req.Body)
	fmt.Println("-----------------req  end--------------------")

	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "Hello, This is my shop. ")

	case "/list":
		//fmt.Fprintf(w, "Item Name: Price\n")
		//for name, price := range mux {
		//	fmt.Fprintf(w, "%s: $%4.2f\n", name, price)
		//}

		mux.showAllitems(w, req)

	case "/items": // host/items?name=name1&name=name2,
		items := req.URL.Query()
		fmt.Printf("Items: %s\n", items)
		itemNames := items["name"]
		fmt.Printf("Item names: %s\n", itemNames)
		for _, n := range itemNames {
			if value, ok := mux[n]; ok {
				fmt.Fprintf(w, "%s\t$%4.2f\n", n, value)
			} else {
				fmt.Fprintf(w, "%s\tNot Found.\n", n)
			}
		}
	case "/update": // host/update?item=name1&price=int,
		newItem := strings.ToLower(req.URL.Query().Get("item"))
		newPrice := req.URL.Query().Get("price")
		fmt.Printf("New Item: %s:%v\n", newItem, newPrice)
		p, err := strconv.ParseFloat(newPrice, 64)
		if err != nil || p < 0 || newItem == "" || len(newItem) > 20 {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Invalid price or items: %s\n", newPrice)
			return
		}
		mux[newItem] = float32(p)
		fmt.Fprintf(w, "New Item added: %s:%v\n", newItem, mux[newItem])

	default:
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "url error. %s", req.URL.Path)
	}

}
func main() {
	// 注册 path to handler
	http.HandleFunc("/", myHttpHandler)                       // 这里go会自动将myHttpHandler转换为http.HandlerFunc类型
	http.HandleFunc("/hello", http.HandlerFunc(helloHandler)) // 这里是手动转换成http.HandlerFunc类型, 两种都可以

	// 启动server， ip:port, 处理器
	addr := "localhost:8111"
	fmt.Println("server1 started at " + addr)
	go func() {
		err := http.ListenAndServe(addr, nil) // handler 使用默认的处理器DefaultServeMux
		if err != nil {
			fmt.Println(err)
		}
	}()

	/* 在Server2中，使用自定义的处理器*/
	mux := MyMux{
		"books": 22.1,
		"foods": 10.50,
		"tools": 50,
		"toys":  30.50,
	}
	addr2 := "localhost:8222"
	fmt.Println("server2 started at " + addr2)
	err := http.ListenAndServe(addr2, mux) // 使用自己的mux
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("server started at " + addr)

}
