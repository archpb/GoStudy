package main

import "fmt"

func main() {

	var myMap map[int]string = make(map[int]string) // definition 1（make）, initiated, and allocated mem
	fmt.Println(myMap)                              // nil
	myMap[1] = "1st element"
	myMap[2] = "2nd element"

	fmt.Println(myMap)

	var myMap2 map[int]string         // definition 2(nil), but not init, no memory is allocated
	fmt.Printf("myMap2:%v\n", myMap2) // nil
	myMap2 = make(map[int]string)
	myMap2[1] = "11dd11" // add element
	myMap2[2] = "22aa22"
	myMap2[3] = "22332哈喽"
	fmt.Printf("after set, myMap2:%v\n", myMap2) // nil

	myMap3 := map[int]string{ // definition 3, init, and set key:value pairs
		10: "sss",
		20: "aaa",
		30: "444asdf", // 注意这个有个逗号
	}
	fmt.Println(myMap3)

	// delete
	delete(myMap3, 1)  // key 1 not existed, do nothing
	delete(myMap3, 20) // key found, then del
	fmt.Printf("after delete, myMap3: %v\n", myMap3)

	// traverse
	for k, v := range myMap3 {
		fmt.Println(k, v)
	}

	// find
	v, ok := myMap2[1] // found
	if ok {
		fmt.Printf("myMap2[1] exited, v=%v\n", v)
	} else {
		fmt.Printf("myMap2[1] not exited, v=%v\n", v)

	}
	v2, ok2 := myMap2[2222] // NOT found
	if ok2 {
		fmt.Printf("myMap2[2222] exited, v2=%v\n", v2)
	} else {
		fmt.Printf("myMap2[2222] not exited, v2=%v\n", v2)
	}

}
