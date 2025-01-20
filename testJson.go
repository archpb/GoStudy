package main

import (
	"encoding/json"
	"fmt"
)

type Contact struct {
	QQ      string `json:"qq"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
	Address string `json:"address"`
}
type Personal struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Sex     string
	IsPass  bool
	Contact Contact `json:"contactMethods"`
}

func main() {
	p := Personal{Name: "Alice", Age: 30, Sex: "Male", IsPass: true,
		Contact: Contact{
			QQ: "5652211", Phone: "13945654444", Email: "haha@qq.com", Address: "SHanghai",
		},
	}
	jsonData, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}
	fmt.Println(string(jsonData)) // 输出: {"name":"Alice","age":30}

	var p2 Personal
	err = json.Unmarshal(jsonData, &p2)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}
	fmt.Println(p2)

	fmt.Println("==========use Indent print json============")
	jsonData2, err := json.MarshalIndent(p, "", "   ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return

	}
	fmt.Println(string(jsonData2))
}
