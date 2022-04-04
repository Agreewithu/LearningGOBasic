package main

import "fmt"

func main() {
	var a string
	a = "rx"
	//pair<type:string,value:"rx">
	var allType interface{} //定义一个万能指针
	allType = a             //pair<type:string,value:"rx">
	//通过断言找到string,value
	str, _ := allType.(string)
	fmt.Println(str)
}
