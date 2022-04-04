package main

import "fmt"

func main() {

	//使用一下defer
	defer fmt.Println("main: hello go end1")
	defer fmt.Println("main: hello go end2") //它先出栈
	fmt.Println("main: hello go 1")
	fmt.Println("main: hello go 2")
}
