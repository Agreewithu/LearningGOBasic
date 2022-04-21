package main

import "fmt"

func main() {
	var j float64 = 3.5
	var i int = int(j)
	fmt.Printf("强制类型不能转常量的...\n")
	fmt.Printf("将float64类型的变量j(%f)->int类型的i(%d)\n", j, i)
	var k interface{} = nil
	fmt.Printf("查看k的数据类型%T\n", k)
	fmt.Printf("查看k的值%v\n", k)
}
